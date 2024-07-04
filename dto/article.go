package dto

import "github.com/michaelwp/goblog/entities"

type ArticleExtend struct {
	UserName     string  `json:"user_name,omitempty"`
	CategoryName string  `json:"category_name,omitempty"`
	Page         *string `json:"page,omitempty"`
	Avatar       *string `json:"avatar,omitempty"`
}

type ArticleWithExtend struct {
	entities.Article
	ArticleExtend
}
