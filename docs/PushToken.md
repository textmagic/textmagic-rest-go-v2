# PushToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | type of the token: * **GCM** — Google Cloud Messaging * **APN** — Apple Push Notification * **FCM** — Firebase Cloud Messaging  | 
**Token** | **string** | Push token value. | 

## Methods

### NewPushToken

`func NewPushToken(type_ string, token string, ) *PushToken`

NewPushToken instantiates a new PushToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTokenWithDefaults

`func NewPushTokenWithDefaults() *PushToken`

NewPushTokenWithDefaults instantiates a new PushToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PushToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PushToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PushToken) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *PushToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushToken) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


