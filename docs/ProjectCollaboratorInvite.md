# ProjectCollaboratorInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Text sent when inviting the user to the project | [optional] 
**Email** | Pointer to **string** | Collaborator email | [optional] 
**RoleName** | **string** | Role of the the collaborator inside the project | 
**UserId** | Pointer to **int64** | User id of the collaborator | [optional] 

## Methods

### NewProjectCollaboratorInvite

`func NewProjectCollaboratorInvite(roleName string, ) *ProjectCollaboratorInvite`

NewProjectCollaboratorInvite instantiates a new ProjectCollaboratorInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCollaboratorInviteWithDefaults

`func NewProjectCollaboratorInviteWithDefaults() *ProjectCollaboratorInvite`

NewProjectCollaboratorInviteWithDefaults instantiates a new ProjectCollaboratorInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ProjectCollaboratorInvite) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ProjectCollaboratorInvite) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ProjectCollaboratorInvite) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ProjectCollaboratorInvite) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEmail

`func (o *ProjectCollaboratorInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProjectCollaboratorInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProjectCollaboratorInvite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProjectCollaboratorInvite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRoleName

`func (o *ProjectCollaboratorInvite) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ProjectCollaboratorInvite) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ProjectCollaboratorInvite) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUserId

`func (o *ProjectCollaboratorInvite) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProjectCollaboratorInvite) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProjectCollaboratorInvite) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ProjectCollaboratorInvite) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


