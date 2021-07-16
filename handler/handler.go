package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/incrypto32/airdropbot/tgbotgo"
	tgtypes "github.com/incrypto32/airdropbot/tgbotgo/types"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	bot             *tgbotgo.TgBot
	commandHandlers map[string]func(*tgtypes.Update)
}

func New(e *echo.Echo, bot *tgbotgo.TgBot) *Handler {
	h := &Handler{bot: bot, commandHandlers: map[string]func(*tgtypes.Update){}}

	e.GET("/", h.webHookHandler)
	e.POST("/", h.webHookHandler)

	return h
}

func (h *Handler) registerCmd(cmd string, f func(*tgtypes.Update)) {
	h.commandHandlers[cmd] = f
}

func (h *Handler) registerCommandHandlers() {
	h.registerCmd("/start", func(u *tgtypes.Update) {
		h.bot.SendTextMessage(u.Message.Chat.ID, "Hi this is jimbralakka bot")
	})
	h.registerCmd("/whoami", func(u *tgtypes.Update) {
		h.bot.SendTextMessage(u.Message.Chat.ID, "Hi you suck")
	})
}

func (h *Handler) webHookHandler(c echo.Context) error {
	u := new(tgtypes.Update)

	if err := c.Bind(u); err != nil {
		log.Println(err)
	}

	h.registerCommandHandlers()

	if text := u.Message.Text; strings.HasPrefix(text, "/") {
		h.commandHandlers[text](u)
	}

	return c.NoContent(http.StatusOK)
}
