# ShortCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Custom field id | 
**Name** | **string** | Custom field name | 
**FieldType** | **string** | Custom field type | 
**Settings** | Pointer to **map[string]interface{}** | Settings for the custom field | [optional] 
**Order** | Pointer to **int64** | Order of the field in the group | [optional] 
**IsMandatory** | Pointer to **bool** | Whether the field is mandatory or not | [optional] 

## Methods

### NewShortCustomField

`func NewShortCustomField(id int64, name string, fieldType string, ) *ShortCustomField`

NewShortCustomField instantiates a new ShortCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortCustomFieldWithDefaults

`func NewShortCustomFieldWithDefaults() *ShortCustomField`

NewShortCustomFieldWithDefaults instantiates a new ShortCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShortCustomField) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortCustomField) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortCustomField) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ShortCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShortCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShortCustomField) SetName(v string)`

SetName sets Name field to given value.


### GetFieldType

`func (o *ShortCustomField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ShortCustomField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ShortCustomField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetSettings

`func (o *ShortCustomField) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ShortCustomField) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ShortCustomField) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ShortCustomField) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetOrder

`func (o *ShortCustomField) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ShortCustomField) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ShortCustomField) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ShortCustomField) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsMandatory

`func (o *ShortCustomField) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *ShortCustomField) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *ShortCustomField) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *ShortCustomField) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


