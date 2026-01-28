# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | User id | 
**FirstName** | **string** | First Name | 
**LastName** | **string** | Last Name | 
**Name** | **string** | Full Name | 
**IsActive** | **bool** | Account activity status | 
**UrlName** | **string** | Name that appears in website url | 
**IsPublic** | **bool** | Account public status | 
**JobTitle** | **string** | User Job title | 
**OrcidId** | **string** | Orcid associated to this User | 

## Methods

### NewUser

`func NewUser(id int64, firstName string, lastName string, name string, isActive bool, urlName string, isPublic bool, jobTitle string, orcidId string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *User) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *User) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *User) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUrlName

`func (o *User) GetUrlName() string`

GetUrlName returns the UrlName field if non-nil, zero value otherwise.

### GetUrlNameOk

`func (o *User) GetUrlNameOk() (*string, bool)`

GetUrlNameOk returns a tuple with the UrlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlName

`func (o *User) SetUrlName(v string)`

SetUrlName sets UrlName field to given value.


### GetIsPublic

`func (o *User) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *User) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *User) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetJobTitle

`func (o *User) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *User) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *User) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### GetOrcidId

`func (o *User) GetOrcidId() string`

GetOrcidId returns the OrcidId field if non-nil, zero value otherwise.

### GetOrcidIdOk

`func (o *User) GetOrcidIdOk() (*string, bool)`

GetOrcidIdOk returns a tuple with the OrcidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcidId

`func (o *User) SetOrcidId(v string)`

SetOrcidId sets OrcidId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


