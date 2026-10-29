package application

import (
	"errors"
	"net/http"

	"github.com/ReiiSky/SharingVisionBackend/application/dto"
	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/gin-gonic/gin"
)

type updateArticleController struct {
	kernel  IKernel
	service updateArticleService
}

func NewUpdateArticleController(kernel IKernel) updateArticleController {
	return updateArticleController{
		kernel,
		updateArticleService{kernel},
	}
}

func (ctrl updateArticleController) Handle(c *gin.Context) {
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
	id := packages.ParseIntMin(0, 0, c.Param("id"))
	if err := ctrl.service.UpdateArticle(ctx, id, articlePayload); err != nil {

		if errors.Is(err, ArticleNotFoundError) {
			c.JSON(http.StatusNotFound, NewNotFoundResponse("article"))

			return
		}

		c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse[any](err.Error()))

		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse[any]("article updated", nil))
}
