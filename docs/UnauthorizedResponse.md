# UnauthorizedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error code. Meanings of error codes are similar to [HTTP response codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). | [optional] 
**Message** | Pointer to **string** | Brief error message. You could display this message to your user or save it in a log. | [optional] 

## Methods

### NewUnauthorizedResponse

`func NewUnauthorizedResponse() *UnauthorizedResponse`

NewUnauthorizedResponse instantiates a new UnauthorizedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedResponseWithDefaults

`func NewUnauthorizedResponseWithDefaults() *UnauthorizedResponse`

NewUnauthorizedResponseWithDefaults instantiates a new UnauthorizedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnauthorizedResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnauthorizedResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UnauthorizedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnauthorizedResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


