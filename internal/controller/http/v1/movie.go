package v1

import (
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type movieRoutes struct {
	m usecase.Movie
	l logger.Interface
}

func newMovieRoutes(handler *gin.RouterGroup, m usecase.Movie, l logger.Interface) {
	r := &movieRoutes{m, l}

	h := handler.Group("/movies")
	{
		h.POST("/index")
	}
}

type indexMovieRequest struct {
	ReleaseYear string `json:"releaseYear" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Director    string `json:"director" binding:"required"`
	Cast        string `json:"cast" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Synopsis    string `json:"synopsis" binding:"required"`
}

// @Summary  create and index a document
// @Description create and index a document
// @ID          index
// @Tags  	    movies
// @Accept      json
// @Produce     json
// @Param       request body indexMovieRequest true "Set up movie"
// @Success     200 {object} movieResponse
// @Failure     500 {object} response
// @Router      /movies/index [post]
func (r *movieRoutes) doIndex(c *gin.Context){
	var request indexMovieRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		r.l.Error(err, "http - v1 - doIndex")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	movie , err := r.m.Index(
		c.Request.Context(),
		entity.Movie{
			ReleaseYear : request.ReleaseYear,
			Title       : request.Title,
			Director    : request.Director,
			Cast        : request.Cast,
			Genre       string `json:"genre"`
			Synopsis    string `json:"synopsis"`
		})
}
