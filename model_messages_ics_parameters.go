/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type MessagesIcsParameters struct {
	// Scheduled message text.
	Text string `json:"text"`
	Recipients *MessagesIcsParametersRecipients `json:"recipients"`
}
