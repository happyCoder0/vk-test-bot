package messagehandlers

import (
	"math/rand"
	"vk-test-task/main/consts/templates"
	errorhandlers "vk-test-task/main/error-handlers"
	messageutils "vk-test-task/main/message-utils"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

var emojis = [5]string{
	"🤑",
	"😎",
	"👽",
	"🤡",
	"💩",
}

type RandomEmojiHandler struct{}

func (h *RandomEmojiHandler) Handle(vk *api.VK, ev events.MessageNewObject) {
	response := emojis[rand.Intn(5)]

	_, err := vk.MessagesSend(messageutils.CreateMessage(
		response,
		ev.Message.PeerID,
	))

	errorhandlers.HandleFatal(err)
}

func (h *RandomEmojiHandler) Predicate(command string) bool {
	return command == templates.RandomEmoji
}
