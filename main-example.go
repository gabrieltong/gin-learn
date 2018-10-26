package main

import (
	"fmt"
	"net/http"
	"github.com/gabrieltong/gin-learn/sr"  
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	f1 := func(c *gin.Context) {
		c.String(http.StatusOK, src.Oli)
	}

	// Simple group: v1
	v1 := r.Group("/v1")
		gin.BasicAuth
		v1.POST("/login", f1)
		v1.POST("/submit", f1)
		v1.POST("/read", f1)
	

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", f1)
		v2.POST("/submit", f1)
		v2.POST("/read", f1)
	}


	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			s := gin.H{"user": user, "value": value}
			fmt.Printf(s["user"].(string))
			c.JSON(http.StatusOK, s)
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	
	authorized.POST("admin", func(c *gin.Context) {
		c.HTML
		user := c.MustGet(gin.AuthUserKey).(string)
		c.ShouldBind
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
