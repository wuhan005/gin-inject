package gin_inject

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type formType interface{}

func BindJSON(form interface{}) func(c *gin.Context) (formType, interface{}) {
	return func(c *gin.Context) (formType, interface{}) {
		ft := reflect.TypeOf(form)
		fv := reflect.New(ft)
		_ = c.BindJSON(fv.Interface())
		return formType(fv.Interface()), fv.Interface()
	}
}
