# Contact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Contact ID. | [default to null]
**Favorited** | **bool** | Is the Contact favourite? [Custom fields list](http://docs.textmagictesting.com/#operation/getFavourites). | [default to null]
**Blocked** | **bool** | Is the Contact blocked? [Custom fields list](http://docs.textmagictesting.com/#operation/getBlockedContacts). | [default to null]
**FirstName** | **string** | Contact first name. | [default to null]
**LastName** | **string** | Contact last name. | [default to null]
**CompanyName** | **string** | Company name. | [default to null]
**Phone** | **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | [default to null]
**Email** | **string** | Contact email address. | [default to null]
**Country** | [***Country**](Country.md) | Contact country. | [default to null]
**CustomFields** | [**[]ContactCustomField**](ContactCustomField.md) | See [Custom Fields](http://docs.textmagictesting.com/#tag/Custom-Fields) section. | [default to null]
**User** | [***User**](User.md) |  | [default to null]
**Lists** | [**[]List**](List.md) |  | [default to null]
**PhoneType** | **string** | Phone number type: * **0** if it is fixed-line; * **1** if it is mobile; * **2** if it is mobile or fixed-line (in case we cannot distingush between fixed-line or mobile); * **3** if it is toll-free; * **4** if it is a premium rate phone; * **5** if it is a shared cost phone; * **6** if it is a VoIP; * **7** if it is a [Personal Number](); * **8** if it is a pager; * **9** if it is an Universal Access Number; * **10** if the phone type is unknown; * **-1** if the phone type is not yet processed or cannot be determined.  | [default to null]
**Avatar** | [***ContactImage**](ContactImage.md) |  | [default to null]
**Notes** | [**[]ContactNote**](ContactNote.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


