# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCriteria** | **string** | HR code associated with group, if code exists | 
**Id** | **int64** | Group id | 
**Name** | **string** | Group name | 
**ParentId** | **int64** | Parent group if any | 
**ResourceId** | **string** | Group resource id | 

## Methods

### NewGroup

`func NewGroup(associationCriteria string, id int64, name string, parentId int64, resourceId string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCriteria

`func (o *Group) GetAssociationCriteria() string`

GetAssociationCriteria returns the AssociationCriteria field if non-nil, zero value otherwise.

### GetAssociationCriteriaOk

`func (o *Group) GetAssociationCriteriaOk() (*string, bool)`

GetAssociationCriteriaOk returns a tuple with the AssociationCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCriteria

`func (o *Group) SetAssociationCriteria(v string)`

SetAssociationCriteria sets AssociationCriteria field to given value.


### GetId

`func (o *Group) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Group) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Group) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Group) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetResourceId

`func (o *Group) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Group) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Group) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


