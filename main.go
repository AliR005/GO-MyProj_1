package main

import (
	"flag"
	"log"
)

func main(){


}

func mustToken () string {
	token := flag.String("token_bot", "", "token for access tg bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	
	return *token
}