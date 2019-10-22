# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User ID. | [default to null]
**DisplayTimeFormat** | **string** | User&#39;s prefered format of time display * *12h* - AM/PM format * *24h* - 24 hour clock format  | [optional] [default to null]
**Username** | **string** | Username. | [default to null]
**FirstName** | **string** | Account first name. | [default to null]
**LastName** | **string** | Account last name. | [default to null]
**Email** | **string** | User email address. | [default to null]
**Status** | **string** | Current account status: * **A** for Active * **T** for Trial.  | [default to null]
**Balance** | **float32** | Account balance (in account currency). | [default to null]
**Phone** | **string** | User phone number | [default to null]
**Company** | **string** | Account company name. | [default to null]
**Currency** | [***Currency**](Currency.md) |  | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]
**Timezone** | [***Timezone**](Timezone.md) |  | [default to null]
**SubaccountType** | **string** | Type of account: * **P** for Parent User * **A** for Administrator Sub-Account * **U** for Regular User  | [default to null]
**EmailAccepted** | **bool** | Is account has confirmed Email. | [default to null]
**PhoneAccepted** | **bool** | Is account has confirmed Phone number. | [default to null]
**Avatar** | [***UserImage**](UserImage.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


