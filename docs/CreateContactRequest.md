# CreateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | Contact first name. | [optional] 
**LastName** | Pointer to **string** | Contact last name. | [optional] 
**Phone** | Pointer to **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | [optional] 
**Email** | Pointer to **string** | Contact email address. | [optional] 
**CompanyName** | Pointer to **string** | Company name. | [optional] 
**Lists** | Pointer to **string** | Contact [list](https://docs.textmagic.com/#tag/Lists) ID. Each contact must be assigned to at least one list. | [optional] 
**Favorited** | Pointer to **bool** | Is the contact marked as favorite? | [optional] 
**Blocked** | Pointer to **bool** | Is the contact blocked for outgoing and incoming messaging? | [optional] 
**Type** | Pointer to **int32** | Force type of phone. Possible values: 0 is landline; 1 is mobile; default is -1 (auto-detection). | [optional] 
**CustomFieldValues** | Pointer to [**[]CustomFieldListItem**](CustomFieldListItem.md) |  | [optional] 
**Local** | Pointer to **int32** | Treat phone numbers passed in the request body as local. | [optional] 
**Country** | Pointer to **string** | The 2-letter ISO country code for local phone numbers, used when local is  set to true. Default is the account country. | [optional] 
**Tags** | Pointer to **[]int32** | An array of tag IDs that will be assigned to the contact. | [optional] 
**Owner** | Pointer to **int32** | Contact Owner ID | [optional] 

## Methods

### NewCreateContactRequest

`func NewCreateContactRequest() *CreateContactRequest`

NewCreateContactRequest instantiates a new CreateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactRequestWithDefaults

`func NewCreateContactRequestWithDefaults() *CreateContactRequest`

NewCreateContactRequestWithDefaults instantiates a new CreateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CreateContactRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateContactRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateContactRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateContactRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CreateContactRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateContactRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateContactRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateContactRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *CreateContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *CreateContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompanyName

`func (o *CreateContactRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateContactRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateContactRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreateContactRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetLists

`func (o *CreateContactRequest) GetLists() string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *CreateContactRequest) GetListsOk() (*string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *CreateContactRequest) SetLists(v string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *CreateContactRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetFavorited

`func (o *CreateContactRequest) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *CreateContactRequest) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *CreateContactRequest) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *CreateContactRequest) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetBlocked

`func (o *CreateContactRequest) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *CreateContactRequest) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *CreateContactRequest) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *CreateContactRequest) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetType

`func (o *CreateContactRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateContactRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateContactRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateContactRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *CreateContactRequest) GetCustomFieldValues() []CustomFieldListItem`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *CreateContactRequest) GetCustomFieldValuesOk() (*[]CustomFieldListItem, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *CreateContactRequest) SetCustomFieldValues(v []CustomFieldListItem)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *CreateContactRequest) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetLocal

`func (o *CreateContactRequest) GetLocal() int32`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *CreateContactRequest) GetLocalOk() (*int32, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *CreateContactRequest) SetLocal(v int32)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *CreateContactRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetCountry

`func (o *CreateContactRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateContactRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateContactRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateContactRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTags

`func (o *CreateContactRequest) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContactRequest) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContactRequest) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *CreateContactRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateContactRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateContactRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateContactRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


