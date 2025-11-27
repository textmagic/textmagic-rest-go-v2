# UserStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User statement ID. | 
**UserId** | **NullableInt32** | User ID. | 
**Date** | **time.Time** | User statement date. | 
**Balance** | **NullableFloat64** |  | 
**Delta** | **NullableFloat32** | Balance change amount. | 
**Type** | **string** | Type of statement (what you have been charged for): *   **sms** - for sending SMS *   **number** - for renewing a dedicated number; *   **schedule** - for scheduling text messages; *   **topup** - for adding credits to your account.  | 
**Value** | **NullableString** | Value differs by **type**: *   for **sms**, it is the sent messages amount; *   for **number**, it is a dedicated phone number; *   for **schedule**, it is a scheduled messages amount; *   for **top-up**, it is an invoice ID.  | 
**Comment** | **NullableString** | Optional comment. | 

## Methods

### NewUserStatement

`func NewUserStatement(id int32, userId NullableInt32, date time.Time, balance NullableFloat64, delta NullableFloat32, type_ string, value NullableString, comment NullableString, ) *UserStatement`

NewUserStatement instantiates a new UserStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatementWithDefaults

`func NewUserStatementWithDefaults() *UserStatement`

NewUserStatementWithDefaults instantiates a new UserStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserStatement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserStatement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserStatement) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserStatement) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserStatement) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserStatement) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *UserStatement) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserStatement) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDate

`func (o *UserStatement) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UserStatement) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UserStatement) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetBalance

`func (o *UserStatement) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *UserStatement) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *UserStatement) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *UserStatement) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *UserStatement) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetDelta

`func (o *UserStatement) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *UserStatement) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *UserStatement) SetDelta(v float32)`

SetDelta sets Delta field to given value.


### SetDeltaNil

`func (o *UserStatement) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *UserStatement) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetType

`func (o *UserStatement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserStatement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserStatement) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UserStatement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserStatement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserStatement) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *UserStatement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UserStatement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetComment

`func (o *UserStatement) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UserStatement) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UserStatement) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *UserStatement) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UserStatement) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


