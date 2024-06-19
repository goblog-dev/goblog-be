package model

type ArticleExtend struct {
	UserName     string  `json:"user_name"`
	CategoryName string  `json:"category_name"`
	Page         *string `json:"page"`
}

type ArticleWithExtend struct {
	Article
	ArticleExtend
}
