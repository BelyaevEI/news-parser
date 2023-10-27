package utils

import "github.com/BelyaevEI/news-parser/internal/model"

func RemoveValue(slice []model.Article) []model.Article {
	var result []model.Article
	for _, item := range slice {
		if !item.Posted {
			result = append(result, item)
		}
	}
	return result
}
