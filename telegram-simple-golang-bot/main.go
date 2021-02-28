package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var chat_id int64
	chat_id = 123444444444444 //important ! changeto your receiver chat_id
	trd := sendMessage("Yes i'am talking to you!", chat_id, "<put you bot key here>")
	fmt.Println("Message sent ? " + strconv.FormatBool(trd.Ok))
	fmt.Println("Sender : " + trd.Result.From.Username)
	fmt.Println("Date : " + time.Unix(trd.Result.Date, 0).String()) //converting unix time to local time
	fmt.Println("Message : " + trd.Result.Text)
	tud := getUpdates("<put you bot key here>", 1) // getting only the last message push
	for _, msg := range tud.Result {
		fmt.Println(msg.Message.From.FirstName + ":" + msg.Message.Text)

	}

}
func sendMessage(text string, chat_id int64, botkey string) *TelgramResponseData { // function to send simple text message using chat_id of receiver, you need to create a bot using
	Url := "https://api.telegram.org/bot" + botkey + "/sendmessage?" // BotFather bot in telegram to get a key for your bot
	params := "text=" + text + "&chat_id=" + strconv.FormatInt(chat_id, 10)
	httpResponse, err := http.Get(fmt.Sprintf(Url + params))
	if err != nil {
		log.Fatalln(err)
	}

	defer httpResponse.Body.Close()

	responsebody, _ := ioutil.ReadAll(httpResponse.Body)
	var telgramResponseData TelgramResponseData
	if err := json.Unmarshal(responsebody, &telgramResponseData); err != nil {
		log.Println()
	}
	return &telgramResponseData
}
func getUpdates(botkey string, limit int) *TelegramUpdatesData { //this function gets all messages sent to the bot by other users,
	Url := "https://api.telegram.org/bot" + botkey + "/getUpdates?limit=" + strconv.Itoa(limit) //Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.

	httpResponse, err := http.Get(fmt.Sprintf(Url))
	if err != nil {
		log.Fatalln(err)
	}

	defer httpResponse.Body.Close()

	responsebody, _ := ioutil.ReadAll(httpResponse.Body)
	var telegramUpdatesData TelegramUpdatesData
	if err := json.Unmarshal(responsebody, &telegramUpdatesData); err != nil { // converting json data to TelgramUpdatesData struct
		log.Println()
	}
	return &telegramUpdatesData
}

type TelgramResponseData struct { // struct representing the data received after sending a message
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int64  `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

type TelegramUpdatesData struct { //struct representing the messages received by the bot
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID           int    `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LastName     string `json:"last_name"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}
