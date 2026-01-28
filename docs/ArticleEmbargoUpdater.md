# ArticleEmbargoUpdater

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEmbargoed** | **bool** | Embargo status | 
**EmbargoDate** | **string** | Date when the embargo expires and the article gets published, &#39;0&#39; value will set up permanent embargo | 
**EmbargoType** | **string** | Embargo can be enabled at the article or the file level. Possible values: article, file | 
**EmbargoTitle** | Pointer to **string** | Title for embargo | [optional] 
**EmbargoReason** | Pointer to **string** | Reason for setting embargo | [optional] 
**EmbargoOptions** | Pointer to **[]map[string]interface{}** | List of embargo permissions to be associated with the article. The list must contain &#x60;id&#x60; and can also contain &#x60;group_ids&#x60;(a field that only applies to &#39;logged_in&#39; permissions). The new list replaces old options in the database, and an empty list removes all permissions for this article. Administration permission has to be set up alone but logged in and IP range permissions can be set up together. | [optional] 

## Methods

### NewArticleEmbargoUpdater

`func NewArticleEmbargoUpdater(isEmbargoed bool, embargoDate string, embargoType string, ) *ArticleEmbargoUpdater`

NewArticleEmbargoUpdater instantiates a new ArticleEmbargoUpdater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleEmbargoUpdaterWithDefaults

`func NewArticleEmbargoUpdaterWithDefaults() *ArticleEmbargoUpdater`

NewArticleEmbargoUpdaterWithDefaults instantiates a new ArticleEmbargoUpdater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEmbargoed

`func (o *ArticleEmbargoUpdater) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *ArticleEmbargoUpdater) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *ArticleEmbargoUpdater) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetEmbargoDate

`func (o *ArticleEmbargoUpdater) GetEmbargoDate() string`

GetEmbargoDate returns the EmbargoDate field if non-nil, zero value otherwise.

### GetEmbargoDateOk

`func (o *ArticleEmbargoUpdater) GetEmbargoDateOk() (*string, bool)`

GetEmbargoDateOk returns a tuple with the EmbargoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoDate

`func (o *ArticleEmbargoUpdater) SetEmbargoDate(v string)`

SetEmbargoDate sets EmbargoDate field to given value.


### GetEmbargoType

`func (o *ArticleEmbargoUpdater) GetEmbargoType() string`

GetEmbargoType returns the EmbargoType field if non-nil, zero value otherwise.

### GetEmbargoTypeOk

`func (o *ArticleEmbargoUpdater) GetEmbargoTypeOk() (*string, bool)`

GetEmbargoTypeOk returns a tuple with the EmbargoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoType

`func (o *ArticleEmbargoUpdater) SetEmbargoType(v string)`

SetEmbargoType sets EmbargoType field to given value.


### GetEmbargoTitle

`func (o *ArticleEmbargoUpdater) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *ArticleEmbargoUpdater) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *ArticleEmbargoUpdater) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.

### HasEmbargoTitle

`func (o *ArticleEmbargoUpdater) HasEmbargoTitle() bool`

HasEmbargoTitle returns a boolean if a field has been set.

### GetEmbargoReason

`func (o *ArticleEmbargoUpdater) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *ArticleEmbargoUpdater) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *ArticleEmbargoUpdater) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.

### HasEmbargoReason

`func (o *ArticleEmbargoUpdater) HasEmbargoReason() bool`

HasEmbargoReason returns a boolean if a field has been set.

### GetEmbargoOptions

`func (o *ArticleEmbargoUpdater) GetEmbargoOptions() []map[string]interface{}`

GetEmbargoOptions returns the EmbargoOptions field if non-nil, zero value otherwise.

### GetEmbargoOptionsOk

`func (o *ArticleEmbargoUpdater) GetEmbargoOptionsOk() (*[]map[string]interface{}, bool)`

GetEmbargoOptionsOk returns a tuple with the EmbargoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoOptions

`func (o *ArticleEmbargoUpdater) SetEmbargoOptions(v []map[string]interface{})`

SetEmbargoOptions sets EmbargoOptions field to given value.

### HasEmbargoOptions

`func (o *ArticleEmbargoUpdater) HasEmbargoOptions() bool`

HasEmbargoOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


