/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type MessageOut struct {
	// Message ID.
	Id int32 `json:"id"`
	// Message sender (phone number or alphanumeric Sender ID).
	Sender string `json:"sender,omitempty"`
	// Recipient`s phone number.
	Receiver string `json:"receiver,omitempty"`
	Text string `json:"text"`
	// Delivery status of the message. See [message delivery statuses](https://docs.textmagic.com/#section/Delivery-status-codes) for details. 
	Status string `json:"status"`
	// Recipient contact ID.
	ContactId int32 `json:"contactId"`
	// Message Session ID of a message.
	SessionId int32 `json:"sessionId"`
	// Sending time.
	MessageTime string `json:"messageTime"`
	Avatar string `json:"avatar"`
	// Indicates that the message has been deleted.
	Deleted bool `json:"deleted,omitempty"`
	// Message charset. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS. 
	Charset string `json:"charset"`
	// Human-readable message charset label. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS; *   **Voice** for voice services (Text-to-Speech or Voice Broadcast) messages. 
	CharsetLabel string `json:"charsetLabel"`
	// Contact first name. Could be substituted from your [Contacts](https://docs.textmagic.com/#tag/Contacts) (even if you submitted the phone number instead of the contact ID). 
	FirstName string `json:"firstName"`
	// Contact last name.
	LastName string `json:"lastName"`
	// The 2-letter ISO country code of the recipient's phone number. 
	Country string `json:"country"`
	// Receipent`s phone number.
	Phone string `json:"phone,omitempty"`
	// Message price.
	Price float32 `json:"price,omitempty"`
	// Message parts (multiples of 160 characters) count.
	PartsCount int32 `json:"partsCount"`
	// The user email which this message came from. For Email2SMS and Distribution Lists the messages, it is an original email address - in other cases, it is an account email address.
	FromEmail string `json:"fromEmail,omitempty"`
	// The Phone number used to send the SMS.
	FromNumber string `json:"fromNumber,omitempty"`
}
