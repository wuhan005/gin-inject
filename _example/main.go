package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	inject "github.com/wuhan005/gin-inject"
)

type inputForm struct {
	Name  string `json:"name"`
	Motto string `json:"motto"`
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/init", inject.Warp(
		func(c *gin.Context, sess sessions.Session) {
			// Set session.
			sess.Set("name", "E99p1ant")
			sess.Save()

			c.JSON(200, "Hello World")
		}),
	)

	r.POST("/input", inject.Warp(
		inject.BindJSON(inputForm{}),
		func(c *gin.Context, f *inputForm) {
			c.JSON(200, gin.H{
				"your_name":  f.Name,
				"your_motto": f.Motto,
			})
		}),
	)

	r.GET("/session", inject.Warp(
		func(c *gin.Context, sess sessions.Session) {
			c.JSON(200, gin.H{
				"name": sess.Get("name"),
			})
		}),
	)

	_ = r.Run()
}
