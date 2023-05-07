package messagehandlers

import (
	"log"
	"vk-test-task/main/consts/templates"
	errorhandlers "vk-test-task/main/error-handlers"
	factfetchers "vk-test-task/main/fact-fetchers"
	"vk-test-task/main/keyboards"
	messageutils "vk-test-task/main/message-utils"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

type AnimalFactHandler struct{}

func (h *AnimalFactHandler) Handle(vk *api.VK, ev events.MessageNewObject) {
	if ev.Message.Text == templates.AnimalFact {
		log.Println(templates.AnimalFact)
		_, err := vk.MessagesSend(messageutils.CreateMessageWithKeyboard(
			"О каком животном вы хотите узнать больше?",
			ev.Message.PeerID,
			keyboards.GetSecondLayerKeyboard(),
		))

		errorhandlers.HandleFatal(err)
	} else {
		log.Println(templates.KittyFace)
		var (
			ch chan string
			f  func(chan string)
		)

		ch = make(chan string)

		if ev.Message.Text == templates.KittyFace {
			f = factfetchers.FetchCatFact
		} else {
			f = factfetchers.FetchDogFact
		}

		go f(ch)
		response := <-ch

		_, err := vk.MessagesSend(messageutils.CreateMessageWithKeyboard(
			response,
			ev.Message.PeerID,
			keyboards.GetFirstLayerKeyboard(),
		))

		errorhandlers.HandleFatal(err)
	}
}

func (h *AnimalFactHandler) Predicate(command string) bool {
	return command == templates.AnimalFact ||
		command == templates.KittyFace ||
		command == templates.DoggoFace
}
