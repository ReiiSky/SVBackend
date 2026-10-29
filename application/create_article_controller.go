package application

import (
	"net/http"

	"github.com/ReiiSky/SharingVisionBackend/application/dto"
	"github.com/gin-gonic/gin"
)

type createArticleController struct {
	kernel  IKernel
	service createArticleService
}

func NewCreateArticleController(kernel IKernel) createArticleController {
	return createArticleController{
		kernel,
		createArticleService{kernel},
	}
}

func (ctrl createArticleController) Handle(c *gin.Context) {
	articlePayload := dto.CreateArticleInput{}

	if err := c.ShouldBindJSON(&articlePayload); err != nil {
		c.JSON(http.StatusBadRequest, NewInvalidBodyResponse())

		return
	}

	if result := dto.Validate(ctrl.kernel.Validator(), articlePayload); result != nil {
		c.JSON(http.StatusBadRequest, NewInvalidBodyValueResponse(result.Msg()))

		return
	}

	ctx := c.Request.Context()
	if err := ctrl.service.CreateArticle(ctx, articlePayload); err != nil {
		c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse[any](err.Error()))

		return
	}

	c.JSON(http.StatusCreated, NewSuccessResponse[any]("article created", nil))
}
