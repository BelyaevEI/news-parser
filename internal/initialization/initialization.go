package initialization

import (
	"sync"
	"time"

	newsmakerbot "github.com/BelyaevEI/news-parser/internal/newsmakerbot/bot"
	"github.com/BelyaevEI/news-parser/internal/publisher/bloomberg/servicebloomberg"
	"github.com/BelyaevEI/news-parser/internal/publisher/economist/serviceco"
	"github.com/BelyaevEI/news-parser/internal/publisher/nyt/servicenyt"
	"github.com/BelyaevEI/news-parser/internal/publisher/rbc/servicerbc"
	"github.com/BelyaevEI/news-parser/internal/publisher/wsj/servicewsj"
	"github.com/BelyaevEI/news-parser/internal/storage"
)

type Services struct {
	Rbc       *servicerbc.Service
	Bloomberg *servicebloomberg.Service
	Nyt       *servicenyt.Service
	Wsj       *servicewsj.Service
	Eco       *serviceco.Service
}

// Initialization services
func New() *Services {

	// Init connect to bot
	bot, err := newsmakerbot.Connect()
	if err != nil {
		return &Services{}
	}

	// Init storage for services
	storagerbc := storage.New()
	storagebloomberg := storage.New()
	storagenyt := storage.New()
	storgaewsj := storage.New()
	storageco := storage.New()

	// Init services
	rbc := servicerbc.New(storagerbc, bot)
	bloomberg := servicebloomberg.New(storagebloomberg, bot)
	nyt := servicenyt.New(storagenyt, bot)
	wsj := servicewsj.New(storgaewsj, bot)
	eco := serviceco.New(storageco, bot)

	return &Services{Rbc: rbc,
		Bloomberg: bloomberg,
		Nyt:       nyt,
		Wsj:       wsj,
		Eco:       eco}
}

func (service *Services) RunServices() {

	var wg sync.WaitGroup

	// Run bloomberg
	go service.Bloomberg.RunBloomberg()
	go service.Bloomberg.DeletePostedArt()
	time.Sleep(time.Minute * 5)

	// Run economist
	go service.Eco.RunEconomist()
	go service.Eco.DeletePostedArt()
	time.Sleep(time.Minute * 5)

	// Run New York Times
	go service.Nyt.RunNewYorkTimes()
	go service.Nyt.DeletePostedArt()
	time.Sleep(time.Minute * 5)

	// Run rbc
	go service.Rbc.RunRBC()
	go service.Rbc.DeletePostedArt()
	time.Sleep(time.Minute * 5)

	// Run wall street journal
	go service.Wsj.RunWallStreetJournal()
	go service.Wsj.DeletePostedArt()

	wg.Add(1)

	wg.Wait()
}
