func main() {

	fmt.Println("Starting server on :4446")

	http.HandleFunc("/webhook", Verify)
	log.Fatal(http.ListenAndServe(":4446", nil))

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
			rp := strings.NewReplacer("<at>", "", "chatbot", "", "</at>", "", "search", "", "&", "", "nbsp", "", ";", "")
			output := rp.Replace(ti)
			fmt.Println("***output****", output)

			resp, resp1 := Devopsresp(*input, output)

			caction := new(Actioncard)
			input := new(Message)

			caction.Type = "MessageCard"
			caction.ThemeColor = "0072C6"
			caction.Title = "Resutls - If nothing returns try diffrent search term"

			//caction.Text = resp
			caction.Text = resp
			input.Text = resp1
			input.ReplyToID = input.From.ID

			respo, _ := json.Marshal(caction)
			respo1, _ := json.Marshal(input)

			http.Post(url1, "application/json",
				bytes.NewReader(respo))

			http.Post(url1, "application/json",
				bytes.NewReader(respo1))
			return
		}

		caction := new(Actioncard)

		caction.Type = "MessageCard"
		caction.ThemeColor = "0072C6"
		caction.Title = "Welcome to the aks wiki searchbot bot"
		caction.Text = "Usage: search termsearch \n @chatbot search node issues"

		s, _ := json.Marshal(caction)
		//s := Ma()
		//fmt.Println(string(s))

		r, err := http.Post(url1, "application/json",
			bytes.NewReader(s))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r.Status, string(s))

		return

	}

}
