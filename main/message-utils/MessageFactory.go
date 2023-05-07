package messageutils

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
)

func CreateMessage(text string, peerId int) api.Params {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message(text)
	b.PeerID(peerId)
	return b.Params
}

func CreateMessageWithKeyboard(text string,
	peerId int,
	keyboard *object.MessagesKeyboard) api.Params {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message(text)
	b.PeerID(peerId)
	b.Keyboard(keyboard)
	return b.Params
}
