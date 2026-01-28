# ProjectCompletePrivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | ID of the account owning the project | 
**Collaborators** | [**[]Collaborator**](Collaborator.md) | List of project collaborators | 
**CreatedDate** | **string** | Date when project was created | 
**CustomFields** | [**[]ShortCustomField**](ShortCustomField.md) | Collection custom fields | 
**Description** | **string** | Project description | 
**Funding** | **string** | Project funding | 
**FundingList** | [**[]FundingInformation**](FundingInformation.md) | Full Project funding information | 
**GroupId** | **int64** | Group of project if any | 
**ModifiedDate** | **string** | Date when project was last modified | 
**Quota** | **int64** | Project quota | 
**UsedQuota** | **int64** | Project used quota | 
**UsedQuotaPrivate** | **int64** | Project private quota used | 
**UsedQuotaPublic** | **int64** | Project public quota used | 
**Role** | **string** | Role inside this project | 
**Storage** | **string** | Project storage type | 
**Id** | **int64** | Project id | 
**Title** | **string** | Project title | 
**Url** | **string** | Api endpoint | 

## Methods

### NewProjectCompletePrivate

`func NewProjectCompletePrivate(accountId int64, collaborators []Collaborator, createdDate string, customFields []ShortCustomField, description string, funding string, fundingList []FundingInformation, groupId int64, modifiedDate string, quota int64, usedQuota int64, usedQuotaPrivate int64, usedQuotaPublic int64, role string, storage string, id int64, title string, url string, ) *ProjectCompletePrivate`

NewProjectCompletePrivate instantiates a new ProjectCompletePrivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCompletePrivateWithDefaults

`func NewProjectCompletePrivateWithDefaults() *ProjectCompletePrivate`

NewProjectCompletePrivateWithDefaults instantiates a new ProjectCompletePrivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ProjectCompletePrivate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ProjectCompletePrivate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ProjectCompletePrivate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetCollaborators

`func (o *ProjectCompletePrivate) GetCollaborators() []Collaborator`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *ProjectCompletePrivate) GetCollaboratorsOk() (*[]Collaborator, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *ProjectCompletePrivate) SetCollaborators(v []Collaborator)`

SetCollaborators sets Collaborators field to given value.


### GetCreatedDate

`func (o *ProjectCompletePrivate) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectCompletePrivate) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectCompletePrivate) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetCustomFields

`func (o *ProjectCompletePrivate) GetCustomFields() []ShortCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectCompletePrivate) GetCustomFieldsOk() (*[]ShortCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectCompletePrivate) SetCustomFields(v []ShortCustomField)`

SetCustomFields sets CustomFields field to given value.


### GetDescription

`func (o *ProjectCompletePrivate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCompletePrivate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCompletePrivate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFunding

`func (o *ProjectCompletePrivate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ProjectCompletePrivate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ProjectCompletePrivate) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *ProjectCompletePrivate) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ProjectCompletePrivate) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ProjectCompletePrivate) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetGroupId

`func (o *ProjectCompletePrivate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProjectCompletePrivate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProjectCompletePrivate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetModifiedDate

`func (o *ProjectCompletePrivate) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectCompletePrivate) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectCompletePrivate) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetQuota

`func (o *ProjectCompletePrivate) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ProjectCompletePrivate) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ProjectCompletePrivate) SetQuota(v int64)`

SetQuota sets Quota field to given value.


### GetUsedQuota

`func (o *ProjectCompletePrivate) GetUsedQuota() int64`

GetUsedQuota returns the UsedQuota field if non-nil, zero value otherwise.

### GetUsedQuotaOk

`func (o *ProjectCompletePrivate) GetUsedQuotaOk() (*int64, bool)`

GetUsedQuotaOk returns a tuple with the UsedQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuota

`func (o *ProjectCompletePrivate) SetUsedQuota(v int64)`

SetUsedQuota sets UsedQuota field to given value.


### GetUsedQuotaPrivate

`func (o *ProjectCompletePrivate) GetUsedQuotaPrivate() int64`

GetUsedQuotaPrivate returns the UsedQuotaPrivate field if non-nil, zero value otherwise.

### GetUsedQuotaPrivateOk

`func (o *ProjectCompletePrivate) GetUsedQuotaPrivateOk() (*int64, bool)`

GetUsedQuotaPrivateOk returns a tuple with the UsedQuotaPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaPrivate

`func (o *ProjectCompletePrivate) SetUsedQuotaPrivate(v int64)`

SetUsedQuotaPrivate sets UsedQuotaPrivate field to given value.


### GetUsedQuotaPublic

`func (o *ProjectCompletePrivate) GetUsedQuotaPublic() int64`

GetUsedQuotaPublic returns the UsedQuotaPublic field if non-nil, zero value otherwise.

### GetUsedQuotaPublicOk

`func (o *ProjectCompletePrivate) GetUsedQuotaPublicOk() (*int64, bool)`

GetUsedQuotaPublicOk returns a tuple with the UsedQuotaPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaPublic

`func (o *ProjectCompletePrivate) SetUsedQuotaPublic(v int64)`

SetUsedQuotaPublic sets UsedQuotaPublic field to given value.


### GetRole

`func (o *ProjectCompletePrivate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectCompletePrivate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectCompletePrivate) SetRole(v string)`

SetRole sets Role field to given value.


### GetStorage

`func (o *ProjectCompletePrivate) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectCompletePrivate) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectCompletePrivate) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetId

`func (o *ProjectCompletePrivate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCompletePrivate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCompletePrivate) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *ProjectCompletePrivate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectCompletePrivate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectCompletePrivate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *ProjectCompletePrivate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectCompletePrivate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectCompletePrivate) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


