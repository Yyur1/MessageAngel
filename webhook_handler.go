package main

import (
	"errors"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"net/http"
)

// WebhookHandler stores text message to the database
func WebhookHandler() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		cb, err := webhook.ParseRequest(channelSecret, request)
		if err != nil {
			logger.Warn("Cannot parse request: %+v", err)
			if errors.Is(err, webhook.ErrInvalidSignature) {
				writer.WriteHeader(400)
			} else {
				writer.WriteHeader(500)
			}
			return
		}

		for _, event := range cb.Events {
			switch e := event.(type) {
			case webhook.MessageEvent:
				if source, ok := e.Source.(*webhook.GroupSource); ok {
					if message, ok := e.Message.(*webhook.TextMessageContent); ok {
						logger.Infof("SENT: At Group: %+v, User: %+v SENT text message: %+v",
							source.GroupId, source.UserId, message)
						err := CreateGroupMessage(source.GroupId, source.UserId, message)
						if err != nil {
							logger.Errorf("Create: Cannot create message event: %+v", err)
						}
					}
				}
			case webhook.UnsendEvent:
				if source, ok := e.Source.(*webhook.GroupSource); ok {
					logger.Infof("UNSENT: At Group: %+v, User: %+v UNSENT message_id: %+v", source.GroupId, source.UserId, e.Unsend.MessageId)
					err := DeleteGroupMessage(source.GroupId, source.UserId, e.Unsend.MessageId)
					if err != nil {
						logger.Errorf("Delete: Cannot delete message event: %+v", err)
					}
				}
			}
		}
	}
}

func CreateGroupMessage(groupId string, userId string, message *webhook.TextMessageContent) error {
	newMessageRecord := MessageRecord{GroupId: groupId, UserId: userId, Message: message.Text, MessageId: message.Id}
	result := db.Create(&newMessageRecord)
	if result.Error != nil {
		logger.Errorf("Failed to insert. GroupId: %v, UserId: %v, Message: %v", groupId, userId, message)
		return result.Error
	}
	return nil
}

func DeleteGroupMessage(groupId string, userId string, messageId string) error {
	result := db.Where("group_id = ? AND user_id = ? AND message_id = ?", groupId, userId, messageId).Delete(&MessageRecord{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
