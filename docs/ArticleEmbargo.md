# ArticleEmbargo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEmbargoed** | **bool** | True if embargoed | 
**EmbargoDate** | **NullableString** | Date when embargo lifts | 
**EmbargoType** | **NullableString** | Embargo type | 
**EmbargoTitle** | **string** | Title for embargo | 
**EmbargoReason** | **string** | Reason for embargo | 
**EmbargoOptions** | **[]map[string]interface{}** | List of embargo permissions that are associated with the article. If the type is logged_in and the group_ids list is empty, then the whole institution can see the article; if there are multiple group_ids, then only users that are under those groups can see the article. | 

## Methods

### NewArticleEmbargo

`func NewArticleEmbargo(isEmbargoed bool, embargoDate NullableString, embargoType NullableString, embargoTitle string, embargoReason string, embargoOptions []map[string]interface{}, ) *ArticleEmbargo`

NewArticleEmbargo instantiates a new ArticleEmbargo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleEmbargoWithDefaults

`func NewArticleEmbargoWithDefaults() *ArticleEmbargo`

NewArticleEmbargoWithDefaults instantiates a new ArticleEmbargo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEmbargoed

`func (o *ArticleEmbargo) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *ArticleEmbargo) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *ArticleEmbargo) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetEmbargoDate

`func (o *ArticleEmbargo) GetEmbargoDate() string`

GetEmbargoDate returns the EmbargoDate field if non-nil, zero value otherwise.

### GetEmbargoDateOk

`func (o *ArticleEmbargo) GetEmbargoDateOk() (*string, bool)`

GetEmbargoDateOk returns a tuple with the EmbargoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoDate

`func (o *ArticleEmbargo) SetEmbargoDate(v string)`

SetEmbargoDate sets EmbargoDate field to given value.


### SetEmbargoDateNil

`func (o *ArticleEmbargo) SetEmbargoDateNil(b bool)`

 SetEmbargoDateNil sets the value for EmbargoDate to be an explicit nil

### UnsetEmbargoDate
`func (o *ArticleEmbargo) UnsetEmbargoDate()`

UnsetEmbargoDate ensures that no value is present for EmbargoDate, not even an explicit nil
### GetEmbargoType

`func (o *ArticleEmbargo) GetEmbargoType() string`

GetEmbargoType returns the EmbargoType field if non-nil, zero value otherwise.

### GetEmbargoTypeOk

`func (o *ArticleEmbargo) GetEmbargoTypeOk() (*string, bool)`

GetEmbargoTypeOk returns a tuple with the EmbargoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoType

`func (o *ArticleEmbargo) SetEmbargoType(v string)`

SetEmbargoType sets EmbargoType field to given value.


### SetEmbargoTypeNil

`func (o *ArticleEmbargo) SetEmbargoTypeNil(b bool)`

 SetEmbargoTypeNil sets the value for EmbargoType to be an explicit nil

### UnsetEmbargoType
`func (o *ArticleEmbargo) UnsetEmbargoType()`

UnsetEmbargoType ensures that no value is present for EmbargoType, not even an explicit nil
### GetEmbargoTitle

`func (o *ArticleEmbargo) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *ArticleEmbargo) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *ArticleEmbargo) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.


### GetEmbargoReason

`func (o *ArticleEmbargo) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *ArticleEmbargo) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *ArticleEmbargo) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.


### GetEmbargoOptions

`func (o *ArticleEmbargo) GetEmbargoOptions() []map[string]interface{}`

GetEmbargoOptions returns the EmbargoOptions field if non-nil, zero value otherwise.

### GetEmbargoOptionsOk

`func (o *ArticleEmbargo) GetEmbargoOptionsOk() (*[]map[string]interface{}, bool)`

GetEmbargoOptionsOk returns a tuple with the EmbargoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoOptions

`func (o *ArticleEmbargo) SetEmbargoOptions(v []map[string]interface{})`

SetEmbargoOptions sets EmbargoOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


