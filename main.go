package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var helpMessage = "使い方:。\n選択肢数(2~9)\n選択肢１\n選択肢２\n...\n選択肢n"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(parse(message.Text))).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func parse(message string) string {
	rand.Seed(time.Now().UnixNano())
	if startsWithN(message) {
		if n, err := strconv.Atoi(message[0:0]); err == nil {
			if nm := strings.SplitN(message, "\n", n+1); nm != nil {
				ch := rand.Intn(n) + 1
				text := "乱数の女神さまの厳正な判断の元選ばれたのは\n" + nm[ch] + "\nでした。"
				return text
			} else {
				return "選択肢の数が合わないよ、改行区切りで最後は改行しないでね！\n" + helpMessage
			}
		}
	} else {
		return helpMessage
	}
}

func startsWithN(str string) bool {
	for i := 2; i < 10; i++ {
		if strings.HasPrefix(str, strconv.Itoa(i)) {
			return true
		}
	}
	return false
}
