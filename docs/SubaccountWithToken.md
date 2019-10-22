# SubaccountWithToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Sub-account ID. | [default to null]
**Username** | **string** | Username. | [default to null]
**FirstName** | **string** | Account first name. | [default to null]
**LastName** | **string** | Account last name. | [default to null]
**Email** | **string** | Account Email address. | [default to null]
**Status** | **string** | Current account status: * **A** for Active * **T** for Trial.  | [default to null]
**Balance** | **float64** | Account balance (in account currency). | [default to null]
**Phone** | **string** | Contact phone number. | [default to null]
**Company** | **string** | Account company name. | [default to null]
**Currency** | [***Currency**](Currency.md) |  | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]
**Timezone** | [***Timezone**](Timezone.md) |  | [default to null]
**SubaccountType** | **string** | Type of account: *   **A** for Administrator sub-account *   **U** for Regular User  | [default to null]
**EmailAccepted** | **bool** | Is account has confirmed Email. | [default to null]
**PhoneAccepted** | **bool** | Is account has confirmed Phone number. | [default to null]
**Avatar** | [***UserImage**](UserImage.md) |  | [default to null]
**Token** | **string** | Access token of account. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


