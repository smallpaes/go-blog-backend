package v1

import "github.com/gin-gonic/gin"

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
// @Success 200 {object} model.Tag
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/vi/tags [get]
func (t Tag) List(c *gin.Context) {

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
