package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "GoQuest",
    })
}

func main() {
    router := gin.Default()
    router.Static("/css", "./static/css")
    router.Static("/img", "./static/img")
    router.LoadHTMLGlob("templates/*")

    router.GET("/", indexHandler)
    router.Run(":9090")
}
