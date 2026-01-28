# Institution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Institution id | 
**Name** | **string** | Institution name | 
**Domain** | **NullableString** | Institution domain | 

## Methods

### NewInstitution

`func NewInstitution(id int64, name string, domain NullableString, ) *Institution`

NewInstitution instantiates a new Institution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionWithDefaults

`func NewInstitutionWithDefaults() *Institution`

NewInstitutionWithDefaults instantiates a new Institution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Institution) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Institution) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Institution) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Institution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Institution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Institution) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *Institution) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Institution) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Institution) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *Institution) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Institution) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


