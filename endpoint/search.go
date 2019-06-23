package endpoint

import (
	"log"
	"net/http"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/domain"

	"github.com/gin-gonic/gin"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/utils"
)

// SearchPost .
func SearchPost(c *gin.Context) {
	keyword := c.Query("q")

	if keyword == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	err := domain.SearchByText(c.Request.Context(), keyword, 0, 10)
	if err != nil {
		log.Println(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")
		return
	}
}
