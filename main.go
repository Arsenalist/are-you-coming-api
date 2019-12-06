package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS", "DELETE"}}))
	// Get user value
	r.GET("/event/:hash", func(c *gin.Context) {
		hash := c.Params.ByName("hash")
		event, err := GetEvent(hash)
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, event)
		}
	})
	r.PUT("/event", func(c *gin.Context) {
		var json Event
		c.ShouldBind(&json)
		e := CreateEvent(json.Name, json.UserId)
		c.JSON(http.StatusOK, e)
		return
	})
	r.POST("/event", func(c *gin.Context) {
		var json Event
		c.ShouldBind(&json)
		event, _ := GetEvent(json.Hash)
		if event.UserId == json.UserId {
			event.UpdateEventAttributes(json)
			SaveEvent(event)
		}
		c.JSON(http.StatusOK, event)
		return
	})
	r.POST("/event/rsvp", func(c *gin.Context) {
		var json Rsvp
		c.Bind(&json)
		eventHash := json.EventHash
		event, _ := GetEvent(eventHash)
		SaveRsvp(event, json.Name, json.UserId, json.Rsvp)
		c.JSON(http.StatusOK, event)
		return
	})
	r.DELETE("/event/rsvp", func(c *gin.Context) {
		var json Rsvp
		c.ShouldBind(&json)
		eventHash := json.EventHash
		event, _ := GetEvent(eventHash)
		DeleteRsvp(event, json.Name, json.UserId)
		c.JSON(http.StatusOK, event)
		return
	})
	return r
}

func main() {
	r := setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port not set, forcing to 8080")
		port = "8080"
	}
	r.Run(":" + port)
}
