package main

type MessageType interface {
	Message_List_Request | Message_Simple_Request
}

type Message_List_Request struct {
	Messaging_Product string `json:"messaging_product"`
	To                string `json:"to"`
	Type              string `json:"type"`
	Interactive       struct {
		Type   string `json:"type"`
		Header struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"header"`
		Body struct {
			Text string `json:"text"`
		} `json:"body"`
		Action struct {
			Button   string `json:"button"`
			Sections []struct {
				Title string `json:"title"`
				Rows  []struct {
					Id    string `json:"id"`
					Title string `json:"title"`
				} `json:"rows"`
			} `json:"sections"`
		} `json:"action"`
	} `json:"interactive"`
}

type Message_Simple_Request struct {
	Messaging_Product string `json:"messaging_product"`
	To                string `json:"to"`
	Type              string `json:"type"`
	Text              Text   `json:"text"`
}

type Text struct {
	Body string `json:"body"`
}
