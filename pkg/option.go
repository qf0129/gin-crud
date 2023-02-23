package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QueryOption func(tx *gorm.DB) *gorm.DB

type FixedOption struct {
	ClosePaging bool   `form:"close_paging"` // 关闭分页，默认false
	Page        int    `form:"page"`         // 页数，默认1
	PageSize    int    `form:"page_size"`    // 每页数量，默认10
	OrderBy     string `form:"order_by"`     // 排序字段名
	Descending  bool   `form:"desc"`         // 是否倒序，默认false
	Preload     string `form:"preload"`      // 预加载表名，以英文逗号分隔
}

const (
	OPTION_CLOSE_PAGING = "close_paging"
	OPTION_PAGE         = "page"
	OPTION_PAGE_SIZE    = "page_size"
	OPTION_ORDER_BY     = "order_by"
	OPTION_DESCENDING   = "desc"
	OPTION_PRELOAD      = "preload"
)

var FIXED_OPTIONS = []string{OPTION_CLOSE_PAGING, OPTION_PAGE, OPTION_PAGE_SIZE, OPTION_ORDER_BY, OPTION_DESCENDING, OPTION_PRELOAD}

func PraseFilterOptions(c *gin.Context) ([]QueryOption, error) {
	var options []QueryOption
	for k := range c.Request.URL.Query() {
		if !ArrHasStr(FIXED_OPTIONS, k) {
			options = append(options, OptionFilterBy(k, c.Query(k)))
		}
	}
	return options, nil
}

func OptionPreload(field string, options ...QueryOption) QueryOption {
	return func(tx *gorm.DB) *gorm.DB {
		if field == "" {
			return tx
		} else if field == "*" {
			return tx.Preload(clause.Associations)
		} else {
			return tx.Preload(cases.Title(language.Dutch).String(field), func(tx *gorm.DB) *gorm.DB {
				for _, option := range options {
					tx = option(tx)
				}
				return tx
			})
		}
	}
}

func OptionWithPage(pageIndex int, pageSize int) QueryOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	}
}

func OptionOrderBy(field string, descending bool) QueryOption {
	text := fmt.Sprintf("`%s`", field)
	if descending {
		text += " desc"
	}
	return func(tx *gorm.DB) *gorm.DB {
		if field == "" {
			return tx
		} else {
			return tx.Order(text)
		}
	}
}

func OptionFilterBy(field string, value any) QueryOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(map[string]any{field: value})
	}
}

func OptionWhere(query any, args ...any) QueryOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}
