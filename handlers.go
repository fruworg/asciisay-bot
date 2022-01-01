package main

import (
	"fmt"
	"time"
	"strings"
	"math/rand"
	
	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	reply := 
` < Если тебе было весело, то это не военное преступление. > 
    \
     \    v     v 
      \   |\---/|
       \ ( o_o  )
          \_W__/-..----.
       ___/    '   ' ,""+ \
     (__...'      __\    |'.___.';
      (_,...-'''(_,.'__)/'.....+`
	msg := fmt.Sprintf ("\n*Привет!* Присылай сообщение и я его преобразую.\n" + 
	       "\nС случайным животным:\nЕсли тебе было весело, то это не военное преступление." +
	       "\nС указанием животного:\n*cat* Если тебе было весело, то это не военное преступление.\n" +
	       "\nВместо *cat* можно поставить любое другое животное. \nТолько не забудь про пробел.\n" +
	       "\nЖивотные: *cow*, *cat*, *pig*, *bear*, *bat.*\n" + 
			    "\nПример ответа:\n" + "```%s```", reply)
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}

// Handle the msg command here
func (a *application) msgHandler(m *tbot.Message) {
	fruw := map[string]string{
		"cow": `
    \
     \   ^__^
      \  (oo)\_______
         (__)\       )\/\
              ||----w |
              ||     ||
	`,	//pig by jv
		"pig": `
    \
     \   <'--'\>______
      \  /. .  ''     \
        ('')  ,        @
         '-._,        /
              )-)_/--( >  
             ''''  ''''
	`,	//cat by sk
		"cat": `
    \
     \    v     v	
      \   |\---/|
       \ ( o_o  )
          \_W__/-..----.
       ___/    '   ' ,""+ \
     (__...'      __\    |'.___.';
       (_,...-'''(_,.'__)/'.....+
	`,	
		"aardvarks": `
	       _.---._    /\\       /
  	    ./'       "--'\//	   /
	  ./              o \     /
	 /./\  )______   \__ \	 /
	./  / /\ \   | \ \  \ \ 
	   / /  \ \  | |\ \  \7
 	   "     "    "  "     
	`,	//bat by jgs
		"bat": `
     \
 /\   \             /\
/ \'._ \ (\_/)   _.'/ \
|.''._'--(o.o)--'_.''.|
 \_ / ';=/ " \=;' \ _/
   '\__| \___/ |__/'
        \(_|_)/
         " ' "
	`,	
		"bear": `
     \
   __ \       __			
  /  \.-"""-./  \
  \    -   -    /
   |   o   o   |
   \  .-'''-.  /
    '-\__Y__/-'
       '---'
	`, 	//beaver by jgs
		"beaver": `
	          .="   "=._.---.
 	       ."         c 6 Y 6'p
 	      /   ,       '.  w_/
	      |   '-.   /     / 
	_,..._|      )_-\ \_=.\
       '-....-''------)))'=-'"''"
	`}
	rndmap := map[int]string{
		0: "cow",
		1: "pig",
		2: "cat",
		3: "bat",
		4: "bear"}
	msg := ""
	arr := strings.Split(m.Text, " ")
    	animal := arr[0]
	if fruw[animal] != "" && strings.TrimSpace(m.Text) != animal{
		reply := fruw[animal]
		text := strings.TrimPrefix(m.Text, animal + " ")
		msg = fmt.Sprintf("``` < %s > %s ```", text, reply)
	} else { 
		rand.Seed(time.Now().UnixNano())
		rnd := (rand.Intn(5))
		reply := fruw[rndmap[rnd]]
		msg = fmt.Sprintf("``` < %s > %s ```", m.Text, reply)
	}
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
