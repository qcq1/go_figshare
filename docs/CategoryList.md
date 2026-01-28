# CategoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasChildren** | **bool** | True if category has children | 
**IsSelectable** | **bool** | The selectable status | 
**Id** | **int64** | Category id | 
**ParentId** | **int64** | Parent category | 
**Path** | **string** | Path to all ancestor ids | 
**SourceId** | **string** | ID in original standard taxonomy | 
**TaxonomyId** | **int64** | Internal id of taxonomy the category is part of | 
**Title** | **string** | Category title | 

## Methods

### NewCategoryList

`func NewCategoryList(hasChildren bool, isSelectable bool, id int64, parentId int64, path string, sourceId string, taxonomyId int64, title string, ) *CategoryList`

NewCategoryList instantiates a new CategoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryListWithDefaults

`func NewCategoryListWithDefaults() *CategoryList`

NewCategoryListWithDefaults instantiates a new CategoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasChildren

`func (o *CategoryList) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *CategoryList) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *CategoryList) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.


### GetIsSelectable

`func (o *CategoryList) GetIsSelectable() bool`

GetIsSelectable returns the IsSelectable field if non-nil, zero value otherwise.

### GetIsSelectableOk

`func (o *CategoryList) GetIsSelectableOk() (*bool, bool)`

GetIsSelectableOk returns a tuple with the IsSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelectable

`func (o *CategoryList) SetIsSelectable(v bool)`

SetIsSelectable sets IsSelectable field to given value.


### GetId

`func (o *CategoryList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryList) SetId(v int64)`

SetId sets Id field to given value.


### GetParentId

`func (o *CategoryList) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CategoryList) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CategoryList) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetPath

`func (o *CategoryList) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CategoryList) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CategoryList) SetPath(v string)`

SetPath sets Path field to given value.


### GetSourceId

`func (o *CategoryList) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CategoryList) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CategoryList) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTaxonomyId

`func (o *CategoryList) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *CategoryList) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *CategoryList) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.


### GetTitle

`func (o *CategoryList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CategoryList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CategoryList) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


