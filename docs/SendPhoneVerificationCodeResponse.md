# SendPhoneVerificationCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyId** | **string** | The ID of a verification request. This is required to finish the verification request in the next step. | 
**Price** | **float32** | An amount of credit which will be deducted from your account balance when this verification is successfully completed. | 

## Methods

### NewSendPhoneVerificationCodeResponse

`func NewSendPhoneVerificationCodeResponse(verifyId string, price float32, ) *SendPhoneVerificationCodeResponse`

NewSendPhoneVerificationCodeResponse instantiates a new SendPhoneVerificationCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPhoneVerificationCodeResponseWithDefaults

`func NewSendPhoneVerificationCodeResponseWithDefaults() *SendPhoneVerificationCodeResponse`

NewSendPhoneVerificationCodeResponseWithDefaults instantiates a new SendPhoneVerificationCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyId

`func (o *SendPhoneVerificationCodeResponse) GetVerifyId() string`

GetVerifyId returns the VerifyId field if non-nil, zero value otherwise.

### GetVerifyIdOk

`func (o *SendPhoneVerificationCodeResponse) GetVerifyIdOk() (*string, bool)`

GetVerifyIdOk returns a tuple with the VerifyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyId

`func (o *SendPhoneVerificationCodeResponse) SetVerifyId(v string)`

SetVerifyId sets VerifyId field to given value.


### GetPrice

`func (o *SendPhoneVerificationCodeResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SendPhoneVerificationCodeResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SendPhoneVerificationCodeResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


