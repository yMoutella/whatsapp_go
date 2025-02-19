package main

type dtoMessage struct {
	PhoneNumberId string      `json:"phoneNumberId"`
	PhoneNumber   string      `json:"phoneNumber"`
	ProfileName   string      `json:"profileName"`
	Message       Message     `json:"message"`
	List_Reply    Interactive `json:"list_reply"`
}

func (dto *dtoMessage) createDto(msgData MessageInterface) {
	dto.PhoneNumberId = msgData.Entry[0].Changes[0].Value.Metadata.Phone_number_id
	dto.PhoneNumber = msgData.Entry[0].Changes[0].Value.Messages[0].From
	dto.Message = msgData.Entry[0].Changes[0].Value.Messages[0]
	dto.List_Reply = msgData.Entry[0].Changes[0].Value.Messages[0].Interactive
	dto.ProfileName = msgData.Entry[0].Changes[0].Value.Contacts[0].Profile.Name
}
