# NotFoundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error code. Meanings of error codes are similar to [HTTP response codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). | [optional] 
**Message** | Pointer to **string** | A Brief error message. You could display this message to your user or save it in a log. | [optional] 

## Methods

### NewNotFoundResponse

`func NewNotFoundResponse() *NotFoundResponse`

NewNotFoundResponse instantiates a new NotFoundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundResponseWithDefaults

`func NewNotFoundResponseWithDefaults() *NotFoundResponse`

NewNotFoundResponseWithDefaults instantiates a new NotFoundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NotFoundResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *NotFoundResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *NotFoundResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotFoundResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotFoundResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotFoundResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


