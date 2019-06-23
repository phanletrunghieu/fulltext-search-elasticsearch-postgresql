package http

import (
	"github.com/gin-gonic/gin"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/endpoint"
)

// NewHTTPHandler .
func NewHTTPHandler() *gin.Engine {
	r := gin.Default()

	r.GET("/search", endpoint.SearchPost)

	return r
}
