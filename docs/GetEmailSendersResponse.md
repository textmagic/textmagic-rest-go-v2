# GetEmailSendersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EmailSenderItem**](EmailSenderItem.md) | Array of email sender objects. | 

## Methods

### NewGetEmailSendersResponse

`func NewGetEmailSendersResponse(items []EmailSenderItem, ) *GetEmailSendersResponse`

NewGetEmailSendersResponse instantiates a new GetEmailSendersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailSendersResponseWithDefaults

`func NewGetEmailSendersResponseWithDefaults() *GetEmailSendersResponse`

NewGetEmailSendersResponseWithDefaults instantiates a new GetEmailSendersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetEmailSendersResponse) GetItems() []EmailSenderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetEmailSendersResponse) GetItemsOk() (*[]EmailSenderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetEmailSendersResponse) SetItems(v []EmailSenderItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


