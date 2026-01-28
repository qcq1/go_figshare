# CreateProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **int64** | Figshare ID of the entity | 
**Location** | **string** | Url for entity | 

## Methods

### NewCreateProjectResponse

`func NewCreateProjectResponse(entityId int64, location string, ) *CreateProjectResponse`

NewCreateProjectResponse instantiates a new CreateProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectResponseWithDefaults

`func NewCreateProjectResponseWithDefaults() *CreateProjectResponse`

NewCreateProjectResponseWithDefaults instantiates a new CreateProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *CreateProjectResponse) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CreateProjectResponse) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CreateProjectResponse) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.


### GetLocation

`func (o *CreateProjectResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateProjectResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateProjectResponse) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


