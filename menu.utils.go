package main

import "encoding/json"

func getGoodBye() Message_Simple_Request {

	return Message_Simple_Request{
		MessagingProduct: "whatsapp",
		To:               "",
		Type:             "text",
		Text: Text{
			Body: "Obrigado por utilizar nossos serviços",
		},
	}

}

func getMenu() (Message_List_Request, error) {
	menu := map[int]string{
		1:   `{"messaging_product":"whatsapp","to":"dto.phoneNumber","type":"interactive","interactive":{"type":"list","header":{"type":"text","text":"Menu inicial"},"body":{"text":"Nos ajude a entender como podemos te atender melhor:"},"action":{"button":"Menu","sections":[{"title":"title","rows":[{"id":"2","title":"Dúvidas operacionais"},{"id":"3","title":"Simulação"},{"id":"100","title":"Sair"}]}]}}}`,
		2:   `{"messaging_product":"whatsapp","to":"dto.phoneNumber","type":"interactive","interactive":{"type":"list","header":{"type":"text","text":"Dúvidas operacionais"},"body":{"text":"Nos ajude a entender como podemos te atender melhor:"},"action":{"button":"Menu","sections":[{"title":"title","rows":[{"id":"1","title":"INSS"},{"id":"2","title":"FGTS"},{"id":"3","title":"Convênios Públicos"},{"id":"4","title":"Crédito pessoal"},{"id":"5","title":"Sistemas"},{"id":"99","title":"Voltar"},{"id":"100","title":"Sair"}]}]}}}`,
		3:   `{"messaging_product":"whatsapp","to":"dto.phoneNumber","type":"interactive","interactive":{"type":"list","header":{"type":"text","text":"Simulação"},"body":{"text":"Nos ajude a entender como podemos te atender melhor:"},"action":{"button":"Menu","sections":[{"rows":[{"id":"1","title":"Santander"},{"id":"2","title":"PAN"},{"id":"99","title":"Voltar"},{"id":"100","title":"Sair"}]}]}}}`,
		100: `{"messaging_product":"whatsapp","to":"dto.phoneNumber","type":"text","text":{"body":"Obrigado por utilizar nossos serviços"}}`,
	}

	var message Message_List_Request

	err := json.Unmarshal([]byte(menu[1]), &message)

	return message, err
}
