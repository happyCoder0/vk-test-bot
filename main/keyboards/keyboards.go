package keyboards

import (
	"vk-test-task/main/consts/colors"
	"vk-test-task/main/consts/templates"

	"github.com/SevereCloud/vksdk/v2/object"
)

var firstLayerKeyboard = object.NewMessagesKeyboard(false)
var secondLayerKeyboard = object.NewMessagesKeyboard(false)

func GetFirstLayerKeyboard() *object.MessagesKeyboard {
	if len(firstLayerKeyboard.Buttons) == 0 {
		row1 := firstLayerKeyboard.AddRow()
		row1.AddTextButton(templates.RandomNumber, "", colors.Secondary)
		row1.AddTextButton(templates.DayOfWeek, "", colors.Secondary)
		row2 := firstLayerKeyboard.AddRow()
		row2.AddTextButton(templates.RandomEmoji, "", colors.Secondary)
		row2.AddTextButton(templates.AnimalFact, "", colors.Secondary)
	}

	return firstLayerKeyboard
}

func GetSecondLayerKeyboard() *object.MessagesKeyboard {
	if len(secondLayerKeyboard.Buttons) == 0 {
		row := secondLayerKeyboard.AddRow()
		row.AddTextButton(templates.KittyFace, "", colors.Secondary)
		row.AddTextButton(templates.DoggoFace, "", colors.Secondary)
	}

	return secondLayerKeyboard
}
