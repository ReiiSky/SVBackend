package application

import (
	"net/http"

	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-gonic/gin"
)

type getArticleController struct {
	kernel  IKernel
	service getArticleService
}

func NewGetArticleController(kernel IKernel) getArticleController {
	return getArticleController{
		kernel,
		getArticleService{kernel},
	}
}

func (ctrl getArticleController) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	offset := packages.ParseIntMin(0, 0, c.Query("offset"))
	limit := packages.ParseIntMin(10, 1, c.Query("limit"))

	articles, err := ctrl.service.GetArticles(ctx, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse[any](err.Error()))

		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("article created", articles))
}
