package application

import (
	"errors"
	"net/http"

	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-gonic/gin"
)

type deleteArticleController struct {
	kernel  IKernel
	service deleteArticleService
}

func NewDeleteArticleController(kernel IKernel) deleteArticleController {
	return deleteArticleController{
		kernel,
		deleteArticleService{kernel},
	}
}

func (ctrl deleteArticleController) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	id := packages.ParseIntMin(0, 0, c.Param("id"))

	if err := ctrl.service.DeleteArticle(ctx, id); err != nil {
		if errors.Is(err, ArticleNotFoundError) {
			c.JSON(http.StatusNotFound, NewNotFoundResponse("article"))

			return
		}

		c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse[any](err.Error()))
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse[any]("article deleted", nil))
}
