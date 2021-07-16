package tgbotgo_test

import (
	"testing"

	"github.com/incrypto32/airdropbot/tgbotgo"
	"github.com/stretchr/testify/assert"
)

func TestBot(t *testing.T) {
	t.Parallel()

	_, err := tgbotgo.New("TOKEN", "httpsgoogle.com")
	assert.Error(t, err)

	bot, err := tgbotgo.New("TOKEN", "https://google.com")
	assert.NoError(t, err)
	assert.Equal(t, bot.Url, "https://google.com")
}
