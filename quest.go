package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func battleHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "battle.tmpl", gin.H{
        "user": "Anyad",
    })
}

// AuthRequired will authorize requests.
func AuthRequired(c *gin.Context) {
    c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{})
    c.Abort() // This is how to stop rendering further if the auth failed
}

func main() {
    router := gin.Default()
    router.Static("/css", "./static/css")
    router.Static("/img", "./static/img")
    router.LoadHTMLGlob("templates/*")

    router.GET("/", indexHandler)

    authorized := router.Group("/battle")
    authorized.Use(AuthRequired)
    {
        authorized.GET("/", battleHandler)
    }

    router.Run(":9090")
}
