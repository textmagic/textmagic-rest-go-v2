/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

// Response contains paginated list of data items.
type ResourceLinkResponse struct {
	// Resource ID.
	Id int32 `json:"id"`
	// A link to this resource. If you want to fetch it, just **GET** this address.
	Href string `json:"href"`
}
