# ArticleSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDoi** | Pointer to **string** | Only return articles with this resource_doi | [optional] 
**ItemType** | Pointer to **int64** | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | [optional] 
**Doi** | Pointer to **string** | Only return articles with this doi | [optional] 
**Handle** | Pointer to **string** | Only return articles with this handle | [optional] 
**ProjectId** | Pointer to **int64** | Only return articles in this project | [optional] 
**Order** | Pointer to **string** | The field by which to order | [optional] [default to "created_date"]
**SearchFor** | Pointer to **string** | Search term | [optional] 
**Page** | Pointer to **int64** | Page number. Used for pagination with page_size | [optional] 
**PageSize** | Pointer to **int64** | The number of results included on a page. Used for pagination with page | [optional] [default to 10]
**Limit** | Pointer to **int64** | Number of results included on a page. Used for pagination with query | [optional] 
**Offset** | Pointer to **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | [optional] 
**OrderDirection** | Pointer to **string** | Direction of ordering | [optional] [default to "desc"]
**Institution** | Pointer to **int32** | only return collections from this institution | [optional] 
**PublishedSince** | Pointer to **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] 
**ModifiedSince** | Pointer to **string** | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] 
**Group** | Pointer to **int32** | only return collections from this group | [optional] 

## Methods

### NewArticleSearch

`func NewArticleSearch() *ArticleSearch`

NewArticleSearch instantiates a new ArticleSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleSearchWithDefaults

`func NewArticleSearchWithDefaults() *ArticleSearch`

NewArticleSearchWithDefaults instantiates a new ArticleSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDoi

`func (o *ArticleSearch) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ArticleSearch) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ArticleSearch) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *ArticleSearch) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetItemType

`func (o *ArticleSearch) GetItemType() int64`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ArticleSearch) GetItemTypeOk() (*int64, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ArticleSearch) SetItemType(v int64)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ArticleSearch) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetDoi

`func (o *ArticleSearch) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ArticleSearch) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ArticleSearch) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *ArticleSearch) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *ArticleSearch) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ArticleSearch) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ArticleSearch) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *ArticleSearch) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetProjectId

`func (o *ArticleSearch) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ArticleSearch) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ArticleSearch) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ArticleSearch) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetOrder

`func (o *ArticleSearch) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ArticleSearch) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ArticleSearch) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ArticleSearch) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSearchFor

`func (o *ArticleSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *ArticleSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *ArticleSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *ArticleSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.

### GetPage

`func (o *ArticleSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ArticleSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ArticleSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ArticleSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ArticleSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ArticleSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ArticleSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ArticleSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetLimit

`func (o *ArticleSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ArticleSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ArticleSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ArticleSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ArticleSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ArticleSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ArticleSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ArticleSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderDirection

`func (o *ArticleSearch) GetOrderDirection() string`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *ArticleSearch) GetOrderDirectionOk() (*string, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *ArticleSearch) SetOrderDirection(v string)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *ArticleSearch) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetInstitution

`func (o *ArticleSearch) GetInstitution() int32`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *ArticleSearch) GetInstitutionOk() (*int32, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *ArticleSearch) SetInstitution(v int32)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *ArticleSearch) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetPublishedSince

`func (o *ArticleSearch) GetPublishedSince() string`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *ArticleSearch) GetPublishedSinceOk() (*string, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *ArticleSearch) SetPublishedSince(v string)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *ArticleSearch) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetModifiedSince

`func (o *ArticleSearch) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *ArticleSearch) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *ArticleSearch) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *ArticleSearch) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetGroup

`func (o *ArticleSearch) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ArticleSearch) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ArticleSearch) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ArticleSearch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


