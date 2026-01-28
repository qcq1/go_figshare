# ArticleEmbargo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbargoOptions** | **[]map[string]interface{}** | List of embargo permissions that are associated with the article. If the type is logged_in and the group_ids list is empty, then the whole institution can see the article; if there are multiple group_ids, then only users that are under those groups can see the article. | 
**EmbargoReason** | **string** | Reason for embargo | 
**EmbargoTitle** | **string** | Title for embargo | 
**IsEmbargoed** | **bool** | True if embargoed | 

## Methods

### NewArticleEmbargo

`func NewArticleEmbargo(embargoOptions []map[string]interface{}, embargoReason string, embargoTitle string, isEmbargoed bool, ) *ArticleEmbargo`

NewArticleEmbargo instantiates a new ArticleEmbargo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleEmbargoWithDefaults

`func NewArticleEmbargoWithDefaults() *ArticleEmbargo`

NewArticleEmbargoWithDefaults instantiates a new ArticleEmbargo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


