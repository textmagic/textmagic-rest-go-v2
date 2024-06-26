/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetContactsAutocompleteResponseItem struct {
	// Id of entity. 0 if object is a reply.
	EntityId int32 `json:"entityId"`
	// Entry type: * **contact** if it is related to a contact; * **list** if it is related to a contact list; * **reply** if it is related to an incoming message. 
	EntityType string `json:"entityType"`
	// ID of the contact/list if entityType is contact/list OR phone number if entityType is reply.
	Value string `json:"value"`
	// Name of the contact/list if entityType is contact/list OR phone number if entityType is reply.
	Label string `json:"label"`
	// If contact or list was shared by another sub-account, the name of this user will be shown.
	SharedBy string `json:"sharedBy"`
	// If contact or list was shared by another sub-account then `true` will be set.
	IsShared bool `json:"isShared"`
	// Contact avatar URI.
	Avatar string `json:"avatar"`
	// If contact has been marked as favorite.
	Favorited bool `json:"favorited"`
	// Owner ID of the contact/list (if it was shared).
	UserId int32 `json:"userId"`
	CountryName string `json:"countryName"`
	Qposition int32 `json:"qposition"`
	Rposition int32 `json:"rposition"`
}
