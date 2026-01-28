# ProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]interface{}** | List of key, values pairs to be associated with the project | [optional] 
**CustomFieldsList** | Pointer to [**[]CustomArticleFieldAdd**](CustomArticleFieldAdd.md) | List of custom fields values, supersedes custom_fields parameter | [optional] 
**Description** | Pointer to **string** | Project description | [optional] 
**Funding** | Pointer to **string** | Grant number or organization(s) that funded this project. Up to 2000 characters permitted. | [optional] 
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**GroupId** | Pointer to **int64** | Only if project type is group. | [optional] 
**Title** | **string** | The title for this project - mandatory. 3 - 1000 characters. | 

## Methods

### NewProjectCreate

`func NewProjectCreate(title string, ) *ProjectCreate`

NewProjectCreate instantiates a new ProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateWithDefaults

`func NewProjectCreateWithDefaults() *ProjectCreate`

NewProjectCreateWithDefaults instantiates a new ProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *ProjectCreate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectCreate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectCreate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProjectCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *ProjectCreate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *ProjectCreate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *ProjectCreate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *ProjectCreate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunding

`func (o *ProjectCreate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ProjectCreate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ProjectCreate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *ProjectCreate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *ProjectCreate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ProjectCreate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ProjectCreate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *ProjectCreate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetGroupId

`func (o *ProjectCreate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProjectCreate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProjectCreate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProjectCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTitle

`func (o *ProjectCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectCreate) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


