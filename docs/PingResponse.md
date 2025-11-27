# PingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Current user Id. | 
**Ping** | **string** | Pong. | 
**UtcDateTime** | **string** | Current date and time. | 

## Methods

### NewPingResponse

`func NewPingResponse(userId int32, ping string, utcDateTime string, ) *PingResponse`

NewPingResponse instantiates a new PingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingResponseWithDefaults

`func NewPingResponseWithDefaults() *PingResponse`

NewPingResponseWithDefaults instantiates a new PingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PingResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PingResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PingResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetPing

`func (o *PingResponse) GetPing() string`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *PingResponse) GetPingOk() (*string, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *PingResponse) SetPing(v string)`

SetPing sets Ping field to given value.


### GetUtcDateTime

`func (o *PingResponse) GetUtcDateTime() string`

GetUtcDateTime returns the UtcDateTime field if non-nil, zero value otherwise.

### GetUtcDateTimeOk

`func (o *PingResponse) GetUtcDateTimeOk() (*string, bool)`

GetUtcDateTimeOk returns a tuple with the UtcDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcDateTime

`func (o *PingResponse) SetUtcDateTime(v string)`

SetUtcDateTime sets UtcDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


