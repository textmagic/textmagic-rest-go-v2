# DoAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Token** | **string** |  | 
**Expires** | **NullableTime** |  | 
**MinVersions** | [**DoAuthResponseMinVersions**](DoAuthResponseMinVersions.md) |  | 
**DisallowedRules** | **[]string** |  | 

## Methods

### NewDoAuthResponse

`func NewDoAuthResponse(username string, token string, expires NullableTime, minVersions DoAuthResponseMinVersions, disallowedRules []string, ) *DoAuthResponse`

NewDoAuthResponse instantiates a new DoAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoAuthResponseWithDefaults

`func NewDoAuthResponseWithDefaults() *DoAuthResponse`

NewDoAuthResponseWithDefaults instantiates a new DoAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DoAuthResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DoAuthResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DoAuthResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetToken

`func (o *DoAuthResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DoAuthResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DoAuthResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpires

`func (o *DoAuthResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DoAuthResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DoAuthResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *DoAuthResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *DoAuthResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetMinVersions

`func (o *DoAuthResponse) GetMinVersions() DoAuthResponseMinVersions`

GetMinVersions returns the MinVersions field if non-nil, zero value otherwise.

### GetMinVersionsOk

`func (o *DoAuthResponse) GetMinVersionsOk() (*DoAuthResponseMinVersions, bool)`

GetMinVersionsOk returns a tuple with the MinVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersions

`func (o *DoAuthResponse) SetMinVersions(v DoAuthResponseMinVersions)`

SetMinVersions sets MinVersions field to given value.


### GetDisallowedRules

`func (o *DoAuthResponse) GetDisallowedRules() []string`

GetDisallowedRules returns the DisallowedRules field if non-nil, zero value otherwise.

### GetDisallowedRulesOk

`func (o *DoAuthResponse) GetDisallowedRulesOk() (*[]string, bool)`

GetDisallowedRulesOk returns a tuple with the DisallowedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedRules

`func (o *DoAuthResponse) SetDisallowedRules(v []string)`

SetDisallowedRules sets DisallowedRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


