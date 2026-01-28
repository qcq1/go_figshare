# CustomArticleField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Custom  metadata name | 
**Value** | **map[string]interface{}** | Custom metadata value (can be either a string or an array of strings) | 
**FieldType** | **string** | Custom field type | 
**Settings** | **map[string]interface{}** | Settings for the custom field | 
**Order** | **int64** | Order of the custom field | 
**IsMandatory** | **bool** | Whether the field is mandatory or not | 

## Methods

### NewCustomArticleField

`func NewCustomArticleField(name string, value map[string]interface{}, fieldType string, settings map[string]interface{}, order int64, isMandatory bool, ) *CustomArticleField`

NewCustomArticleField instantiates a new CustomArticleField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomArticleFieldWithDefaults

`func NewCustomArticleFieldWithDefaults() *CustomArticleField`

NewCustomArticleFieldWithDefaults instantiates a new CustomArticleField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomArticleField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomArticleField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomArticleField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CustomArticleField) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomArticleField) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomArticleField) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetFieldType

`func (o *CustomArticleField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *CustomArticleField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *CustomArticleField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetSettings

`func (o *CustomArticleField) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CustomArticleField) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CustomArticleField) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetOrder

`func (o *CustomArticleField) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CustomArticleField) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CustomArticleField) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetIsMandatory

`func (o *CustomArticleField) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *CustomArticleField) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *CustomArticleField) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


