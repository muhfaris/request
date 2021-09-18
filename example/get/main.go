package main

import (
	"log"

	"github.com/muhfaris/request"
)

type QuoteModel struct {
	Anime     string
	Character string
	Quote     string
}

func getWithParse() {
	var quoteModel QuoteModel
	config := &request.Config{URL: "https://animechan.vercel.app/api/random"}
	response := request.Get(config).Parse(&quoteModel)
	if response.Error != nil {
		log.Printf("error get quote anime, %v", response.Error)
		return
	}

	log.Println("Quote:>")
	log.Println(quoteModel.Anime)
	log.Println(quoteModel.Character)
	log.Println(quoteModel.Quote)

}

func getWithoutParse() {
	config := &request.Config{URL: "https://animechan.vercel.app/api/random"}
	response := request.Get(config)
	if response.Error != nil {
		log.Printf("error get quote anime, %v", response.Error)
		return
	}

	log.Println("Response:")
	log.Println(string(response.Body))
}

func main() {
	getWithParse()
	getWithoutParse()
}
