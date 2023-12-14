package storage

import (
	"sync"

	"github.com/BelyaevEI/news-parser/internal/model"
	"github.com/BelyaevEI/news-parser/internal/utils"
)

type Storage struct {
	store []model.Article
	mutex sync.Mutex
}

// Create a new storage for article
func New() *Storage {
	var s Storage

	s.store = make([]model.Article, 0)
	return &s
}

// Check status article in storage
func (s *Storage) CheckerArticles(article model.Article) bool {
	for _, val := range s.store {
		if val.Title == article.Title && val.Description == article.Description {
			return val.Posted
		}
	}
	return false
}

// Add article in storage
func (s *Storage) AddArticle2Storage(article model.Article) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store = append(s.store, article)
}

// Change status article in storage
func (s *Storage) ChangeStatusArticle(article model.Article) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i := 0; i < len(s.store); i++ {
		if article.Title == s.store[i].Title && article.Description == s.store[i].Description {
			s.store[i].Posted = true
		}
	}
}

// Delete posted article
func (s *Storage) DeletePostedArticle() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store = utils.RemoveValue(s.store)
}
