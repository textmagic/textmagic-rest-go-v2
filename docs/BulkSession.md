# BulkSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Bulk Session ID. | 
**Status** | **string** | * **n** – bulk session is just created * **w** - work in progress * **f** - failed * **c** - completed with success * **s** - suspended  | 
**ItemsProcessed** | **NullableInt32** | Amount of messages already processed. | 
**ItemsTotal** | **NullableInt32** | Total amount of messages to be processed. | 
**CreatedAt** | **time.Time** | Creation date and time of a Bulk Session. | 
**Session** | [**NullableMessageSession**](MessageSession.md) |  | 
**Text** | **NullableString** | Message text of a Bulk Session. | 

## Methods

### NewBulkSession

`func NewBulkSession(id int32, status string, itemsProcessed NullableInt32, itemsTotal NullableInt32, createdAt time.Time, session NullableMessageSession, text NullableString, ) *BulkSession`

NewBulkSession instantiates a new BulkSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSessionWithDefaults

`func NewBulkSessionWithDefaults() *BulkSession`

NewBulkSessionWithDefaults instantiates a new BulkSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkSession) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkSession) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *BulkSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetItemsProcessed

`func (o *BulkSession) GetItemsProcessed() int32`

GetItemsProcessed returns the ItemsProcessed field if non-nil, zero value otherwise.

### GetItemsProcessedOk

`func (o *BulkSession) GetItemsProcessedOk() (*int32, bool)`

GetItemsProcessedOk returns a tuple with the ItemsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsProcessed

`func (o *BulkSession) SetItemsProcessed(v int32)`

SetItemsProcessed sets ItemsProcessed field to given value.


### SetItemsProcessedNil

`func (o *BulkSession) SetItemsProcessedNil(b bool)`

 SetItemsProcessedNil sets the value for ItemsProcessed to be an explicit nil

### UnsetItemsProcessed
`func (o *BulkSession) UnsetItemsProcessed()`

UnsetItemsProcessed ensures that no value is present for ItemsProcessed, not even an explicit nil
### GetItemsTotal

`func (o *BulkSession) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *BulkSession) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *BulkSession) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### SetItemsTotalNil

`func (o *BulkSession) SetItemsTotalNil(b bool)`

 SetItemsTotalNil sets the value for ItemsTotal to be an explicit nil

### UnsetItemsTotal
`func (o *BulkSession) UnsetItemsTotal()`

UnsetItemsTotal ensures that no value is present for ItemsTotal, not even an explicit nil
### GetCreatedAt

`func (o *BulkSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSession

`func (o *BulkSession) GetSession() MessageSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *BulkSession) GetSessionOk() (*MessageSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *BulkSession) SetSession(v MessageSession)`

SetSession sets Session field to given value.


### SetSessionNil

`func (o *BulkSession) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *BulkSession) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetText

`func (o *BulkSession) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BulkSession) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BulkSession) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *BulkSession) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *BulkSession) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


