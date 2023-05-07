package messagehandlers

import (
	"time"
	"vk-test-task/main/consts/templates"
	errorhandlers "vk-test-task/main/error-handlers"
	messageutils "vk-test-task/main/message-utils"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

type DayOfWeekHandler struct{}

func (h *DayOfWeekHandler) Handle(vk *api.VK, ev events.MessageNewObject) {
	response := time.Now().Weekday().String()

	_, err := vk.MessagesSend(messageutils.CreateMessage(
		response,
		ev.Message.PeerID,
	))

	errorhandlers.HandleFatal(err)
}

func (h *DayOfWeekHandler) Predicate(command string) bool {
	return command == templates.DayOfWeek
}
