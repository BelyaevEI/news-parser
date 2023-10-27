package newsmakerbot

import (
	"github.com/BelyaevEI/news-parser/internal/newsmakerbot/models"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type NewsMaker struct {
	Bot *tgbotapi.BotAPI
}

func Connect() (*NewsMaker, error) {
	bot, err := tgbotapi.NewBotAPI(models.TOKEN)
	if err != nil {
		return nil, err
	}
	return &NewsMaker{
		Bot: bot,
	}, nil
}

func (newsmaker *NewsMaker) SendMessage(msg string) error {
	message := tgbotapi.NewMessageToChannel(models.CHANNELUSERNAME, msg)
	_, err := newsmaker.Bot.Send(message)
	if err != nil {
		return err
	}
	return nil
}
