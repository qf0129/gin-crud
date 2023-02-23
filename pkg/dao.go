package pkg

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func QueryPage[T GormModel](fixed FixedOption, filterOptions []QueryOption) (any, error) {
	var total int64
	var items []T

	query := DB.Model(new(T))
	for _, option := range filterOptions {
		query = option(query)
	}
	query.Count(&total)

	for _, field := range strings.Split(fixed.Preload, ",") {
		query = OptionPreload(field)(query)
	}
	query = OptionOrderBy(fixed.OrderBy, fixed.Descending)(query)

	if fixed.ClosePaging {
		ret := query.Find(&items)
		if ret.Error != nil {
			logrus.Error(ret.Error)
			return nil, ret.Error
		} else {
			return items, nil
		}
	} else {
		if fixed.Page == 0 {
			fixed.Page = DEFAULT_PAGE_INDEX
		}
		if fixed.PageSize == 0 {
			fixed.PageSize = DEFAULT_PAGE_SIZE
		}
		query = OptionWithPage(fixed.Page, fixed.PageSize)(query)
		ret := query.Find(&items)
		if ret.Error != nil {
			logrus.Error(ret.Error)
			return nil, ret.Error
		} else {
			return PageBody{
				List:     &items,
				Page:     fixed.Page,
				PageSize: fixed.PageSize,
				Total:    total,
			}, nil
		}
	}
}

func QueryAll[T GormModel](fixed FixedOption, filterOptions []QueryOption) (any, error) {
	var items []T

	query := DB.Model(new(T))
	for _, option := range filterOptions {
		query = option(query)
	}

	for _, field := range strings.Split(fixed.Preload, ",") {
		query = OptionPreload(field)(query)
	}
	query = OptionOrderBy(fixed.OrderBy, fixed.Descending)(query)

	ret := query.Find(&items)
	if ret.Error != nil {
		logrus.Error(ret.Error)
		return nil, ret.Error
	} else {
		return items, nil
	}
}

func QueryOne[T GormModel](modelId any, preload string) (any, error) {
	var item T
	query := DB.Model(new(T)).Where(map[string]any{PRIMARY_KEY: modelId})
	for _, field := range strings.Split(preload, ",") {
		query = OptionPreload(field)(query)
	}

	ret := query.Take(&item)
	if ret.Error != nil {
		logrus.Error(ret.Error)
		return nil, ret.Error
	} else {
		return item, nil
	}
}

func QueryByID[T GormModel](modelId any, result any) error {
	return DB.Model(new(T)).Where(map[string]any{PRIMARY_KEY: modelId}).Take(&result).Error
}

func CreateOne[T GormModel](obj any) error {
	return DB.Create(obj).Error
}
func CreateOneWithMap[T GormModel](obj map[string]any) error {
	return DB.Model(new(T)).Create(obj).Error
}

func UpdateOne[T GormModel](modelId any, params map[string]any) error {
	return DB.Model(new(T)).Where(map[string]any{PRIMARY_KEY: modelId}).Updates(params).Error
}

func DeleteOne[T GormModel](modelId any) error {
	var existModel T
	err := QueryByID[T](modelId, &existModel)
	if err != nil {
		return err
	}
	return DB.Delete(&existModel).Error
}
