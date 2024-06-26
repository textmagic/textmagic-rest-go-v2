/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UserStatement struct {
	// User statement ID.
	Id int32 `json:"id"`
	// User ID.
	UserId int32 `json:"userId"`
	// User statement date.
	Date string `json:"date"`
	Balance float64 `json:"balance"`
	// Balance change amount.
	Delta float32 `json:"delta"`
	// Type of statement (what you have been charged for): *   **sms** - for sending SMS *   **number** - for renewing a dedicated number; *   **schedule** - for scheduling text messages; *   **topup** - for adding credits to your account. 
	Type_ string `json:"type"`
	// Value differs by **type**: *   for **sms**, it is the sent messages amount; *   for **number**, it is a dedicated phone number; *   for **schedule**, it is a scheduled messages amount; *   for **top-up**, it is an invoice ID. 
	Value string `json:"value"`
	// Optional comment.
	Comment string `json:"comment"`
}
