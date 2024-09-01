package dummygen

import (
	"someweb/cmd/models"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateDummy() []models.Article {
	var articles []models.Article

	for i := 0; i < 10; i++ {
		articles = append(articles,
			models.Article{
				Id:         i,
				PostedDate: time.Now(),
				Title:      gofakeit.Quote(),
				Content:    gofakeit.HipsterParagraph(2, 3, 20, ""),
			})
	}
	return articles
}
