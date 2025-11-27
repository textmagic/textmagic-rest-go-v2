# UsersInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Dedicated number ID. | 
**DisplayTimeFormat** | Pointer to **string** | Format for representation of time. | [optional] 
**Phone** | Pointer to **string** | Dedicated phone number. | [optional] 
**User** | [**NullableUser**](User.md) |  | 
**PurchasedAt** | **time.Time** | Time when the dedicated number was purchased. | 
**ExpireAt** | **time.Time** | Dedicated number subscription expiration time. | 
**Status** | **string** | Number status: *   **U** for Unused. No messages have been sent from (or received to) this number; *   **A** for Active.  | 
**Country** | [**NullableCountry**](Country.md) |  | 

## Methods

### NewUsersInbound

`func NewUsersInbound(id int32, user NullableUser, purchasedAt time.Time, expireAt time.Time, status string, country NullableCountry, ) *UsersInbound`

NewUsersInbound instantiates a new UsersInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersInboundWithDefaults

`func NewUsersInboundWithDefaults() *UsersInbound`

NewUsersInboundWithDefaults instantiates a new UsersInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersInbound) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersInbound) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersInbound) SetId(v int32)`

SetId sets Id field to given value.


### GetDisplayTimeFormat

`func (o *UsersInbound) GetDisplayTimeFormat() string`

GetDisplayTimeFormat returns the DisplayTimeFormat field if non-nil, zero value otherwise.

### GetDisplayTimeFormatOk

`func (o *UsersInbound) GetDisplayTimeFormatOk() (*string, bool)`

GetDisplayTimeFormatOk returns a tuple with the DisplayTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimeFormat

`func (o *UsersInbound) SetDisplayTimeFormat(v string)`

SetDisplayTimeFormat sets DisplayTimeFormat field to given value.

### HasDisplayTimeFormat

`func (o *UsersInbound) HasDisplayTimeFormat() bool`

HasDisplayTimeFormat returns a boolean if a field has been set.

### GetPhone

`func (o *UsersInbound) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UsersInbound) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UsersInbound) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UsersInbound) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUser

`func (o *UsersInbound) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsersInbound) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsersInbound) SetUser(v User)`

SetUser sets User field to given value.


### SetUserNil

`func (o *UsersInbound) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UsersInbound) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPurchasedAt

`func (o *UsersInbound) GetPurchasedAt() time.Time`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *UsersInbound) GetPurchasedAtOk() (*time.Time, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *UsersInbound) SetPurchasedAt(v time.Time)`

SetPurchasedAt sets PurchasedAt field to given value.


### GetExpireAt

`func (o *UsersInbound) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *UsersInbound) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *UsersInbound) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.


### GetStatus

`func (o *UsersInbound) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersInbound) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersInbound) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCountry

`func (o *UsersInbound) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UsersInbound) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UsersInbound) SetCountry(v Country)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *UsersInbound) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *UsersInbound) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


