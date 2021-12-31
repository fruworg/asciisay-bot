package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "\n*Привет!* Присылай сообщение и я его переведу.\n\nДля обычного текста:\n*en**ru*Hello, World!" +
		"\nДля текста на фотографии:\n*en**ru*https://website.com/file\n\nГде *en* - это язык текста, " +
		"а *ru* - язык перевода.\nДля *OCR* без перевода используй одинаковые языки.\n\nКоды языков:\n*ru* - Русский; " +
		"*en* - Английский; *ar* - Арабский; *zh* - Китайский; \n*fr* - Французкий; *de* - Немецкий; *hi* - Индийский; " +
		"\n*id* - Индонезийский; *ga* - Ирландский; *it* - Итальянский; \n*ja* - Японский; *ko* - Корейский; " +
		"*pl* - Польский; *pt* - Португальский; \n*es* - Испанский; *tr* - Турецкий; *vi* - Вьетнамский."
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
	cow = fmt.Sprintf("```\n%s\n```", cow)
	msg := fmt.Sprintf("%s\n%s\n%s%s", topLine, textLine, bottomLine, cow)
	a.client.SendMessage(m.Chat.ID, msg, tbot.OptParseModeMarkdown)
}
