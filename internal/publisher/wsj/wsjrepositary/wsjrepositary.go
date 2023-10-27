package wsjrepositary

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BelyaevEI/news-parser/internal/config"
	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/model"
	"github.com/BelyaevEI/news-parser/internal/publisher/wsj/models"
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
	index1 := strings.Index(string(bodyBytes), "e1rxbks3 css-16ofghr-HeadlineTextBlock")
	text1 := string(bodyBytes)[index1 : index1+1000]
	index2 := strings.Index(text1, "</p>")
	title := text1[len("e1rxbks3 css-16ofghr-HeadlineTextBlock")+2 : index2]
	article.Title = title

	// Find description article in text parsing
	index4 := strings.Index(text1, "</style><p class=\"css-1mj1ort\">")
	text2 := text1[index4 : index4+250]
	index5 := strings.Index(text2, "</p>")
	description := text2[len("</style><p class=\"css-1mj1ort\">"):index5]
	article.Description = description

	// Fill source article
	article.Source = "Источник: " + "wsj.com"

	return article, nil
}
