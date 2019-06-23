package endpoint

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/domain"
	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/utils"
)

// SearchResponse .
type SearchResponse struct {
	Results []*domain.Post `json:"results,omitempty"`
	Took    uint           `json:"took,omitempty"`
	Total   uint           `json:"total,omitempty"`
}

// SearchPost .
func SearchPost(c *gin.Context) {
	keyword := c.Query("q")

	if keyword == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	posts, took, total, err := domain.SearchByText(c.Request.Context(), keyword, 0, 10)
	if err != nil {
		log.Println(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")
		return
	}

	res := &SearchResponse{
		Results: posts,
		Took:    took,
		Total:   total,
	}

	c.JSON(200, res)
}
