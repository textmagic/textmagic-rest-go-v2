# BadRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error code. Meanings of error codes are similar to [HTTP response codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). | [optional] 
**Message** | Pointer to **string** | Brief error message. You could display this message to your user or save it in a log. | [optional] 
**Errors** | Pointer to [**NullableBadRequestResponseErrors**](BadRequestResponseErrors.md) |  | [optional] 

## Methods

### NewBadRequestResponse

`func NewBadRequestResponse() *BadRequestResponse`

NewBadRequestResponse instantiates a new BadRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestResponseWithDefaults

`func NewBadRequestResponseWithDefaults() *BadRequestResponse`

NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BadRequestResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *BadRequestResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *BadRequestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadRequestResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *BadRequestResponse) GetErrors() BadRequestResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BadRequestResponse) GetErrorsOk() (*BadRequestResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BadRequestResponse) SetErrors(v BadRequestResponseErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BadRequestResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BadRequestResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BadRequestResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


