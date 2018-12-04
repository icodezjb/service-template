package controller

type ArticleController struct {
	baseController
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}
