# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Category id | 
**ParentId** | **int64** | Parent category | 
**Path** | **string** | Path to all ancestor ids | 
**SourceId** | **string** | ID in original standard taxonomy | 
**TaxonomyId** | **int64** | Internal id of taxonomy the category is part of | 
**Title** | **string** | Category title | 

## Methods

### NewCategory

`func NewCategory(id int64, parentId int64, path string, sourceId string, taxonomyId int64, title string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int64)`

SetId sets Id field to given value.


### GetParentId

`func (o *Category) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Category) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Category) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetPath

`func (o *Category) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Category) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Category) SetPath(v string)`

SetPath sets Path field to given value.


### GetSourceId

`func (o *Category) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Category) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Category) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTaxonomyId

`func (o *Category) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *Category) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *Category) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.


### GetTitle

`func (o *Category) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Category) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Category) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


