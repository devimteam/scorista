package main

import (
	"github.com/l-vitaly/scorista"
	"log"
	"os"
)

func main() {
	s := scorista.New("login", "secret")

	resp, err := s.CreditDecision(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(resp.Data), resp.Error.Message, resp.Status == scorista.ST_DONE)
}
