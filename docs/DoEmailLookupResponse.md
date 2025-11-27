# DoEmailLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The email address passed to the call. | 
**Status** | **string** | The email is &#x60;valid&#x60; or &#x60;invalid&#x60;. | 
**Deliverability** | **string** | The delivery status of the email address is&#x60;deliverable&#x60;, &#x60;undeliverable&#x60;. or &#x60;unknown&#x60;. | 
**Reason** | **NullableString** | The reason why the checked email is invalid/undeliverable. | 
**Risk** | **string** | The risk score of the email is&#x60;high&#x60;, &#x60;medium&#x60;, &#x60;low&#x60; or &#x60;null&#x60;. | 
**AddressType** | **string** | The email address type (domain) is &#x60;free&#x60; or &#x60;corporate&#x60;. | 
**IsDisposableAddress** | **bool** | This is &#x60;true&#x60; if the domain is in the list of disposable email addresses; otherwise, it returns as &#x60;false&#x60;. | 
**Suggestion** | **NullableString** | Null if nothing is suggested; however, if there is a potential typo in the email address, the closest suggestion is provided. | 
**EmailRole** | **NullableString** | Checks the mailbox part of the email to see whether it matches a specific role type (‘admin’, ‘sales’, ‘webmaster’). | 
**LocalPart** | **string** | The local part of the email address. | 
**DomainPart** | **string** | The domain part of the email address. | 
**Exchange** | **NullableString** | Email exchange server domain name (MX record value). | 
**Preference** | **NullableInt32** | MX record preference. | 
**IsInWhiteList** | **bool** | &#x60;true&#x60; if the email address exists in the Textmagic whitelist.  | 
**IsInBlackList** | **bool** | &#x60;true&#x60; if the email address exists in the Textmagic blacklist.  | 
**HasMx** | **bool** | &#x60;true&#x60; if the email address domain has an MX record.  | 
**HasAa** | **bool** | &#x60;true&#x60; if the email address domain has an A record (IPv4).  | 
**HasAaaa** | **bool** | &#x60;true&#x60; if the email address domain has an AAAA record (IPv6).  | 

## Methods

### NewDoEmailLookupResponse

`func NewDoEmailLookupResponse(address string, status string, deliverability string, reason NullableString, risk string, addressType string, isDisposableAddress bool, suggestion NullableString, emailRole NullableString, localPart string, domainPart string, exchange NullableString, preference NullableInt32, isInWhiteList bool, isInBlackList bool, hasMx bool, hasAa bool, hasAaaa bool, ) *DoEmailLookupResponse`

NewDoEmailLookupResponse instantiates a new DoEmailLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoEmailLookupResponseWithDefaults

`func NewDoEmailLookupResponseWithDefaults() *DoEmailLookupResponse`

NewDoEmailLookupResponseWithDefaults instantiates a new DoEmailLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DoEmailLookupResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DoEmailLookupResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DoEmailLookupResponse) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetStatus

`func (o *DoEmailLookupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DoEmailLookupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DoEmailLookupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeliverability

`func (o *DoEmailLookupResponse) GetDeliverability() string`

GetDeliverability returns the Deliverability field if non-nil, zero value otherwise.

### GetDeliverabilityOk

`func (o *DoEmailLookupResponse) GetDeliverabilityOk() (*string, bool)`

GetDeliverabilityOk returns a tuple with the Deliverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverability

`func (o *DoEmailLookupResponse) SetDeliverability(v string)`

SetDeliverability sets Deliverability field to given value.


### GetReason

`func (o *DoEmailLookupResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DoEmailLookupResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DoEmailLookupResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *DoEmailLookupResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DoEmailLookupResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRisk

`func (o *DoEmailLookupResponse) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *DoEmailLookupResponse) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *DoEmailLookupResponse) SetRisk(v string)`

SetRisk sets Risk field to given value.


### GetAddressType

`func (o *DoEmailLookupResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *DoEmailLookupResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *DoEmailLookupResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetIsDisposableAddress

`func (o *DoEmailLookupResponse) GetIsDisposableAddress() bool`

GetIsDisposableAddress returns the IsDisposableAddress field if non-nil, zero value otherwise.

### GetIsDisposableAddressOk

`func (o *DoEmailLookupResponse) GetIsDisposableAddressOk() (*bool, bool)`

GetIsDisposableAddressOk returns a tuple with the IsDisposableAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisposableAddress

`func (o *DoEmailLookupResponse) SetIsDisposableAddress(v bool)`

SetIsDisposableAddress sets IsDisposableAddress field to given value.


### GetSuggestion

`func (o *DoEmailLookupResponse) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *DoEmailLookupResponse) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *DoEmailLookupResponse) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.


### SetSuggestionNil

`func (o *DoEmailLookupResponse) SetSuggestionNil(b bool)`

 SetSuggestionNil sets the value for Suggestion to be an explicit nil

### UnsetSuggestion
`func (o *DoEmailLookupResponse) UnsetSuggestion()`

UnsetSuggestion ensures that no value is present for Suggestion, not even an explicit nil
### GetEmailRole

`func (o *DoEmailLookupResponse) GetEmailRole() string`

GetEmailRole returns the EmailRole field if non-nil, zero value otherwise.

### GetEmailRoleOk

`func (o *DoEmailLookupResponse) GetEmailRoleOk() (*string, bool)`

GetEmailRoleOk returns a tuple with the EmailRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRole

`func (o *DoEmailLookupResponse) SetEmailRole(v string)`

SetEmailRole sets EmailRole field to given value.


### SetEmailRoleNil

`func (o *DoEmailLookupResponse) SetEmailRoleNil(b bool)`

 SetEmailRoleNil sets the value for EmailRole to be an explicit nil

### UnsetEmailRole
`func (o *DoEmailLookupResponse) UnsetEmailRole()`

UnsetEmailRole ensures that no value is present for EmailRole, not even an explicit nil
### GetLocalPart

`func (o *DoEmailLookupResponse) GetLocalPart() string`

GetLocalPart returns the LocalPart field if non-nil, zero value otherwise.

### GetLocalPartOk

`func (o *DoEmailLookupResponse) GetLocalPartOk() (*string, bool)`

GetLocalPartOk returns a tuple with the LocalPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPart

`func (o *DoEmailLookupResponse) SetLocalPart(v string)`

SetLocalPart sets LocalPart field to given value.


### GetDomainPart

`func (o *DoEmailLookupResponse) GetDomainPart() string`

GetDomainPart returns the DomainPart field if non-nil, zero value otherwise.

### GetDomainPartOk

`func (o *DoEmailLookupResponse) GetDomainPartOk() (*string, bool)`

GetDomainPartOk returns a tuple with the DomainPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPart

`func (o *DoEmailLookupResponse) SetDomainPart(v string)`

SetDomainPart sets DomainPart field to given value.


### GetExchange

`func (o *DoEmailLookupResponse) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DoEmailLookupResponse) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DoEmailLookupResponse) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### SetExchangeNil

`func (o *DoEmailLookupResponse) SetExchangeNil(b bool)`

 SetExchangeNil sets the value for Exchange to be an explicit nil

### UnsetExchange
`func (o *DoEmailLookupResponse) UnsetExchange()`

UnsetExchange ensures that no value is present for Exchange, not even an explicit nil
### GetPreference

`func (o *DoEmailLookupResponse) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DoEmailLookupResponse) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *DoEmailLookupResponse) SetPreference(v int32)`

SetPreference sets Preference field to given value.


### SetPreferenceNil

`func (o *DoEmailLookupResponse) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *DoEmailLookupResponse) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetIsInWhiteList

`func (o *DoEmailLookupResponse) GetIsInWhiteList() bool`

GetIsInWhiteList returns the IsInWhiteList field if non-nil, zero value otherwise.

### GetIsInWhiteListOk

`func (o *DoEmailLookupResponse) GetIsInWhiteListOk() (*bool, bool)`

GetIsInWhiteListOk returns a tuple with the IsInWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInWhiteList

`func (o *DoEmailLookupResponse) SetIsInWhiteList(v bool)`

SetIsInWhiteList sets IsInWhiteList field to given value.


### GetIsInBlackList

`func (o *DoEmailLookupResponse) GetIsInBlackList() bool`

GetIsInBlackList returns the IsInBlackList field if non-nil, zero value otherwise.

### GetIsInBlackListOk

`func (o *DoEmailLookupResponse) GetIsInBlackListOk() (*bool, bool)`

GetIsInBlackListOk returns a tuple with the IsInBlackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInBlackList

`func (o *DoEmailLookupResponse) SetIsInBlackList(v bool)`

SetIsInBlackList sets IsInBlackList field to given value.


### GetHasMx

`func (o *DoEmailLookupResponse) GetHasMx() bool`

GetHasMx returns the HasMx field if non-nil, zero value otherwise.

### GetHasMxOk

`func (o *DoEmailLookupResponse) GetHasMxOk() (*bool, bool)`

GetHasMxOk returns a tuple with the HasMx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMx

`func (o *DoEmailLookupResponse) SetHasMx(v bool)`

SetHasMx sets HasMx field to given value.


### GetHasAa

`func (o *DoEmailLookupResponse) GetHasAa() bool`

GetHasAa returns the HasAa field if non-nil, zero value otherwise.

### GetHasAaOk

`func (o *DoEmailLookupResponse) GetHasAaOk() (*bool, bool)`

GetHasAaOk returns a tuple with the HasAa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAa

`func (o *DoEmailLookupResponse) SetHasAa(v bool)`

SetHasAa sets HasAa field to given value.


### GetHasAaaa

`func (o *DoEmailLookupResponse) GetHasAaaa() bool`

GetHasAaaa returns the HasAaaa field if non-nil, zero value otherwise.

### GetHasAaaaOk

`func (o *DoEmailLookupResponse) GetHasAaaaOk() (*bool, bool)`

GetHasAaaaOk returns a tuple with the HasAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAaaa

`func (o *DoEmailLookupResponse) SetHasAaaa(v bool)`

SetHasAaaa sets HasAaaa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


