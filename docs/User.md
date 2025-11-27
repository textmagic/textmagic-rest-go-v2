# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User ID. | 
**DisplayTimeFormat** | Pointer to **string** | User&#39;s preferred format of time display: * *12h* - AM/PM format; * *24h* - 24-hour clock format.  | [optional] 
**Username** | **string** | Username. | 
**FirstName** | **string** | Account first name. | 
**LastName** | **string** | Account last name. | 
**Email** | **string** | User email address. | 
**Status** | **string** | Current account status: * **A** for Active; * **T** for Trial.  | 
**Balance** | **float32** | Account balance (in account currency). | 
**Phone** | **NullableString** | User&#39;s phone number. | 
**Company** | **NullableString** | Account company name. | 
**Currency** | [**Currency**](Currency.md) |  | 
**Country** | [**NullableCountry**](Country.md) |  | 
**Timezone** | [**Timezone**](Timezone.md) |  | 
**SubaccountType** | **string** | Type of account: * **P** for Parent User; * **A** for Administrator Sub-Account; * **U** for Regular User.  | 
**EmailAccepted** | **bool** | Does the account have a confirmed email? | 
**PhoneAccepted** | **bool** | Does the account have a confirmed phone number? | 
**Avatar** | [**NullableUserImage**](UserImage.md) |  | 

## Methods

### NewUser

`func NewUser(id int32, username string, firstName string, lastName string, email string, status string, balance float32, phone NullableString, company NullableString, currency Currency, country NullableCountry, timezone Timezone, subaccountType string, emailAccepted bool, phoneAccepted bool, avatar NullableUserImage, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetDisplayTimeFormat

`func (o *User) GetDisplayTimeFormat() string`

GetDisplayTimeFormat returns the DisplayTimeFormat field if non-nil, zero value otherwise.

### GetDisplayTimeFormatOk

`func (o *User) GetDisplayTimeFormatOk() (*string, bool)`

GetDisplayTimeFormatOk returns a tuple with the DisplayTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimeFormat

`func (o *User) SetDisplayTimeFormat(v string)`

SetDisplayTimeFormat sets DisplayTimeFormat field to given value.

### HasDisplayTimeFormat

`func (o *User) HasDisplayTimeFormat() bool`

HasDisplayTimeFormat returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBalance

`func (o *User) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *User) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *User) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *User) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *User) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCompany

`func (o *User) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *User) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *User) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *User) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *User) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCurrency

`func (o *User) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *User) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *User) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetCountry

`func (o *User) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v Country)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *User) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *User) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetTimezone

`func (o *User) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *User) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *User) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.


### GetSubaccountType

`func (o *User) GetSubaccountType() string`

GetSubaccountType returns the SubaccountType field if non-nil, zero value otherwise.

### GetSubaccountTypeOk

`func (o *User) GetSubaccountTypeOk() (*string, bool)`

GetSubaccountTypeOk returns a tuple with the SubaccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccountType

`func (o *User) SetSubaccountType(v string)`

SetSubaccountType sets SubaccountType field to given value.


### GetEmailAccepted

`func (o *User) GetEmailAccepted() bool`

GetEmailAccepted returns the EmailAccepted field if non-nil, zero value otherwise.

### GetEmailAcceptedOk

`func (o *User) GetEmailAcceptedOk() (*bool, bool)`

GetEmailAcceptedOk returns a tuple with the EmailAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAccepted

`func (o *User) SetEmailAccepted(v bool)`

SetEmailAccepted sets EmailAccepted field to given value.


### GetPhoneAccepted

`func (o *User) GetPhoneAccepted() bool`

GetPhoneAccepted returns the PhoneAccepted field if non-nil, zero value otherwise.

### GetPhoneAcceptedOk

`func (o *User) GetPhoneAcceptedOk() (*bool, bool)`

GetPhoneAcceptedOk returns a tuple with the PhoneAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneAccepted

`func (o *User) SetPhoneAccepted(v bool)`

SetPhoneAccepted sets PhoneAccepted field to given value.


### GetAvatar

`func (o *User) GetAvatar() UserImage`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*UserImage, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v UserImage)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *User) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *User) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


