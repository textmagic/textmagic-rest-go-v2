# GetContactsAutocompleteResponseItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **int32** | Id of entity. 0 if object is a Reply | [default to null]
**EntityType** | **string** | Entry type: * **contact** if it is related to a contact * **list** if it is related to a contact list * **reply** if it is related to an incoming message  | [default to null]
**Value** | **string** | Id of contact/list if entityType is contact/list OR phone number if entityType is reply. | [default to null]
**Label** | **string** | Name of the contact/list if entityType is contact/list OR phone number if entityType is reply. | [default to null]
**SharedBy** | **string** | If contact or list was shared by another sub-account then name if this user will be shown. | [default to null]
**Avatar** | **string** | Contact avatar URI. | [default to null]
**Favorited** | **bool** | If contact has been marked as favorite. | [default to null]
**UserId** | **int32** | Owner id of the contact/list (if it was shared). | [default to null]
**CountryName** | **string** |  | [default to null]
**Qposition** | **int32** |  | [default to null]
**Rposition** | **int32** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


