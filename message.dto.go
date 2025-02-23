package main

type dtoMessage struct {
	PhoneNumberId string
	PhoneNumber   string
	ProfileName   string
	Message       Message
	List_Reply    string
}

func (dto *dtoMessage) createDto(msgData MessageInterface) {
	dto.PhoneNumberId = msgData.Entry[0].Changes[0].Value.Metadata.Phone_number_id
	dto.PhoneNumber = msgData.Entry[0].Changes[0].Value.Messages[0].From
	dto.Message = msgData.Entry[0].Changes[0].Value.Messages[0]
	dto.ProfileName = msgData.Entry[0].Changes[0].Value.Contacts[0].Profile.Name
	dto.List_Reply = msgData.Entry[0].Changes[0].Value.Messages[0].Interactive.List_Reply.Id
}

func getListReply() string {

	return ""
}
