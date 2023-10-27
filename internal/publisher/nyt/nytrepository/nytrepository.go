package nytrepository

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BelyaevEI/news-parser/internal/config"
	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/model"
	"github.com/BelyaevEI/news-parser/internal/publisher/nyt/models"
)

func GetArticle() (model.Article, error) {

	var article model.Article

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, models.NYT, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", config.SETTINGSQUERY)
	resp, err := client.Do(request)
	if err != nil {
		return model.Article{}, errors.Internal
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Article{}, errors.Internal
	}

	// Find title article in text parsing
	index1 := strings.Index(string(bodyBytes), "css-miszbp e1hr934v2")
	text1 := string(bodyBytes)[index1 : index1+1000]
	index2 := strings.Index(text1, "data-rref=\"\">")
	index3 := strings.Index(text1, "</a>")
	title := text1[index2+len("data-rref=\"\">") : index3]
	article.Title = title

	// Find description article in text parsing
	index4 := strings.Index(text1, "css-tskdi9 e1hr934v5\">")
	index5 := strings.Index(text1, "</p>")
	description := text1[index4+len("css-tskdi9 e1hr934v5\">") : index5]
	article.Description = description

	// Fill source article
	article.Source = "Источник: " + "nytimes.com"

	return article, nil

}
