# UpdateCurrentUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username. | [optional] 
**FirstName** | Pointer to **string** | Account first name. | [optional] 
**LastName** | Pointer to **string** | Account last name. | [optional] 
**Email** | Pointer to **string** | User email address. | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** | Account company name. | [optional] 
**Timezone** | Pointer to **int32** | Internal timezone ID. See [Get timezones](https://docs.textmagic.com/#operation/getTimezones). | [optional] 

## Methods

### NewUpdateCurrentUserResponse

`func NewUpdateCurrentUserResponse() *UpdateCurrentUserResponse`

NewUpdateCurrentUserResponse instantiates a new UpdateCurrentUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCurrentUserResponseWithDefaults

`func NewUpdateCurrentUserResponseWithDefaults() *UpdateCurrentUserResponse`

NewUpdateCurrentUserResponseWithDefaults instantiates a new UpdateCurrentUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateCurrentUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCurrentUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCurrentUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateCurrentUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateCurrentUserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateCurrentUserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateCurrentUserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateCurrentUserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateCurrentUserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateCurrentUserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateCurrentUserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateCurrentUserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCurrentUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCurrentUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCurrentUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCurrentUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateCurrentUserResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateCurrentUserResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateCurrentUserResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateCurrentUserResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UpdateCurrentUserResponse) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UpdateCurrentUserResponse) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCompany

`func (o *UpdateCurrentUserResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateCurrentUserResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateCurrentUserResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateCurrentUserResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdateCurrentUserResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdateCurrentUserResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetTimezone

`func (o *UpdateCurrentUserResponse) GetTimezone() int32`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateCurrentUserResponse) GetTimezoneOk() (*int32, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateCurrentUserResponse) SetTimezone(v int32)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateCurrentUserResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


