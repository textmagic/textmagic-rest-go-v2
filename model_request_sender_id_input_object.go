/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type RequestSenderIdInputObject struct {
	// The Sender ID that you are applying for. *   11 characters maximum; *   Only Latin based characters and digits are allowed; *   Should contain at least 1 letter. 
	SenderId string `json:"senderId,omitempty"`
	// Explanation of why you need this Sender ID.
	Explanation string `json:"explanation,omitempty"`
}
