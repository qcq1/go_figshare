# AuthorComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | **int64** | Institution id | 
**GroupId** | **int64** | Group id | 
**FirstName** | **string** | First Name | 
**LastName** | **string** | Last Name | 
**IsPublic** | **int64** | if 1 then the author has published items | 
**JobTitle** | **string** | Job title | 
**Id** | **int64** | Author id | 
**FullName** | **string** | Author full name | 
**IsActive** | **bool** | True if author has published items | 
**UrlName** | **string** | Author url name | 
**OrcidId** | **string** | Author Orcid | 

## Methods

### NewAuthorComplete

`func NewAuthorComplete(institutionId int64, groupId int64, firstName string, lastName string, isPublic int64, jobTitle string, id int64, fullName string, isActive bool, urlName string, orcidId string, ) *AuthorComplete`

NewAuthorComplete instantiates a new AuthorComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorCompleteWithDefaults

`func NewAuthorCompleteWithDefaults() *AuthorComplete`

NewAuthorCompleteWithDefaults instantiates a new AuthorComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *AuthorComplete) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *AuthorComplete) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *AuthorComplete) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.


### GetGroupId

`func (o *AuthorComplete) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AuthorComplete) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AuthorComplete) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetFirstName

`func (o *AuthorComplete) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AuthorComplete) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AuthorComplete) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AuthorComplete) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AuthorComplete) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AuthorComplete) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetIsPublic

`func (o *AuthorComplete) GetIsPublic() int64`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AuthorComplete) GetIsPublicOk() (*int64, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AuthorComplete) SetIsPublic(v int64)`

SetIsPublic sets IsPublic field to given value.


### GetJobTitle

`func (o *AuthorComplete) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *AuthorComplete) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *AuthorComplete) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### GetId

`func (o *AuthorComplete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorComplete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorComplete) SetId(v int64)`

SetId sets Id field to given value.


### GetFullName

`func (o *AuthorComplete) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorComplete) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorComplete) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsActive

`func (o *AuthorComplete) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AuthorComplete) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AuthorComplete) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUrlName

`func (o *AuthorComplete) GetUrlName() string`

GetUrlName returns the UrlName field if non-nil, zero value otherwise.

### GetUrlNameOk

`func (o *AuthorComplete) GetUrlNameOk() (*string, bool)`

GetUrlNameOk returns a tuple with the UrlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlName

`func (o *AuthorComplete) SetUrlName(v string)`

SetUrlName sets UrlName field to given value.


### GetOrcidId

`func (o *AuthorComplete) GetOrcidId() string`

GetOrcidId returns the OrcidId field if non-nil, zero value otherwise.

### GetOrcidIdOk

`func (o *AuthorComplete) GetOrcidIdOk() (*string, bool)`

GetOrcidIdOk returns a tuple with the OrcidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcidId

`func (o *AuthorComplete) SetOrcidId(v string)`

SetOrcidId sets OrcidId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


