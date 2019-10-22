# UpdateContactInputObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Contact first name | [optional] [default to null]
**LastName** | **string** | Contact last name | [optional] [default to null]
**Phone** | **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | [default to null]
**Email** | **string** | Contact email address. | [optional] [default to null]
**CompanyName** | **string** | Contact company name | [optional] [default to null]
**Lists** | **string** | Comma-separated [list](http://docs.textmagictesting.com/#section/Lists) ID. Each contact must be assigned to at least one list. | [default to null]
**Favorited** | **bool** | Is contact marked as favorite. | [optional] [default to null]
**Blocked** | **bool** | Is contact blocked for outgoing and incoming messaging. | [optional] [default to null]
**Type_** | **int32** | Force type of phone. Possible values: 0 - landline, 1 - mobile. Default is -1 (auto detection). | [optional] [default to null]
**CustomFieldValues** | [**[]CustomFieldListItem**](CustomFieldListItem.md) |  | [optional] [default to null]
**Local** | **int32** | Treat phone number passed in request body as **local**. | [optional] [default to null]
**Country** | **string** | 2-letter ISO country code for local phone numbers, used when **local** is set to true. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


