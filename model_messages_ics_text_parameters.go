/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type MessagesIcsTextParameters struct {
	// Cost to check that one number is constant – 0.04 in your account currency.
	Cost float32 `json:"cost"`
	// Message parts (multiples of 160 characters) count.
	Parts int32 `json:"parts"`
	// Characters count.
	Chars int32 `json:"chars"`
	// Message charset. Could be: * **ISO-8859-1** – for plaintext SMS; * **UTF-16BE** – for Unicode SMS. 
	Encoding string `json:"encoding"`
	Countries []string `json:"countries"`
	// Human-readable message charset label. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS; *   **Voice** for voice services (Text-to-Speech or Voice Broadcast) messages. 
	CharsetLabel string `json:"charsetLabel"`
}
