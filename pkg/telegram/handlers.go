package telegram

import (
	"context"
	"github.com/dstotijn/go-notion"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// Команды
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды не существует :(")

	// Создание клиента для Notion API
	notionClient := notion.NewClient("NOTION_API_TOKEN")

	switch message.Command() {
	case "start":
		msg.Text = "Привет! Я твой бот :) " +
			"\nСписок команд: " +
			"\n/newtask - добавить задачу"
		_, err := b.bot.Send(msg)
		return err
	case "newtask":

		// Получение текста задачи после команды
		taskText := message.CommandArguments()

		// Создание новой страницы в Notion
		_, err := notionClient.CreatePage(context.Background(), notion.CreatePageParams{
			ParentType: notion.ParentTypeDatabase,
			ParentID:   "NOTION_DB_ID",
			Title: []notion.RichText{
				{
					Text: &notion.Text{
						Content: taskText,
					},
				},
			},
		})
		// Отправка сообщения пользователю
		if err != nil {
			msg.Text = "Ошибка создания задачи в Notion: " + err.Error()
		} else {
			msg.Text = "Задача создана в Notion!"
		}
	default:
		_, err := b.bot.Send(msg)
		return err
	}
	_, err := b.bot.Send(msg)
	return err
}
// Эхо бот
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
