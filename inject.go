package gin_inject

import (
	"log"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-macaron/inject"
)

var globalInjection = inject.New()

func Warp(handlers ...interface{}) func(c *gin.Context) {
	inj := inject.New()
	inj.SetParent(globalInjection)

	return func(c *gin.Context) {
		session := sessions.Default(c)
		inj.Map(c)
		inj.Map(session)

		for _, handler := range handlers {
			val, err := inj.Invoke(handler)
			if err != nil {
				log.Fatalf("Failed to invoke: %v", err)
			}
			if len(val) != 0 {
				switch val[0].Interface().(type) {
				case formType:
					inj.Map(reflect.ValueOf(val[0].Interface()).Interface())
				}
			}
		}
	}
}
