package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
)

func main() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Need BOT_TOKEN env var")
	}

	ownerIDStr, ok := os.LookupEnv("OWNER_ID")
	if !ok {
		log.Fatal("Need OWNER_ID env var")
	}
	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		log.Fatalf("OWNER_ID is not an integer")
	}

	//proxyClient := createSocks5ProxyClient()
	//bot, err := tgbotapi.NewBotAPIWithClient(token, proxyClient)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("can't create bot: %v", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] received %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(int64(ownerID), update.Message.Text)

		bot.Send(msg)
	}
}

//func createSocks5ProxyClient() *http.Client {
//	auth := &proxy.Auth{
//		User:     "tgproxy",
//		Password: "fuckrkn",
//	}
//	dialSocksProxy, err := proxy.SOCKS5("tcp", "tp.grishka.me:1080", auth, proxy.Direct)
//	if err != nil {
//		fmt.Println("Error connecting to proxy:", err)
//	}
//	tr := &http.Transport{Dial: dialSocksProxy.Dial}
//
//	// Create client
//	return &http.Client{
//		Transport: tr,
//	}
//}
