package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	CHANNEL_SECRET = "5d7381ad6a95bde9fa069f3783078327"
	CHANNEL_TOKEN  = "SaCwQG7BYVMGVMzRFVZDml2p8dBfUEZ/KRaHaNc0+2LhyM11Q4IeQZoZOKS1unZarPu2PxpIJJ2BAnFd1k8wQIzXqwtSUaCCTBoQhkHnjmptFD+M5hW7Zw7Zl3bNhOIGgjpssCW/6dj1WGvN56wfrwdB04t89/1O/w1cDnyilFU="
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	http.HandleFunc("/", Handler)

	address, err := determineListenAddress()
	if err != nil {
		log.Fatal("$PORT not set")
	}

	log.Fatal(http.ListenAndServe(address, nil))
}

type Message struct {
	events      []Event
	destination string
}

type Event struct {
	replyToken string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := linebot.New(CHANNEL_SECRET, CHANNEL_TOKEN)
	if err != nil {
		log.Fatal("Can't create bot")
	}

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var message Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(message)

	// bot.ReplyMessage("", linebot.NewTextMessage("Hello").Do(); err != nil {

	// })
}
