package main

import (
	"fmt"
	"time"
	"strings"
	"math/rand"
	
	"github.com/yanzay/tbot/v2"
)

var rndmap = map[int]string{
		0: "cow",
		1: "pig",
		2: "cat",
		3: "bat",
		4: "bear",
		5: "bison",
		6: "cock",
		7: "duck",
		8: "camel",
	        9: "deer",
		10: "dog",
		11: "dolphin",
		12: "elephant"}

var lenmap = len(rndmap)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	org := ""
	for i := 0; i < lenmap; i++{
		org = fmt.Sprintf("%s*%s*", org, rndmap[i])
		if i < lenmap-1{
			org = org + ", "}
	}
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
	       "\nСо случайным животным:\nЕсли тебе было весело, то это не военное преступление.\n" +
	       "\nС указанием животного:\n*cat* Если тебе было весело, то это не военное преступление.\n" +
	       "\nВместо *cat* можно подставить любое другое животное. Только не забудь\nпро пробел, десу.\n" +
	       "\nЖивотные: %s.\n" + 
			    "\nПример ответа:\n" + "```%s```", org, reply)
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
	`,	//aardvarks
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
	`,	//beer
		"bear": `
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
	`,	//bison by cp97
		"bison": `
     \
      ((_,...,_))    
         |o o|
         \   /
          ^_^
		
	`,	//cock by sk 
		"cock": `
     \
      \  /""\      ,
        < ^  L____/|
         #) /'   , /
          \ '---' /
           ''";\)'
             _/_Y
		`,//duck buy hjw
		"duck": `
     \  __
      <(o )___
       ( ._> /
        '---' 
		`,//camel by jgs
		"camel": `
     \       _
      \  .--' |
        /___^ |     .--.
            ) |    /    \
           /  |  /'      '.
          |   '-'    /     \
          \         |      |\
           \    /   \      /\|
            \  /'----'\   /
            |||       \\ |
            ((|        ((|
            |||        |||
           //_(       //_(
	   	`,//deer
		"deer": `
    \
     \    ))    ((
      \  //      \\
         \\____// |
        ~/~    ~\/~
       ( o/  _/o  ~
        /  /     ,|
       (~~~)__.-\ |
        \'~~    | |
         |      | |
		`,//dog
		"dog":`
     \
      \    |\_/|            _    
       \   @ @ |         ((| |))
        \ <>   |           | |       
          /\_   \------____| |
           |                 |  
       ____|_       ___|    / 
      /_/_____/____/_______|
		`,//dolphin by Morfina
		"dolphin":`
     \ 
      \           ,-.
       \         /  (  '
        \  _.--'!   '--._
         ,'              ''.
   *    :                    \
      _.'  O      ___       ! '
     (_.-^, __..-'  ''''--.   |
         /,'        '    _.' /
     '         *    .-''    |
                   (..--^.  ' 
                         | /
                         '
		`,//elephant by Row
		"elephant":`
   \ /		
    :    ___     _,.--.,_
   /  .-~   ~--"~-.   ._ "-.
  :  /      ./_    Y    "-. \
  | Y       :~     !         Y
 /  l       |     /         .|
|_   \O  O-, l    /          |j
()\___) |/   \_/";          !
 \._____.-~\  .  ~\.      ./
            Y_ Y_. "vr"~  T
            (  (    |L    j 
            [nn[nn..][nn..]`}
	msg := ""
	arr := strings.Split(m.Text, " ")
    	animal := arr[0]
	if fruw[animal] != "" && strings.TrimSpace(m.Text) != animal{
		reply := fruw[animal]
		text := strings.TrimPrefix(m.Text, animal + " ")
		msg = fmt.Sprintf("``` < %s > %s ```", text, reply)
	} else { 
		rand.Seed(time.Now().UnixNano())
		rnd := (rand.Intn(lenmap))
		reply := fruw[rndmap[rnd]]
		msg = fmt.Sprintf("``` < %s > %s ```", m.Text, reply)
	}
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
