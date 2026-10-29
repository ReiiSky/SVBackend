package application

import (
	"net/http"

	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-gonic/gin"
)

type getDetailArticleController struct {
	kernel  IKernel
	service getDetailArticleService
}

func NewGetDetailArticleController(kernel IKernel) getDetailArticleController {
	return getDetailArticleController{
		kernel,
		getDetailArticleService{kernel},
	}
}

func (ctrl getDetailArticleController) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	id := packages.ParseIntMin(0, 0, c.Param("id"))
	article, err := ctrl.service.GetArticleByID(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse[any](err.Error()))

		return
	}

	if article == nil {
		c.JSON(http.StatusNotFound, NewNotFoundResponse("article"))

		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("article found", article))
}
