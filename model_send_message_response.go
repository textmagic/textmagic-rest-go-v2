/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SendMessageResponse struct {
	// Message ID.
	Id int32 `json:"id"`
	// URI of the message session.
	Href string `json:"href"`
	// Message response type: * **message** – when the message is sent to a single recipient. * **session** – when the message is sent is to multiple recipients. * **schedule** - when the message is scheduled for sending. * **bulk** - when the message is sent to multiple recipients and the number of recipients requires asynchronous processing See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session). 
	Type_ string `json:"type"`
	// Message session ID.
	SessionId int32 `json:"sessionId"`
	// Bulk Session ID. See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session).
	BulkId int32 `json:"bulkId"`
	// Message ID.
	MessageId int32 `json:"messageId"`
	// Message Schedule ID.
	ScheduleId int32 `json:"scheduleId"`
	// Message Chat ID.
	ChatId int32 `json:"chatId"`
}
