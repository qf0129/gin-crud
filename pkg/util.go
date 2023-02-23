package pkg

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ArrHasStr(arr []string, target string) bool {
	for _, element := range arr {
		if target == element {
			return true
		}
	}
	return false
}

func GetModelNameLower(obj any) string {
	reflectType := reflect.TypeOf(obj)
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return CamelToSnakeCase(reflectType.Name())
}

func GetIdKey(obj any) string {
	reflectType := reflect.TypeOf(obj)
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return CamelToSnakeCase(reflectType.Name()) + "_id"
}

// AbcDef > abc_def
func CamelToSnakeCase(text string) string {
	temp := regexp.MustCompile(`([A-Z])`).ReplaceAllString(text, "_$1")
	temp = cases.Lower(language.Und).String(temp)
	temp = strings.TrimLeft(temp, "_")
	return temp
}

// abc_def > AbcDef
func SnakeToCamelCase(text string) string {
	temp := strings.ReplaceAll(text, "_", " ")
	temp = cases.Title(language.Und).String(temp)
	return strings.ReplaceAll(temp, " ", "")
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type PageBody struct {
	List     any   `json:"list"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

func RespSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
	c.Abort()
}

func RespFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: 40000,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
}
