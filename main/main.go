package main

import (
	"context"
	"log"
	errorhandlers "vk-test-task/main/error-handlers"
	messagehandlers "vk-test-task/main/message-handlers"
	tokenutils "vk-test-task/main/token-utils"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

var token = tokenutils.RetrieveToken()

func main() {
	vk := api.NewVK(token)

	handlers := []messagehandlers.Handler{
		&messagehandlers.RandomNumberHandler{},
		&messagehandlers.DayOfWeekHandler{},
		&messagehandlers.RandomEmojiHandler{},
		&messagehandlers.GreetHandler{},
		&messagehandlers.AnimalFactHandler{},
	}

	group, err := vk.GroupsGetByID(nil)

	errorhandlers.HandleFatal(err)
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)

	errorhandlers.HandleFatal(err)

	lp.MessageNew(func(ctx context.Context, ev events.MessageNewObject) {
		log.Println("A message received:", ev.Message.Text)

		for _, h := range handlers {
			if h.Predicate(ev.Message.Text) {
				h.Handle(vk, ev)
				break
			}
		}
	})

	err = lp.Run()

	errorhandlers.HandleFatal(err)
}
