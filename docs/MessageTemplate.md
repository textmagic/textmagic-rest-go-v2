# MessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Template ID. | 
**Name** | **string** | Template name. | 
**Content** | **string** | Template text. May contain dynamic fields inside braces. See the [Custom fields list](https://docs.textmagic.com/#tag/Templates/Custom-fields-list-(Merge-dynamic-fields)). | 
**LastModified** | **NullableTime** | Time when the template was last modified. | 

## Methods

### NewMessageTemplate

`func NewMessageTemplate(id int32, name string, content string, lastModified NullableTime, ) *MessageTemplate`

NewMessageTemplate instantiates a new MessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTemplateWithDefaults

`func NewMessageTemplateWithDefaults() *MessageTemplate`

NewMessageTemplateWithDefaults instantiates a new MessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *MessageTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *MessageTemplate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageTemplate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageTemplate) SetContent(v string)`

SetContent sets Content field to given value.


### GetLastModified

`func (o *MessageTemplate) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MessageTemplate) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MessageTemplate) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *MessageTemplate) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *MessageTemplate) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


