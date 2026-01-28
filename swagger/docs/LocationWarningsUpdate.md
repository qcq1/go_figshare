# LocationWarningsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Url for entity | 
**Warnings** | **[]string** | Issues encountered during the operation | 

## Methods

### NewLocationWarningsUpdate

`func NewLocationWarningsUpdate(location string, warnings []string, ) *LocationWarningsUpdate`

NewLocationWarningsUpdate instantiates a new LocationWarningsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWarningsUpdateWithDefaults

`func NewLocationWarningsUpdateWithDefaults() *LocationWarningsUpdate`

NewLocationWarningsUpdateWithDefaults instantiates a new LocationWarningsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *LocationWarningsUpdate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationWarningsUpdate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationWarningsUpdate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetWarnings

`func (o *LocationWarningsUpdate) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *LocationWarningsUpdate) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *LocationWarningsUpdate) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


