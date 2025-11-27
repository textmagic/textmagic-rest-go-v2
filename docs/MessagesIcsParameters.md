# MessagesIcsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Scheduled message text. | 
**Recipients** | [**MessagesIcsParametersRecipients**](MessagesIcsParametersRecipients.md) |  | 

## Methods

### NewMessagesIcsParameters

`func NewMessagesIcsParameters(text string, recipients MessagesIcsParametersRecipients, ) *MessagesIcsParameters`

NewMessagesIcsParameters instantiates a new MessagesIcsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesIcsParametersWithDefaults

`func NewMessagesIcsParametersWithDefaults() *MessagesIcsParameters`

NewMessagesIcsParametersWithDefaults instantiates a new MessagesIcsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MessagesIcsParameters) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessagesIcsParameters) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessagesIcsParameters) SetText(v string)`

SetText sets Text field to given value.


### GetRecipients

`func (o *MessagesIcsParameters) GetRecipients() MessagesIcsParametersRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *MessagesIcsParameters) GetRecipientsOk() (*MessagesIcsParametersRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *MessagesIcsParameters) SetRecipients(v MessagesIcsParametersRecipients)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


