package handler

import (
	"log"

	tele "github.com/gesemaya/telegram"
)

func (h handler) OnError(err error, c tele.Context) {
	log.Println(c.Sender().Recipient(), err)
}
