# DoEmailLookupResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The email address passed to the call. | [default to null]
**Status** | **string** | The email is &#x60;valid&#x60; or &#x60;invalid&#x60;. | [default to null]
**Deliverability** | **string** | The delivery status of the email address is&#x60;deliverable&#x60;, &#x60;undeliverable&#x60;  or &#x60;unknown&#x60;. | [default to null]
**Reason** | **string** | The reason why the checked email is invalid/undeliverable. | [default to null]
**Risk** | **string** | The risk score of the email is&#x60;high&#x60;, &#x60;medium&#x60;, &#x60;low&#x60; or &#x60;null&#x60;. | [default to null]
**AddressType** | **string** | The email address type (domain) is &#x60;free&#x60; or &#x60;corporate&#x60;. | [default to null]
**IsDisposableAddress** | **bool** | This is be &#x60;true&#x60; if the domain is in the list of disposable email addresses, otherwise returns as &#x60;false&#x60;. | [default to null]
**Suggestion** | **string** | Null if nothing is suggested, however, if there is a potential typo in the email address, the closest suggestion is provided. | [default to null]
**EmailRole** | **string** | Checks the mailbox part of the email whether it matches a specific role type (‘admin’, ‘sales’, ‘webmaster’) | [default to null]
**LocalPart** | **string** | The local part of the email address. | [default to null]
**DomainPart** | **string** | The domain part of the email address. | [default to null]
**Exchange** | **string** | Email exchange server domain name (MX record value). | [default to null]
**Preference** | **int32** | MX record preference. | [default to null]
**IsInWhiteList** | **bool** | &#x60;true&#x60; if the email address exists in TextMagic whitelist.  | [default to null]
**IsInBlackList** | **bool** | &#x60;true&#x60; if the email address exists in TextMagic blacklist.  | [default to null]
**HasMx** | **bool** | &#x60;true&#x60; if the email address domain has an MX record.  | [default to null]
**HasAa** | **bool** | &#x60;true&#x60; if the email address domain has an A record (IPv4).  | [default to null]
**HasAaaa** | **bool** | &#x60;true&#x60; if the email address domain has an AAAA record (IPv6).  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


