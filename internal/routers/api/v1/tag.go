package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/smallpaes/go-blog-backend/global"
	"github.com/smallpaes/go-blog-backend/pkg/app"
	"github.com/smallpaes/go-blog-backend/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary Get multiple tags
// @Produce json
// @Param name query string false "Tag name" maxlength(10)
// @Param state query int false "State" Enum(0, 1) default(1)
// @Param page query int false "Page number"
// @Param page_size query int false "Amount per page"
// @Success 200 {object} model.TagSwagger "Success"
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	fmt.Println("jijo")
	param := struct {
		Name  string `json:"name" binding:"max=100"`
		State uint8  `json:"state" default:"1" binding:"oneof=0 1"`
	}{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Add new tag
// @Produce json
// @Param name body string true "Tag name" minlength(3) maxlength(100)
// @Param state body int false "State" Enums(0, 1) default(1)
// @Param created_by body string true "Creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "Success"
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

// @Summary Update tag
// @Produce json
// @Param id path int true "Tag ID"
// @Param name body string false "Tag name" minlength(3) maxlength(100)
// @Param state body int false "State" Enums(0, 1) default(1)
// @Param modified_by body string true "Modifier" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "Success"
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}

// @Summary Delete tag
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}
