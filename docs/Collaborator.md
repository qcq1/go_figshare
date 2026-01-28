# Collaborator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | **string** | Collaborator role | 
**UserId** | **int32** | Collaborator id | 
**Name** | **string** | Collaborator name | 

## Methods

### NewCollaborator

`func NewCollaborator(roleName string, userId int32, name string, ) *Collaborator`

NewCollaborator instantiates a new Collaborator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollaboratorWithDefaults

`func NewCollaboratorWithDefaults() *Collaborator`

NewCollaboratorWithDefaults instantiates a new Collaborator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *Collaborator) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Collaborator) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Collaborator) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUserId

`func (o *Collaborator) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Collaborator) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Collaborator) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *Collaborator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collaborator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collaborator) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


