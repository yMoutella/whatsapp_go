package main

type MessageInterface struct {
	Entry []Entry `json: "entry"`
}

type Entry struct {
	Changes []Change `json: "changes"`
}

type Change struct {
	Field string `json: "field"`
	Value Value  `json: "value"`
}

type Value struct {
	Messaging_product string    `json: "messaging_product"`
	Metadata          Metadata  `json: "metadata"`
	Contacts          []Contact `json: "contacts"`
	Messages          []Message `json: "messages"`
}

type Metadata struct {
	Display_phone_number string `json: "display_phone_number`
	Phone_number_id      string `json: "phone_number_id`
}

type Message struct {
	From        string      `json: "from"`
	Id          string      `json: "id"`
	Timestamp   string      `json: "timestamp"`
	Type        string      `json: "type"`
	Text        Txt         `json: "text"`
	Interactive Interactive `json: "interactive"`
}

type Txt struct {
	Body string `json: "body"`
}

type Contact struct {
	Profile Profile `json: "profile"`
	Wa_id   string  `json: "wa_id"`
}

type Profile struct {
	Name string `json: "name"`
}

type Interactive struct {
	Type       string     `json: "type"`
	List_Reply List_Reply `json: "list_reply"`
}

type List_Reply struct {
	Id    string `json: "id"`
	Title string `json: "title"`
}
