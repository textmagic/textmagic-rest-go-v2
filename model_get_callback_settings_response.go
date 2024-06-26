/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetCallbackSettingsResponse struct {
	// This URL is used to push message delivery status updates to your application.
	OutUrl string `json:"outUrl"`
	// This URL is used to push incoming SMS to your application.
	InUrl string `json:"inUrl"`
	// Desired callback data format. m - multipart/form-data, u - application/x-www-form-urlencoded, j - application/json.
	Format string `json:"format"`
}
