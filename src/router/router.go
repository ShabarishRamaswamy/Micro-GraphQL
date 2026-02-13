package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.HandleFunc("/sendQuery", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			type byteArr []byte

			body, err := io.ReadAll(r.Body)
			fmt.Printf("%+v, %+v\n\n", string(body), err)

			HandleIncomingQuery(body)
		}
	})

	return router
}

type QueryStruct struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func HandleIncomingQuery(query []byte) {
	msg := QueryStruct{}
	json.Unmarshal(query, &msg)

	// player, valid := msg.Query["player"].(string)
	// if !valid {
	// 	fmt.Println("This does not seem too valid")
	// 	return
	// }

	// fmt.Println(player)
	// fmt.Printf("%+v", player)
	ParseGQL(msg.Query)
}


```
We need to parse things like this:
{
  team {
    player {
      playerName
    }
    location
  }
}

team -> [player, location]
player -> [playerName]
```

func ParseGQL(query string) {
	// {"query":"{\n  team {\n    player {\n      playerName\n    }\n    location\n  }\n}\n","variables":{}}, <nil>
	// fmt.Println("We are gonna parse this: ", query)

	// var queryInsideQuery string

	// json.Unmarshal([]byte(query), &queryInsideQuery)

	// fmt.Println("Query: ", queryInsideQuery)
	for i := 0; i < len(query); i++ {
		fmt.Printf("%c", query[i])

		if query[i] == '{' {
			fmt.Println("Start of Something")
		}

		if query[i] == "}" {

		}

	}

}
