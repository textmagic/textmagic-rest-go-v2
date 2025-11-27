# GetAvailableSenderSettingOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | **[]string** | Array of dedicated number strings. | 
**User** | **[]string** | Array of verified account phone numbers (currently only one). | 
**Shared** | **[]string** | Array of shared number strings. | 
**SenderIds** | **[]string** | Array of alphanumeric sender IDs. | 
**UserCarrierTwilio** | **[]string** | Array of alphanumeric sender IDs. | 
**UserCarrierVonage** | **[]string** | Array of alphanumeric sender IDs. | 
**UserCarrierSinch** | **[]string** | Array of alphanumeric sender IDs. | 
**UCarrierBandwidth** | Pointer to **[]string** | Array of alphanumeric sender IDs. | [optional] 
**UcTwilioSenderId** | Pointer to **[]string** | Array of alphanumeric sender IDs. | [optional] 

## Methods

### NewGetAvailableSenderSettingOptionsResponse

`func NewGetAvailableSenderSettingOptionsResponse(dedicated []string, user []string, shared []string, senderIds []string, userCarrierTwilio []string, userCarrierVonage []string, userCarrierSinch []string, ) *GetAvailableSenderSettingOptionsResponse`

NewGetAvailableSenderSettingOptionsResponse instantiates a new GetAvailableSenderSettingOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAvailableSenderSettingOptionsResponseWithDefaults

`func NewGetAvailableSenderSettingOptionsResponseWithDefaults() *GetAvailableSenderSettingOptionsResponse`

NewGetAvailableSenderSettingOptionsResponseWithDefaults instantiates a new GetAvailableSenderSettingOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *GetAvailableSenderSettingOptionsResponse) GetDedicated() []string`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetDedicatedOk() (*[]string, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *GetAvailableSenderSettingOptionsResponse) SetDedicated(v []string)`

SetDedicated sets Dedicated field to given value.


### GetUser

`func (o *GetAvailableSenderSettingOptionsResponse) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetAvailableSenderSettingOptionsResponse) SetUser(v []string)`

SetUser sets User field to given value.


### GetShared

`func (o *GetAvailableSenderSettingOptionsResponse) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GetAvailableSenderSettingOptionsResponse) SetShared(v []string)`

SetShared sets Shared field to given value.


### GetSenderIds

`func (o *GetAvailableSenderSettingOptionsResponse) GetSenderIds() []string`

GetSenderIds returns the SenderIds field if non-nil, zero value otherwise.

### GetSenderIdsOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetSenderIdsOk() (*[]string, bool)`

GetSenderIdsOk returns a tuple with the SenderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderIds

`func (o *GetAvailableSenderSettingOptionsResponse) SetSenderIds(v []string)`

SetSenderIds sets SenderIds field to given value.


### GetUserCarrierTwilio

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierTwilio() []string`

GetUserCarrierTwilio returns the UserCarrierTwilio field if non-nil, zero value otherwise.

### GetUserCarrierTwilioOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierTwilioOk() (*[]string, bool)`

GetUserCarrierTwilioOk returns a tuple with the UserCarrierTwilio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCarrierTwilio

`func (o *GetAvailableSenderSettingOptionsResponse) SetUserCarrierTwilio(v []string)`

SetUserCarrierTwilio sets UserCarrierTwilio field to given value.


### GetUserCarrierVonage

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierVonage() []string`

GetUserCarrierVonage returns the UserCarrierVonage field if non-nil, zero value otherwise.

### GetUserCarrierVonageOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierVonageOk() (*[]string, bool)`

GetUserCarrierVonageOk returns a tuple with the UserCarrierVonage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCarrierVonage

`func (o *GetAvailableSenderSettingOptionsResponse) SetUserCarrierVonage(v []string)`

SetUserCarrierVonage sets UserCarrierVonage field to given value.


### GetUserCarrierSinch

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierSinch() []string`

GetUserCarrierSinch returns the UserCarrierSinch field if non-nil, zero value otherwise.

### GetUserCarrierSinchOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUserCarrierSinchOk() (*[]string, bool)`

GetUserCarrierSinchOk returns a tuple with the UserCarrierSinch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCarrierSinch

`func (o *GetAvailableSenderSettingOptionsResponse) SetUserCarrierSinch(v []string)`

SetUserCarrierSinch sets UserCarrierSinch field to given value.


### GetUCarrierBandwidth

`func (o *GetAvailableSenderSettingOptionsResponse) GetUCarrierBandwidth() []string`

GetUCarrierBandwidth returns the UCarrierBandwidth field if non-nil, zero value otherwise.

### GetUCarrierBandwidthOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUCarrierBandwidthOk() (*[]string, bool)`

GetUCarrierBandwidthOk returns a tuple with the UCarrierBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUCarrierBandwidth

`func (o *GetAvailableSenderSettingOptionsResponse) SetUCarrierBandwidth(v []string)`

SetUCarrierBandwidth sets UCarrierBandwidth field to given value.

### HasUCarrierBandwidth

`func (o *GetAvailableSenderSettingOptionsResponse) HasUCarrierBandwidth() bool`

HasUCarrierBandwidth returns a boolean if a field has been set.

### GetUcTwilioSenderId

`func (o *GetAvailableSenderSettingOptionsResponse) GetUcTwilioSenderId() []string`

GetUcTwilioSenderId returns the UcTwilioSenderId field if non-nil, zero value otherwise.

### GetUcTwilioSenderIdOk

`func (o *GetAvailableSenderSettingOptionsResponse) GetUcTwilioSenderIdOk() (*[]string, bool)`

GetUcTwilioSenderIdOk returns a tuple with the UcTwilioSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcTwilioSenderId

`func (o *GetAvailableSenderSettingOptionsResponse) SetUcTwilioSenderId(v []string)`

SetUcTwilioSenderId sets UcTwilioSenderId field to given value.

### HasUcTwilioSenderId

`func (o *GetAvailableSenderSettingOptionsResponse) HasUcTwilioSenderId() bool`

HasUcTwilioSenderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


