# GetContactImportSessionProgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | Session status: * **1** - if session has been initialized but not yet started; * **3** - if session is being processed; * **4** - if session has errors; * **5** - if session completed successfully.  | 
**Processed** | **int32** | How many contacts have been imported? | 

## Methods

### NewGetContactImportSessionProgressResponse

`func NewGetContactImportSessionProgressResponse(status int32, processed int32, ) *GetContactImportSessionProgressResponse`

NewGetContactImportSessionProgressResponse instantiates a new GetContactImportSessionProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContactImportSessionProgressResponseWithDefaults

`func NewGetContactImportSessionProgressResponseWithDefaults() *GetContactImportSessionProgressResponse`

NewGetContactImportSessionProgressResponseWithDefaults instantiates a new GetContactImportSessionProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetContactImportSessionProgressResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContactImportSessionProgressResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContactImportSessionProgressResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetProcessed

`func (o *GetContactImportSessionProgressResponse) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *GetContactImportSessionProgressResponse) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *GetContactImportSessionProgressResponse) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


