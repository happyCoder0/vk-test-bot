package messagehandlers

import (
	"vk-test-task/main/consts/templates"
	errorhandlers "vk-test-task/main/error-handlers"
	messageutils "vk-test-task/main/message-utils"

	"math/rand"

	"strconv"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

type RandomNumberHandler struct{}

func (h *RandomNumberHandler) Handle(vk *api.VK, ev events.MessageNewObject) {
	response := rand.Intn(6) + 1

	_, err := vk.MessagesSend(messageutils.CreateMessage(
		strconv.Itoa(response),
		ev.Message.PeerID,
	))
	errorhandlers.HandleFatal(err)
}

func (h *RandomNumberHandler) Predicate(command string) bool {
	return command == templates.RandomNumber
}
