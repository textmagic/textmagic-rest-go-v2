/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetAvailableDedicatedNumbersResponse struct {
	// Array of phone numbers.
	Numbers []string `json:"numbers"`
	// Dedicated number monthly fee for this country. Returned in the current [account](https://docs.textmagic.com/#tag/User) currency.
	Price float32 `json:"price"`
}
