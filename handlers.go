package main

import (
	"fmt"
	"strings"

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
	    v     v	
   	\   |\---/|
   	 \ ( o_o  )
	    \_v__/-..----.
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
	`}
	msg := ""
	arr := strings.Split(m.Text, " ")
    	animal := arr[0]
	if fruw[animal] != "" && strings.TrimSpace(m.Text) != animal{
		reply := fruw[animal]
		text := strings.TrimPrefix(m.Text, animal + " ")
		msg = fmt.Sprintf("```www www %s > %s```", text, reply)
	} else { 
		reply := fruw["pig"]
		msg = fmt.Sprintf("```www www %s > %s```", m.Text, reply)
	}
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
