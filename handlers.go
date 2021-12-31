package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "бебра?"
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}

// Handle the msg command here
func (a *application) msgHandler(m *tbot.Message) {
	text := strings.TrimPrefix(m.Text, "cowsay ")
	lineLen := utf8.RuneCountInString(text) + 2
	topLine := fmt.Sprintf(" %s ", strings.Repeat("_", lineLen))
	textLine := fmt.Sprintf("< %s >", text)
	bottomLine := fmt.Sprintf(" %s ", strings.Repeat("-", lineLen))
	cow := `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
               ||----w |
               ||     ||
	`
	msg := fmt.Sprintf("```%s\n%s\n%s%s```", topLine, textLine, bottomLine, cow)
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
