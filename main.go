package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var plexEvents = make(chan PlexMessage)
var token string

func main() {
	var ok bool
	token, ok = os.LookupEnv("BOT_TOKEN")
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
	httpPort, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("Need PORT env var")
	}


	//proxyClient := createSocks5ProxyClient()
	//bot, err := tgbotapi.NewBotAPIWithClient(token, proxyClient)
	go processPlexMessages(int64(ownerID), plexEvents)

	http.HandleFunc("/plex", plexWebhookHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil))
}

func processPlexMessages(ownerID int64, events chan PlexMessage) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("can't create bot: %v", err)
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	for plexEvent := range events {
		verb := getVerbByEventType(plexEvent.Event)
		fullTitle := strings.Join(rejectNils([]string{plexEvent.Metadata.GrandparentTitle, plexEvent.Metadata.ParentTitle, plexEvent.Metadata.Title}), " - ")
		msg := fmt.Sprintf("%s %s %s", plexEvent.Account.Title, verb, fullTitle)

		tmsg := tgbotapi.NewMessage(ownerID, msg)
		_, err := bot.Send(tmsg)
		if err != nil {
			log.Println(err)
		}

	}
}

func rejectNils(collection []string) []string {
	result := make([]string, 0, len(collection))
	for _, elem := range collection {
		if len(elem) > 0 {
			result = append(result, elem)
		}
	}
	return result
}

func getVerbByEventType(eventType string) string {
	switch eventType {
	case "media.play":
		return "has started playing"
	case "media.pause":
		return "has paused"
	case "media.resume":
		return "has resumed"
	case "media.stop":
		return "has stopped watching"
	case "media.scrobble":
		return "has watched"
	default:
		return "has " + eventType
	}
}

func plexWebhookHandler(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := mr.NextPart()

		// This is OK, no more parts
		if err == io.EOF {
			break
		}

		// Some error
		if err != nil {
			fmt.Println("error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON 'doc' part
		if part.FormName() == "payload" {
			payload := PlexMessage{}

			jsonDecoder := json.NewDecoder(part)
			err = jsonDecoder.Decode(&payload)
			if err != nil {
				fmt.Println("error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//fmt.Printf("plex %+v\n", payload)
			plexEvents <- payload
		}
	}
	w.WriteHeader(http.StatusOK)
}
