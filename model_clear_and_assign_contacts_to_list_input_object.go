/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type ClearAndAssignContactsToListInputObject struct {
	// Contact ID(s), separated by a comma or \"all\" to add all contacts belonging to the current user.
	Contacts string `json:"contacts,omitempty"`
}
