# MessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Payload type. | 
**MediaPreview** | **string** | Media preview link. | 

## Methods

### NewMessagePayload

`func NewMessagePayload(type_ string, mediaPreview string, ) *MessagePayload`

NewMessagePayload instantiates a new MessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagePayloadWithDefaults

`func NewMessagePayloadWithDefaults() *MessagePayload`

NewMessagePayloadWithDefaults instantiates a new MessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessagePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessagePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessagePayload) SetType(v string)`

SetType sets Type field to given value.


### GetMediaPreview

`func (o *MessagePayload) GetMediaPreview() string`

GetMediaPreview returns the MediaPreview field if non-nil, zero value otherwise.

### GetMediaPreviewOk

`func (o *MessagePayload) GetMediaPreviewOk() (*string, bool)`

GetMediaPreviewOk returns a tuple with the MediaPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaPreview

`func (o *MessagePayload) SetMediaPreview(v string)`

SetMediaPreview sets MediaPreview field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


