package handler

import (
	"github.com/handybots/azartio/azartio"
	"github.com/handybots/azartio/storage"

	"github.com/demget/don"
	tele "github.com/gesemaya/telegram"
	"github.com/gesemaya/telegram/layout"
)

func New(c Handler) handler {
	return handler{
		lt:   c.Layout,
		b:    c.Bot,
		db:   c.DB,
		rt:   c.Roulette,
		dons: c.Dons,
	}
}

type (
	Handler struct {
		Layout   *layout.Layout
		Bot      *tele.Bot
		DB       *storage.DB
		Roulette *azartio.Roulette
		Dons     *don.Dons
	}
	handler struct {
		lt   *layout.Layout
		b    *tele.Bot
		db   *storage.DB
		rt   *azartio.Roulette
		dons *don.Dons
	}
)
