/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

// Returned when request has been denied in reason of access problems.
type UnauthorizedResponse struct {
	// Error code. Meanings of error codes are similar to [HTTP response codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes).
	Code int32 `json:"code,omitempty"`
	// Brief error message. You could display this message to your user or save it in a log.
	Message string `json:"message,omitempty"`
}
