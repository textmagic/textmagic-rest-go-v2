# UpdateCallbackSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutUrl** | Pointer to **NullableString** | This URL is used to push message delivery status updates to your application. | [optional] 
**InUrl** | Pointer to **NullableString** | This URL is used to push incoming SMS to your application. | [optional] 
**Format** | Pointer to **string** | Desired callback data format. m - multipart/form-data, u - application/x-www-form-urlencoded, j - application/json | [optional] 

## Methods

### NewUpdateCallbackSettingsRequest

`func NewUpdateCallbackSettingsRequest() *UpdateCallbackSettingsRequest`

NewUpdateCallbackSettingsRequest instantiates a new UpdateCallbackSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallbackSettingsRequestWithDefaults

`func NewUpdateCallbackSettingsRequestWithDefaults() *UpdateCallbackSettingsRequest`

NewUpdateCallbackSettingsRequestWithDefaults instantiates a new UpdateCallbackSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutUrl

`func (o *UpdateCallbackSettingsRequest) GetOutUrl() string`

GetOutUrl returns the OutUrl field if non-nil, zero value otherwise.

### GetOutUrlOk

`func (o *UpdateCallbackSettingsRequest) GetOutUrlOk() (*string, bool)`

GetOutUrlOk returns a tuple with the OutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutUrl

`func (o *UpdateCallbackSettingsRequest) SetOutUrl(v string)`

SetOutUrl sets OutUrl field to given value.

### HasOutUrl

`func (o *UpdateCallbackSettingsRequest) HasOutUrl() bool`

HasOutUrl returns a boolean if a field has been set.

### SetOutUrlNil

`func (o *UpdateCallbackSettingsRequest) SetOutUrlNil(b bool)`

 SetOutUrlNil sets the value for OutUrl to be an explicit nil

### UnsetOutUrl
`func (o *UpdateCallbackSettingsRequest) UnsetOutUrl()`

UnsetOutUrl ensures that no value is present for OutUrl, not even an explicit nil
### GetInUrl

`func (o *UpdateCallbackSettingsRequest) GetInUrl() string`

GetInUrl returns the InUrl field if non-nil, zero value otherwise.

### GetInUrlOk

`func (o *UpdateCallbackSettingsRequest) GetInUrlOk() (*string, bool)`

GetInUrlOk returns a tuple with the InUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUrl

`func (o *UpdateCallbackSettingsRequest) SetInUrl(v string)`

SetInUrl sets InUrl field to given value.

### HasInUrl

`func (o *UpdateCallbackSettingsRequest) HasInUrl() bool`

HasInUrl returns a boolean if a field has been set.

### SetInUrlNil

`func (o *UpdateCallbackSettingsRequest) SetInUrlNil(b bool)`

 SetInUrlNil sets the value for InUrl to be an explicit nil

### UnsetInUrl
`func (o *UpdateCallbackSettingsRequest) UnsetInUrl()`

UnsetInUrl ensures that no value is present for InUrl, not even an explicit nil
### GetFormat

`func (o *UpdateCallbackSettingsRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateCallbackSettingsRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateCallbackSettingsRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateCallbackSettingsRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


