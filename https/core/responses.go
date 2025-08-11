package https_core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	fn        gin.HandlerFunc
	method    string
	addresses []string
	protected bool
}

var (
	responses = map[string]Response{
		"not-found-screen": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusNotFound, "not-found.html", nil)
			},
			method: "GET",
		},
		"index": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusOK, "index.html", nil)
			},
			method:    "GET",
			addresses: []string{"/", "/index.html"},
		},
		"projects": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusOK, "projects.html", nil)
			},
			method:    "GET",
			addresses: []string{"/projects", "/projects.html"},
		},
		"members": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusOK, "members.html", nil)
			},
			method:    "GET",
			addresses: []string{"/members", "/members.html"},
		},
		"technologies": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusOK, "technologies.html", nil)
			},
			method:    "GET",
			addresses: []string{"/technologies", "/technologies.html"},
		},
		"colleagues": {
			fn: func(c *gin.Context) {
				c.HTML(http.StatusOK, "colleagues.html", nil)
			},
			method:    "GET",
			addresses: []string{"/colleagues", "/colleagues.html"},
		},
	}
)
