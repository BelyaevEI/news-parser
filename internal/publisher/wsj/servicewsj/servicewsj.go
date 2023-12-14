package servicewsj

import (
	"time"

	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/BelyaevEI/news-parser/internal/model"
	newsmakerbot "github.com/BelyaevEI/news-parser/internal/newsmakerbot/bot"
	"github.com/BelyaevEI/news-parser/internal/publisher/wsj/wsjrepositary"
	"github.com/BelyaevEI/news-parser/internal/storage"
	"github.com/BelyaevEI/news-parser/internal/translater"
)

type Service struct {
	store *storage.Storage
	bot   *newsmakerbot.NewsMaker
}

func New(s *storage.Storage, bot *newsmakerbot.NewsMaker) *Service {
	return &Service{store: s,
		bot: bot}
}

// Parse article wsj.com and post to telegram channel
func (s *Service) RunWallStreetJournal() {
	for {
		// Get new article for post
		article, err := s.getNewArticles()
		if err != nil {
			time.Sleep(time.Minute * 5)
			continue
		}

		// Check posted article or not
		if posted := s.checkPostedArticle(article); posted {
			time.Sleep(time.Minute * 15)
			continue
		}

		post := article.Title + article.Description + article.Source

		// Posted article to telegram bot
		err = s.bot.SendMessage(post)
		if err != nil {
			continue
		}

		// Add article to storage
		s.store.AddArticle2Storage(article)

		// Change status of article
		s.changeStatusArticle(article)
	}
}

// Delete posted article in storage
func (s *Service) DeletePostedArt() {
	for {
		time.Sleep(time.Hour * 2)
		s.deletePostedArticle()
	}
}

func (s *Service) getNewArticles() (model.Article, error) {
	var article model.Article

	// Get article
	article, err := wsjrepositary.GetArticle()
	if err != nil || len(article.Title) == 0 || len(article.Description) == 0 {
		return model.Article{}, err
	}

	// Translate title of article to Russian
	article.Title, err = translater.Translate(article.Title)
	if err != nil {
		return model.Article{}, errors.Internal
	}

	article.Title += "." + "\n\n"

	// Translate description of article to Russian
	article.Description, err = translater.Translate(article.Description)
	if err != err {
		return model.Article{}, errors.Internal
	}

	article.Description += "\n\n"

	return article, nil
}

// Check posted article
func (s *Service) checkPostedArticle(article model.Article) bool {
	if posted := s.store.CheckerArticles(article); posted {
		return true
	}
	return false
}

// Change status article after post
func (s *Service) changeStatusArticle(article model.Article) {
	s.store.ChangeStatusArticle(article)
}

// Delete article after post
func (s *Service) deletePostedArticle() {
	s.store.DeletePostedArticle()
}
