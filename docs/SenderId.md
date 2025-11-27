# SenderId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Numeric sender ID. | 
**DisplayTimeFormat** | Pointer to **string** | Format for representation of time. | [optional] 
**SenderId** | **string** | Alphanumeric ID. | 
**User** | [**NullableUser**](User.md) |  | 
**Status** | **string** | *   **P** for Pending - this Sender ID is being reviewed by our support team; *   **R** for Rejected - our support team rejected your application for this Sender ID; *   **A** for Active.  | 

## Methods

### NewSenderId

`func NewSenderId(id int32, senderId string, user NullableUser, status string, ) *SenderId`

NewSenderId instantiates a new SenderId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderIdWithDefaults

`func NewSenderIdWithDefaults() *SenderId`

NewSenderIdWithDefaults instantiates a new SenderId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SenderId) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SenderId) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SenderId) SetId(v int32)`

SetId sets Id field to given value.


### GetDisplayTimeFormat

`func (o *SenderId) GetDisplayTimeFormat() string`

GetDisplayTimeFormat returns the DisplayTimeFormat field if non-nil, zero value otherwise.

### GetDisplayTimeFormatOk

`func (o *SenderId) GetDisplayTimeFormatOk() (*string, bool)`

GetDisplayTimeFormatOk returns a tuple with the DisplayTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimeFormat

`func (o *SenderId) SetDisplayTimeFormat(v string)`

SetDisplayTimeFormat sets DisplayTimeFormat field to given value.

### HasDisplayTimeFormat

`func (o *SenderId) HasDisplayTimeFormat() bool`

HasDisplayTimeFormat returns a boolean if a field has been set.

### GetSenderId

`func (o *SenderId) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SenderId) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SenderId) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetUser

`func (o *SenderId) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SenderId) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SenderId) SetUser(v User)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SenderId) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SenderId) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetStatus

`func (o *SenderId) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SenderId) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SenderId) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


