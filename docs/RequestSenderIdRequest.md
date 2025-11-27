# RequestSenderIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderId** | Pointer to **string** | The Sender ID that you are applying for. *   11 characters maximum; *   Only Latin based characters and digits are allowed; *   Should contain at least 1 letter.  | [optional] 
**Explanation** | Pointer to **string** | Explanation of why you need this Sender ID. | [optional] 

## Methods

### NewRequestSenderIdRequest

`func NewRequestSenderIdRequest() *RequestSenderIdRequest`

NewRequestSenderIdRequest instantiates a new RequestSenderIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSenderIdRequestWithDefaults

`func NewRequestSenderIdRequestWithDefaults() *RequestSenderIdRequest`

NewRequestSenderIdRequestWithDefaults instantiates a new RequestSenderIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderId

`func (o *RequestSenderIdRequest) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *RequestSenderIdRequest) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *RequestSenderIdRequest) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *RequestSenderIdRequest) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetExplanation

`func (o *RequestSenderIdRequest) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *RequestSenderIdRequest) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *RequestSenderIdRequest) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *RequestSenderIdRequest) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


