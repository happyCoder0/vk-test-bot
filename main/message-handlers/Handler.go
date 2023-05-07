package messagehandlers

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

type Handler interface {
	Handle(*api.VK, events.MessageNewObject)
	Predicate(string) bool
}
