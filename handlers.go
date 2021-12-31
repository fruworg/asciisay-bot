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
	fruw := map[string]string{
		cow : `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
               ||----w |
               ||     ||
	`
		pig : `
        \   <'--'\>______
         \  /. .  ''     \
           ('')  ,        @
            '-._,        /
               )-)_/--( >  jv
              ''''  ''''
	`}
	
	text := strings.TrimPrefix(m.Text, "cowsay ")
	lineLen := utf8.RuneCountInString(text) + 2
	topLine := fmt.Sprintf(" %s ", strings.Repeat("_", lineLen))
	textLine := fmt.Sprintf("< %s >", text)
	bottomLine := fmt.Sprintf(" %s ", strings.Repeat("-", lineLen))
	msg := fmt.Sprintf("```%s\n%s\n%s%s```", topLine, textLine, bottomLine, fruw[pig])
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
