package rbcrepository

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/publisher/rbc/models"
	"github.com/PuerkitoBio/goquery"
)

// Get article title
func GetArticleTitle() (string, error) {
	response, err := http.Get(models.RBC)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", fmt.Errorf("create new doc: %s", err)
	}

	// Find article title and save in storage
	title := doc.Find(models.MAINTITLETAG).Text()
	if len(title) == 0 {
		return "", errors.TitleNotFound
	}
	title += "."
	title += "\n\n"
	return title, nil
}

// Get description article
func GetDescriptionArticle() (string, error) {
	response, err := http.Get(models.RBC)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", errors.NewDoc
	}

	defer response.Body.Close()

	// Find link article
	link, exists := doc.Find(models.MAINLINKTAG).Attr("href")
	if !exists {
		return "", errors.LinkNoFound
	}

	responseArticle, err := http.Get(link)
	if err != nil {
		return "", err
	}
	// Find description in new page
	docArticle, err := goquery.NewDocumentFromReader(responseArticle.Body)
	if err != nil {
		return "", errors.NewDoc
	}

	// Find description article
	description := docArticle.Find(models.OVERVIEWTAG).Text()
	if len(description) != 0 {
		description = strings.ReplaceAll(description, "\n", "")
		description = strings.TrimLeft(description, " ")
		description += "\n\n"
		return description, nil
	}
	return "", errors.DescrNotFound
}
