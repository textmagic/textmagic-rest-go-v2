/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type BuyDedicatedNumberInputObject struct {
	// Dedicated phone number.
	Phone string `json:"phone,omitempty"`
	// Country code phone number.
	Country string `json:"country,omitempty"`
	// Assigned dedicated number. This number will be available for this account only. You cannot transfer numbers between sub-accounts. 
	UserId int32 `json:"userId,omitempty"`
}
