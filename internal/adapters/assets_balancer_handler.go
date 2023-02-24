package adapters

import (
	"balancer/internal/boundaries"
	"balancer/internal/ports"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AssetsBalancerHandler struct {
		useCase ports.AssetBalancerUseCase
	}
	errorResult struct {
		Errors []string
	}
)

func newErrorResult(errs ...string) *errorResult {
	return &errorResult{
		Errors: errs,
	}
}
func (h *AssetsBalancerHandler) HandleGetAssetsGroups(c *gin.Context) {
	res := h.useCase.GetAssetsGroups(c)

	c.JSON(http.StatusOK, res)
}

func (h *AssetsBalancerHandler) HandleGetAssetsGroup(c *gin.Context) {

	input := &boundaries.GetAssetsGroupInput{
		Id: uuid.MustParse(c.Param("id")),
	}
	res := h.useCase.GetAssetsGroup(c, input)

	c.JSON(http.StatusOK, res)
}

func (h *AssetsBalancerHandler) HandleCreateAssetsGroup(c *gin.Context) {
	input := &boundaries.CreateAssetsGroupInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	res, err := h.useCase.CreateAssetsGroup(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *AssetsBalancerHandler) HandleCreateAsset(c *gin.Context) {
	input := &boundaries.CreateAssetForGroupInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	res, err := h.useCase.CreateAsset(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *AssetsBalancerHandler) HandleUpdateAsset(c *gin.Context) {
	input := &boundaries.UpdateAssetInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	res, err := h.useCase.UpdateAsset(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AssetsBalancerHandler) HandleUpdateAssetsGroup(c *gin.Context) {
	input := &boundaries.UpdateAssetsGroup{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	res, err := h.useCase.UpdateAssetsGroup(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AssetsBalancerHandler) HandleDeleteAsset(c *gin.Context) {
	input := &boundaries.DeleteAssetInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	res, err := h.useCase.DeleteAsset(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AssetsBalancerHandler) HandleDeleteAssetsGroup(c *gin.Context) {
	input := &boundaries.DeleteAssetsGroupInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, newErrorResult(err.Error()))
		return
	}

	err := h.useCase.DeleteAssetsGroup(c, input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, newErrorResult(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func NewAssetsBalancerHandler(uc ports.AssetBalancerUseCase) *AssetsBalancerHandler {
	return &AssetsBalancerHandler{
		useCase: uc,
	}
}
