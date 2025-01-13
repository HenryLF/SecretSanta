package drafter

import (
	"log"

	"github.com/gorilla/websocket"
)

type Resquest struct {
	Gifters  []Person
	Unstrict bool
}

func DrafterWSApi(conn *websocket.Conn) {
	var req Resquest
	var err error
	for {
		err = conn.ReadJSON(&req)
		if err != nil {
			log.Println("Error reading JSON.", err)
			break
		}
		result, err := Process(req)
		if err != nil {
			log.Println("Processing error.", err)
			continue
		}
		err = conn.WriteJSON(result.toJSON())
		if err != nil {
			log.Println("Error writing JSON.", err)
			continue
		}
	}
	defer conn.Close()
}
