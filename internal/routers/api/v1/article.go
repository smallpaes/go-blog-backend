package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary Get single article
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Internal Error"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {

}

// @Summary Get multiple articles
// @Produce json
// @Param name query string false "Article name"
// @Param tag_id query int false "Tag ID"
// @Param state query int false "State"
// @Param page query int false "Amount of articles per page"
// @Success 200 {object} model.ArticleSwagger "List of articles"
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {

}

// @Summary Create an article
// @Produce json
// @Param tag_id body string true "Tag ID"
// @Param title body string true "Article title"
// @Param desc body string false "Article brief description"
// @Param cover_image_url body string true "Cover image source"
// @Param content body string true "Article content"
// @Param created_by body int true "Creator"
// @Param state body int false "State"
// @Success 200 {object} model.Article
// @Failure 400 {object} errcode.Error "Request error"
// @Failure 500 {object} errcode.Error "Internal error"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {

}

// @Summary Update an article
// @Produce json
// @Param id path int true "Article ID"
// @Param tag_id body string false "Tag ID"
// @Param title body string false "Article title"
// @Param desc body string false "Article brief description"
// @Param cover_image_url body string false "Cover image source"
// @Param content body string false "Article content"
// @Param modified_by body string true "Modifier"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {

}

// @Summary Delete an article
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {

}
