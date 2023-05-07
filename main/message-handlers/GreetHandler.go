package messagehandlers

import (
	"strings"
	"vk-test-task/main/consts/templates"
	errorhandlers "vk-test-task/main/error-handlers"
	"vk-test-task/main/keyboards"
	messageutils "vk-test-task/main/message-utils"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

type GreetHandler struct{}

func (h *GreetHandler) Handle(vk *api.VK, ev events.MessageNewObject) {
	_, err := vk.MessagesSend(messageutils.CreateMessageWithKeyboard(
		"Привет, вот мой функционал:\n"+
			"1. Подкинуть кость\n"+
			"2. Показать день недели\n"+
			"3. Отправить рандомное эмодзи\n"+
			"4. Показать факты о животных\n",
		ev.Message.PeerID,
		keyboards.GetFirstLayerKeyboard(),
	))

	errorhandlers.HandleFatal(err)
}

func (h *GreetHandler) Predicate(command string) bool {
	return command == templates.Greet || strings.ToLower(command) == templates.Hi
}
