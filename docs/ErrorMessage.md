# ErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | A machine friendly error code, used by the dev team to identify the error. | [optional] 
**Message** | Pointer to **string** | A human friendly message explaining the error. | [optional] 

## Methods

### NewErrorMessage

`func NewErrorMessage() *ErrorMessage`

NewErrorMessage instantiates a new ErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageWithDefaults

`func NewErrorMessageWithDefaults() *ErrorMessage`

NewErrorMessageWithDefaults instantiates a new ErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorMessage) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorMessage) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorMessage) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorMessage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


