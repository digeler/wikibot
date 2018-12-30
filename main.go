package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	azureDevOpsResourceID      = "499b84ac-1321-427f-aa17-267ca6975798"
	azureDevOpsOrganizationURL = "http://dev.azure.com/organization"
	url1                       = "https://outlook.office.com/webhook/074e4c99-14b9-4454-98ae-9eff23b77872@72f988bf-86f1-41af-91ab-2d7cd011db47"

	url = "https://almsearch.dev.azure.com/aks-support-demo/_apis/search/"
)

// Message declaration

func main() {

	fmt.Println("Starting server on :4445")
	http.HandleFunc("/webhook", Verify)
	log.Fatal(http.ListenAndServe(":4445", nil))

}

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		input := new(Message)
		if err := json.NewDecoder(r.Body).Decode(input); err == nil {
			//fmt.Println(input.Text, input.From, input.Recipient.Name)
		}
		if strings.Contains(input.Text, "search") {
			fmt.Println(input.Text)
			ti := input.Text
			rp := strings.NewReplacer("<at>", "", "wikibot", "", "</at>", "", "search", "", "&", "", "nbsp", "", ";", "")
			output := rp.Replace(ti)
			fmt.Println(output)

			resp := Devopsresp(*input, output)
			input.Text = resp
			input.ReplyToID = input.From.ID

			respo, _ := json.Marshal(input)

			http.Post(url1, "application/json",
				bytes.NewReader(respo))
			return
		}

		caction := new(Actioncard)
		caction.Type = "MessageCard"
		caction.ThemeColor = "0072C6"
		caction.Title = "Welcome to the aks wiki searchbot bot"
		caction.Text = "Usage: search termsearch \n @servicebot search node issues"
		s, _ := json.Marshal(caction)

		http.Post(url1, "application/json",
			bytes.NewReader(s))

		return

	}

}
