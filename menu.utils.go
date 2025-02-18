package main

var menu map[int]string

func getMenu() map[int]string {
	menu[1] = "{\"messaging_product\":\"whatsapp\",\"to\":\"dto.phoneNumber\",\"type\":\"interactive\",\"interactive\":{\"type\":\"list\",\"header\":{\"type\":\"text\",\"text\":\"Menu inicial\"},\"body\":{\"text\":\"Nos ajude a entender como podemos te atender melhor:\"},\"action\":{\"button\":\"Menu\",\"sections\":[{\"title\":\"title\",\"rows\":[{\"id\":\"2\",\"title\":\"Dúvidas operacionais\"},{\"id\":\"3\",\"title\":\"Simulação\"},{\"id\":\"100\",\"title\":\"Sair\"}]}]}}}"

	menu[2] = "{\"messaging_product\":\"whatsapp\",\"to\":\"dto.phoneNumber\",\"type\":\"interactive\",\"interactive\":{\"type\":\"list\",\"header\":{\"type\":\"text\",\"text\":\"Dúvidas operacionais\"},\"body\":{\"text\":\"Nos ajude a entender como podemos te atender melhor:\"},\"action\":{\"button\":\"Menu\",\"sections\":[{\"title\":\"title\",\"rows\":[{\"id\":\"1\",\"title\":\"INSS\"},{\"id\":\"2\",\"title\":\"FGTS\"},{\"id\":\"3\",\"title\":\"Convênios Públicos\"},{\"id\":\"4\",\"title\":\"Crédito pessoal\"},{\"id\":\"5\",\"title\":\"Sistemas\"},{\"id\":\"99\",\"title\":\"Voltar\"},{\"id\":\"100\",\"title\":\"Sair\"}]}]}}}"

	menu[3] = "{\"messaging_product\":\"whatsapp\",\"to\":\"dto.phoneNumber\",\"type\":\"interactive\",\"interactive\":{\"type\":\"list\",\"header\":{\"type\":\"text\",\"text\":\"Simulação\"},\"body\":{\"text\":\"Nos ajude a entender como podemos te atender melhor:\"},\"action\":{\"button\":\"Menu\",\"sections\":[{\"rows\":[{\"id\":\"1\",\"title\":\"Santander\"},{\"id\":\"2\",\"title\":\"PAN\"},{\"id\":\"99\",\"title\":\"Voltar\"},{\"id\":\"100\",\"title\":\"Sair\"}]}]}}}"

	menu[100] = "{\"messaging_product\":\"whatsapp\",\"to\":\"dto.phoneNumber\",\"type\":\"text\",\"text\":{\"body\":\"Obrigado por utilizar nossos serviços\"}}"

	return menu
}
