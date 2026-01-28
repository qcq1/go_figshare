# CollectionSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDoi** | Pointer to **string** | Only return collections with this resource_doi | [optional] 
**Doi** | Pointer to **string** | Only return collections with this doi | [optional] 
**Handle** | Pointer to **string** | Only return collections with this handle | [optional] 
**Order** | Pointer to **string** | The field by which to order. | [optional] [default to "created_date"]
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

### NewCollectionSearch

`func NewCollectionSearch() *CollectionSearch`

NewCollectionSearch instantiates a new CollectionSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionSearchWithDefaults

`func NewCollectionSearchWithDefaults() *CollectionSearch`

NewCollectionSearchWithDefaults instantiates a new CollectionSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDoi

`func (o *CollectionSearch) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *CollectionSearch) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *CollectionSearch) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *CollectionSearch) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetDoi

`func (o *CollectionSearch) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *CollectionSearch) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *CollectionSearch) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *CollectionSearch) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *CollectionSearch) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CollectionSearch) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CollectionSearch) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CollectionSearch) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetOrder

`func (o *CollectionSearch) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CollectionSearch) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CollectionSearch) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CollectionSearch) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSearchFor

`func (o *CollectionSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *CollectionSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *CollectionSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *CollectionSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.

### GetPage

`func (o *CollectionSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CollectionSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CollectionSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *CollectionSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CollectionSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CollectionSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CollectionSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CollectionSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetLimit

`func (o *CollectionSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CollectionSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CollectionSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CollectionSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CollectionSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CollectionSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CollectionSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CollectionSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderDirection

`func (o *CollectionSearch) GetOrderDirection() string`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *CollectionSearch) GetOrderDirectionOk() (*string, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *CollectionSearch) SetOrderDirection(v string)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *CollectionSearch) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetInstitution

`func (o *CollectionSearch) GetInstitution() int32`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *CollectionSearch) GetInstitutionOk() (*int32, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *CollectionSearch) SetInstitution(v int32)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *CollectionSearch) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetPublishedSince

`func (o *CollectionSearch) GetPublishedSince() string`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *CollectionSearch) GetPublishedSinceOk() (*string, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *CollectionSearch) SetPublishedSince(v string)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *CollectionSearch) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetModifiedSince

`func (o *CollectionSearch) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *CollectionSearch) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *CollectionSearch) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *CollectionSearch) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetGroup

`func (o *CollectionSearch) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CollectionSearch) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CollectionSearch) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CollectionSearch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


