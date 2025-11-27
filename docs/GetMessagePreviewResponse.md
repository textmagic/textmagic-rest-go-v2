# GetMessagePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MessageOut**](MessageOut.md) |  | [optional] 

## Methods

### NewGetMessagePreviewResponse

`func NewGetMessagePreviewResponse() *GetMessagePreviewResponse`

NewGetMessagePreviewResponse instantiates a new GetMessagePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessagePreviewResponseWithDefaults

`func NewGetMessagePreviewResponseWithDefaults() *GetMessagePreviewResponse`

NewGetMessagePreviewResponseWithDefaults instantiates a new GetMessagePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetMessagePreviewResponse) GetMessages() []MessageOut`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetMessagePreviewResponse) GetMessagesOk() (*[]MessageOut, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetMessagePreviewResponse) SetMessages(v []MessageOut)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetMessagePreviewResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


