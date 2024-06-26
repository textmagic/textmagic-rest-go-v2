/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SenderSettingsItem struct {
	// The 2-letter ISO country code of the recipient's phone number. 
	Country string `json:"country"`
	// Phone enabled for sending to a specified country.
	Phone string `json:"phone"`
}
