package ecorepository

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BelyaevEI/news-parser/internal/config"
	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/model"
	"github.com/BelyaevEI/news-parser/internal/publisher/economist/models"
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
	index1 := strings.Index(string(bodyBytes), "css-1nh4cha e1rr6cni0")
	text1 := string(bodyBytes)[index1 : index1+1000]
	index2 := strings.Index(text1, "top_stories:headline_1\">")
	index3 := strings.Index(text1, "</a>")
	title := text1[index2+len("top_stories:headline_1\">") : index3]
	article.Title = title

	// Find description article in text parsing
	index4 := strings.Index(text1, "css-sn9piy er8c6600\">")
	index5 := strings.Index(text1, "</p>")
	description := text1[index4+len("css-sn9piy er8c6600\">") : index5]
	article.Description = description

	// Fill source article
	article.Source = "Источник: " + "economist.com"

	return article, nil

}
