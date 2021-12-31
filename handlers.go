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
	`,
		"cat": `
   	\   |\---/|
   	 \ ( o_o  )
            \_'__/-..----.
 	   ___/ '   ' ,""+ \  sk
	  (__...'   __\    |'.___.';
  	    (_,...'(_,.'__)/'.....+
	`,	
		"aardvarks": `
	       _.---._    /\\       /
  	    ./'       "--'\//	   /
	  ./              o \     /
	 /./\  )______   \__ \	 /
	./  / /\ \   | \ \  \ \ 
	   / /  \ \  | |\ \  \7
 	   "     "    "  "     
	`,	
		"bat": `
	 /\   \             /\
	/ \'._ \ (\_/)   _.'/ \
	|.''._'--(o.o)--'_.''.|
	 \_ / ';=/ " \=;' \ _/
 	  '\__| \___/ |__/'
	jgs     \(_|_)/
         	 " ' "
	`,	
		"bear": `
	 __ \       __			
	/  \.-"""-./  \
	\    -   -    /
	 |   o   o   |
	 \  .-'''-.  /
 	  '-\__Y__/-'
  	     '---'
	`, 	
		"beaver": `
	          .="   "=._.---.
 	       ."         c 6 Y 6'p
 	      /   ,       '.  w_/
	  jgs |   '-.   /     / 
	_,..._|      )_-\ \_=.\
       '-....-''------)))'=-'"''"
	`
	}
	msg := ""
	arr := strings.Split(m.Text, " ")
    	animal := arr[0]
	if fruw[animal] != ""{
		reply := fruw[animal]
		text := strings.TrimPrefix(m.Text, animal + " ")
		lineLen := utf8.RuneCountInString(text) + 2
		topLine := fmt.Sprintf(" %s ", strings.Repeat("_", lineLen))
		textLine := fmt.Sprintf("< %s >", text)
		bottomLine := fmt.Sprintf(" %s ", strings.Repeat("-", lineLen))
		msg = fmt.Sprintf("```%s\n%s\n%s%s```", topLine, textLine, bottomLine, reply)
	} else { 
		reply := fruw["pig"]
		lineLen := utf8.RuneCountInString(m.Text) + 2
		topLine := fmt.Sprintf(" %s ", strings.Repeat("_", lineLen))
		textLine := fmt.Sprintf("< %s >", m.Text)
		bottomLine := fmt.Sprintf(" %s ", strings.Repeat("-", lineLen))
		msg = fmt.Sprintf("```%s\n%s\n%s%s```", topLine, textLine, bottomLine, reply)
	}
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
