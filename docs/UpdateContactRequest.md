# UpdateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | Contact first name. | [optional] 
**LastName** | Pointer to **string** | Contact last name. | [optional] 
**Phone** | Pointer to **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | [optional] 
**Email** | Pointer to **string** | Contact email address. | [optional] 
**CompanyName** | Pointer to **string** | Contact company name. | [optional] 
**Lists** | Pointer to **string** | Comma-separated [list](https://docs.textmagic.com/#section/Lists) ID. Each contact must be assigned to at least one list. | [optional] 
**Favorited** | Pointer to **bool** | Is the contact marked as favorite? | [optional] 
**Blocked** | Pointer to **bool** | Is the contact blocked for outgoing and incoming messaging? | [optional] 
**Type** | Pointer to **int32** | Force type of phone. Possible values: 0 is landline; 1 is mobile; default is -1 (auto-detection). | [optional] 
**CustomFieldValues** | Pointer to [**[]CustomFieldListItem**](CustomFieldListItem.md) |  | [optional] 
**Local** | Pointer to **int32** | Treat phone numbers passed in the request body as **local**. | [optional] 
**Country** | Pointer to **string** | The 2-letter ISO country code for local phone numbers, used when **local** is set to true. | [optional] 
**Tags** | Pointer to **[]int32** | An array of tag IDs that will be assigned to the contact. | [optional] 
**Owner** | Pointer to **int32** | Contact Owner ID | [optional] 

## Methods

### NewUpdateContactRequest

`func NewUpdateContactRequest() *UpdateContactRequest`

NewUpdateContactRequest instantiates a new UpdateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactRequestWithDefaults

`func NewUpdateContactRequestWithDefaults() *UpdateContactRequest`

NewUpdateContactRequestWithDefaults instantiates a new UpdateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateContactRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateContactRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateContactRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateContactRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateContactRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateContactRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateContactRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateContactRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompanyName

`func (o *UpdateContactRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateContactRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateContactRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateContactRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetLists

`func (o *UpdateContactRequest) GetLists() string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *UpdateContactRequest) GetListsOk() (*string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *UpdateContactRequest) SetLists(v string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *UpdateContactRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetFavorited

`func (o *UpdateContactRequest) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *UpdateContactRequest) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *UpdateContactRequest) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *UpdateContactRequest) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetBlocked

`func (o *UpdateContactRequest) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *UpdateContactRequest) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *UpdateContactRequest) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *UpdateContactRequest) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetType

`func (o *UpdateContactRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateContactRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateContactRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateContactRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *UpdateContactRequest) GetCustomFieldValues() []CustomFieldListItem`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *UpdateContactRequest) GetCustomFieldValuesOk() (*[]CustomFieldListItem, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *UpdateContactRequest) SetCustomFieldValues(v []CustomFieldListItem)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *UpdateContactRequest) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetLocal

`func (o *UpdateContactRequest) GetLocal() int32`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *UpdateContactRequest) GetLocalOk() (*int32, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *UpdateContactRequest) SetLocal(v int32)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *UpdateContactRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateContactRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateContactRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateContactRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateContactRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContactRequest) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContactRequest) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContactRequest) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateContactRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateContactRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateContactRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateContactRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


