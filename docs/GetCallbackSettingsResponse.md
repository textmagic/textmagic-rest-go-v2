# GetCallbackSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutUrl** | **NullableString** | This URL is used to push message delivery status updates to your application. | 
**InUrl** | **NullableString** | This URL is used to push incoming SMS to your application. | 
**Format** | **string** | Desired callback data format. m - multipart/form-data, u - application/x-www-form-urlencoded, j - application/json. | 

## Methods

### NewGetCallbackSettingsResponse

`func NewGetCallbackSettingsResponse(outUrl NullableString, inUrl NullableString, format string, ) *GetCallbackSettingsResponse`

NewGetCallbackSettingsResponse instantiates a new GetCallbackSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallbackSettingsResponseWithDefaults

`func NewGetCallbackSettingsResponseWithDefaults() *GetCallbackSettingsResponse`

NewGetCallbackSettingsResponseWithDefaults instantiates a new GetCallbackSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutUrl

`func (o *GetCallbackSettingsResponse) GetOutUrl() string`

GetOutUrl returns the OutUrl field if non-nil, zero value otherwise.

### GetOutUrlOk

`func (o *GetCallbackSettingsResponse) GetOutUrlOk() (*string, bool)`

GetOutUrlOk returns a tuple with the OutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutUrl

`func (o *GetCallbackSettingsResponse) SetOutUrl(v string)`

SetOutUrl sets OutUrl field to given value.


### SetOutUrlNil

`func (o *GetCallbackSettingsResponse) SetOutUrlNil(b bool)`

 SetOutUrlNil sets the value for OutUrl to be an explicit nil

### UnsetOutUrl
`func (o *GetCallbackSettingsResponse) UnsetOutUrl()`

UnsetOutUrl ensures that no value is present for OutUrl, not even an explicit nil
### GetInUrl

`func (o *GetCallbackSettingsResponse) GetInUrl() string`

GetInUrl returns the InUrl field if non-nil, zero value otherwise.

### GetInUrlOk

`func (o *GetCallbackSettingsResponse) GetInUrlOk() (*string, bool)`

GetInUrlOk returns a tuple with the InUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUrl

`func (o *GetCallbackSettingsResponse) SetInUrl(v string)`

SetInUrl sets InUrl field to given value.


### SetInUrlNil

`func (o *GetCallbackSettingsResponse) SetInUrlNil(b bool)`

 SetInUrlNil sets the value for InUrl to be an explicit nil

### UnsetInUrl
`func (o *GetCallbackSettingsResponse) UnsetInUrl()`

UnsetInUrl ensures that no value is present for InUrl, not even an explicit nil
### GetFormat

`func (o *GetCallbackSettingsResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GetCallbackSettingsResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GetCallbackSettingsResponse) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


