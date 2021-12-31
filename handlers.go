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
	reply := ""
	fruw := map[string]string{
		"cow": `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
               ||----w |
               ||     ||
	`,
		"pig": `
        \   <'--'\>______
         \  /. .  ''     \
           ('')  ,        @
            '-._,        /
               )-)_/--( >  jv
              ''''  ''''
	`}
	
	arr := strings.Split(m.Text, " ")
    	fam := arr[0]
	text := strings.TrimPrefix(m.Text, fam + " ")
	lineLen := utf8.RuneCountInString(text) + 2
	topLine := fmt.Sprintf(" %s ", strings.Repeat("_", lineLen))
	textLine := fmt.Sprintf("< %s >", text)
	bottomLine := fmt.Sprintf(" %s ", strings.Repeat("-", lineLen))
	if fruw[animal] != ""{
		reply = fruw[animal]
	} else { 
		reply = fruw["pig"]
	}
	msg := fmt.Sprintf("```%s\n%s\n%s%s```", topLine, textLine, bottomLine, reply)
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
