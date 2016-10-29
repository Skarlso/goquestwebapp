package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthenticateRequest is used to authenticate a request for a certain end-point group.
// TODO: Use this as a middleware to authorize a request going to /battle/*
func AuthenticateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("user-id")
		if v == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// session.Set("count", count)
		// session.Save()
		c.Next()
	}
}
