# ProjectCollaborator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of collaborator invitation | 
**RoleName** | **string** | Collaborator role | 
**UserId** | **int32** | Collaborator id | 
**Name** | **string** | Collaborator name | 

## Methods

### NewProjectCollaborator

`func NewProjectCollaborator(status string, roleName string, userId int32, name string, ) *ProjectCollaborator`

NewProjectCollaborator instantiates a new ProjectCollaborator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCollaboratorWithDefaults

`func NewProjectCollaboratorWithDefaults() *ProjectCollaborator`

NewProjectCollaboratorWithDefaults instantiates a new ProjectCollaborator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectCollaborator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectCollaborator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectCollaborator) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRoleName

`func (o *ProjectCollaborator) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ProjectCollaborator) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ProjectCollaborator) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUserId

`func (o *ProjectCollaborator) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProjectCollaborator) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProjectCollaborator) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *ProjectCollaborator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCollaborator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCollaborator) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


