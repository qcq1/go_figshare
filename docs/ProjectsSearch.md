# ProjectsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **string** | The field by which to order. | [optional] [default to "published_date"]
**Group** | Pointer to **int32** | only return collections from this group | [optional] 
**Institution** | Pointer to **int32** | only return collections from this institution | [optional] 
**Limit** | Pointer to **int64** | Number of results included on a page. Used for pagination with query | [optional] 
**ModifiedSince** | Pointer to **string** | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] 
**Offset** | Pointer to **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | [optional] 
**OrderDirection** | Pointer to **string** | Direction of ordering | [optional] [default to "desc"]
**Page** | Pointer to **int64** | Page number. Used for pagination with page_size | [optional] 
**PageSize** | Pointer to **int64** | The number of results included on a page. Used for pagination with page | [optional] [default to 10]
**PublishedSince** | Pointer to **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] 
**SearchFor** | Pointer to **string** | Search term | [optional] 

## Methods

### NewProjectsSearch

`func NewProjectsSearch() *ProjectsSearch`

NewProjectsSearch instantiates a new ProjectsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSearchWithDefaults

`func NewProjectsSearchWithDefaults() *ProjectsSearch`

NewProjectsSearchWithDefaults instantiates a new ProjectsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ProjectsSearch) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ProjectsSearch) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ProjectsSearch) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ProjectsSearch) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetGroup

`func (o *ProjectsSearch) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProjectsSearch) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProjectsSearch) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ProjectsSearch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInstitution

`func (o *ProjectsSearch) GetInstitution() int32`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *ProjectsSearch) GetInstitutionOk() (*int32, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *ProjectsSearch) SetInstitution(v int32)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *ProjectsSearch) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetLimit

`func (o *ProjectsSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProjectsSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProjectsSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ProjectsSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModifiedSince

`func (o *ProjectsSearch) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *ProjectsSearch) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *ProjectsSearch) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *ProjectsSearch) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetOffset

`func (o *ProjectsSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ProjectsSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ProjectsSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ProjectsSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderDirection

`func (o *ProjectsSearch) GetOrderDirection() string`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *ProjectsSearch) GetOrderDirectionOk() (*string, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *ProjectsSearch) SetOrderDirection(v string)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *ProjectsSearch) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPage

`func (o *ProjectsSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProjectsSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProjectsSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProjectsSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ProjectsSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProjectsSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProjectsSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProjectsSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPublishedSince

`func (o *ProjectsSearch) GetPublishedSince() string`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *ProjectsSearch) GetPublishedSinceOk() (*string, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *ProjectsSearch) SetPublishedSince(v string)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *ProjectsSearch) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetSearchFor

`func (o *ProjectsSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *ProjectsSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *ProjectsSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *ProjectsSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


