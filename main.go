package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static") // For static assets


  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl.html",nil)

  })
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}
