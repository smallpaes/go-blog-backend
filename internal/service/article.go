package service

type ArticleRequest struct {
	ID    uint32 `json:"id" binding:"required,gte=1"`
	State uint8  `json:"state" default:"1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `json:"tag_id" binding:"gte=1"`
	State uint8  `json:"state" default:"1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `json:"tag_id" binding:"required,gte=1"`
	Title         string `json:"title" binding:"required,min=2,max=100"`
	Desc          string `json:"desc" binding:"required,min=2,max=255"`
	Content       string `json:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `json:"cover_image_url" binding:"required,url"`
	CreatedBy     string `json:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `json:"state" default:"1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `json:"id" binding:"required,gte=1"`
	TagID         uint32 `json:"tag_id" binding:"gte=1"`
	Title         string `json:"title" binding:"min=2,max=100"`
	Desc          string `json:"desc" binding:"min=2,max=255"`
	Content       string `json:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `json:"cover_image_url" binding:"url"`
	ModifiedBy    string `json:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `json:"state" default:"1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `json:"id" binding:"required,gte=1"`
}
