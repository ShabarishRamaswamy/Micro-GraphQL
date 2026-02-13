package handlers

import (
	"encoding/json"
	"fmt"
)

type QueryStruct struct {
	Query json.RawMessage `json:"query"`
}

func HandleIncomingQuery(query []byte) {
	msg := QueryStruct{}
	json.Unmarshal(query, &msg)

	fmt.Println(msg)
}
