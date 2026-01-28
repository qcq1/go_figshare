# LocationWarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **int64** | Figshare ID of the entity | 
**Location** | **string** | Url for entity | 
**Warnings** | **[]string** | Issues encountered during the operation | 

## Methods

### NewLocationWarnings

`func NewLocationWarnings(entityId int64, location string, warnings []string, ) *LocationWarnings`

NewLocationWarnings instantiates a new LocationWarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWarningsWithDefaults

`func NewLocationWarningsWithDefaults() *LocationWarnings`

NewLocationWarningsWithDefaults instantiates a new LocationWarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *LocationWarnings) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *LocationWarnings) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *LocationWarnings) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.


### GetLocation

`func (o *LocationWarnings) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationWarnings) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationWarnings) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetWarnings

`func (o *LocationWarnings) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *LocationWarnings) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *LocationWarnings) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


