# RequestNewSubaccountTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | Sub-account ID. | [optional] 
**Password** | Pointer to **string** | Your account password. | [optional] 
**AppName** | Pointer to **string** | Application name. | [optional] 

## Methods

### NewRequestNewSubaccountTokenRequest

`func NewRequestNewSubaccountTokenRequest() *RequestNewSubaccountTokenRequest`

NewRequestNewSubaccountTokenRequest instantiates a new RequestNewSubaccountTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestNewSubaccountTokenRequestWithDefaults

`func NewRequestNewSubaccountTokenRequestWithDefaults() *RequestNewSubaccountTokenRequest`

NewRequestNewSubaccountTokenRequestWithDefaults instantiates a new RequestNewSubaccountTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RequestNewSubaccountTokenRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestNewSubaccountTokenRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestNewSubaccountTokenRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RequestNewSubaccountTokenRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPassword

`func (o *RequestNewSubaccountTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestNewSubaccountTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestNewSubaccountTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestNewSubaccountTokenRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAppName

`func (o *RequestNewSubaccountTokenRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *RequestNewSubaccountTokenRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *RequestNewSubaccountTokenRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *RequestNewSubaccountTokenRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


