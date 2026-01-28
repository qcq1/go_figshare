# ProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]interface{}** | List of key, values pairs to be associated with the project | [optional] 
**CustomFieldsList** | Pointer to [**[]CustomArticleFieldAdd**](CustomArticleFieldAdd.md) | List of custom fields values, supersedes custom_fields parameter | [optional] 
**Description** | Pointer to **string** | Project description | [optional] 
**Funding** | Pointer to **string** | Grant number or organization(s) that funded this project. Up to 2000 characters permitted. | [optional] 
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**Title** | Pointer to **string** | The title for this project - mandatory. 3 - 1000 characters. | [optional] 

## Methods

### NewProjectUpdate

`func NewProjectUpdate() *ProjectUpdate`

NewProjectUpdate instantiates a new ProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateWithDefaults

`func NewProjectUpdateWithDefaults() *ProjectUpdate`

NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *ProjectUpdate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectUpdate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectUpdate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProjectUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *ProjectUpdate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *ProjectUpdate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *ProjectUpdate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *ProjectUpdate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunding

`func (o *ProjectUpdate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ProjectUpdate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ProjectUpdate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *ProjectUpdate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *ProjectUpdate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ProjectUpdate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ProjectUpdate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *ProjectUpdate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetTitle

`func (o *ProjectUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProjectUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


