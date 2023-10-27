package bloombergrepository

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BelyaevEI/news-parser/internal/config"
	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/publisher/bloomberg/models"
)

func GetArticleTitle() (string, error) {
	var article string

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, models.BLOOMBERG, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", config.SETTINGSQUERY)

	resp, err := client.Do(request)
	if err != nil {
		return "", errors.Internal
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Internal
	}

	// Find article in text parsing
	index1 := strings.Index(string(bodyBytes), "styles_storyInfo__0dHbp")
	alltext := string(bodyBytes)[index1 : index1+1000]
	index2 := strings.Index(alltext, "hover:underline focus:underline")
	text2 := alltext[index2 : index2+500]
	index3 := strings.Index(text2, "<a href")
	text3 := text2[index3 : index3+300]
	start := strings.Index(text3, ">")
	stop := strings.Index(text3, "</a>")
	article = text3[start+1 : stop]

	if len(article) == 0 {
		return "", errors.ArtNotFound
	}

	return article + ".", nil
}
