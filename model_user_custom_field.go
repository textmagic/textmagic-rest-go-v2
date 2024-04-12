/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UserCustomField struct {
	// Custom field ID.
	Id int32 `json:"id"`
	// Custom field name.
	Name string `json:"name"`
	// Custom field creation time.
	CreatedAt string `json:"createdAt"`
}
