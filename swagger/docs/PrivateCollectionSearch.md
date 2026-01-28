# PrivateCollectionSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | only return collections with this resource_id | [optional] 
**Doi** | Pointer to **string** | Only return collections with this doi | [optional] 
**Handle** | Pointer to **string** | Only return collections with this handle | [optional] 
**Order** | Pointer to **string** | The field by which to order. | [optional] [default to "created_date"]
**ResourceDoi** | Pointer to **string** | Only return collections with this resource_doi | [optional] 
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

### NewPrivateCollectionSearch

`func NewPrivateCollectionSearch() *PrivateCollectionSearch`

NewPrivateCollectionSearch instantiates a new PrivateCollectionSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateCollectionSearchWithDefaults

`func NewPrivateCollectionSearchWithDefaults() *PrivateCollectionSearch`

NewPrivateCollectionSearchWithDefaults instantiates a new PrivateCollectionSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *PrivateCollectionSearch) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PrivateCollectionSearch) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PrivateCollectionSearch) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PrivateCollectionSearch) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetDoi

`func (o *PrivateCollectionSearch) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *PrivateCollectionSearch) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *PrivateCollectionSearch) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *PrivateCollectionSearch) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *PrivateCollectionSearch) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *PrivateCollectionSearch) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *PrivateCollectionSearch) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *PrivateCollectionSearch) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetOrder

`func (o *PrivateCollectionSearch) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PrivateCollectionSearch) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PrivateCollectionSearch) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PrivateCollectionSearch) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetResourceDoi

`func (o *PrivateCollectionSearch) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *PrivateCollectionSearch) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *PrivateCollectionSearch) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *PrivateCollectionSearch) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetGroup

`func (o *PrivateCollectionSearch) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PrivateCollectionSearch) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PrivateCollectionSearch) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PrivateCollectionSearch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInstitution

`func (o *PrivateCollectionSearch) GetInstitution() int32`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *PrivateCollectionSearch) GetInstitutionOk() (*int32, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *PrivateCollectionSearch) SetInstitution(v int32)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *PrivateCollectionSearch) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetLimit

`func (o *PrivateCollectionSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PrivateCollectionSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PrivateCollectionSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PrivateCollectionSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModifiedSince

`func (o *PrivateCollectionSearch) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *PrivateCollectionSearch) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *PrivateCollectionSearch) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *PrivateCollectionSearch) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetOffset

`func (o *PrivateCollectionSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PrivateCollectionSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PrivateCollectionSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PrivateCollectionSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderDirection

`func (o *PrivateCollectionSearch) GetOrderDirection() string`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *PrivateCollectionSearch) GetOrderDirectionOk() (*string, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *PrivateCollectionSearch) SetOrderDirection(v string)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *PrivateCollectionSearch) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPage

`func (o *PrivateCollectionSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PrivateCollectionSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PrivateCollectionSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *PrivateCollectionSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PrivateCollectionSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PrivateCollectionSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PrivateCollectionSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PrivateCollectionSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPublishedSince

`func (o *PrivateCollectionSearch) GetPublishedSince() string`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *PrivateCollectionSearch) GetPublishedSinceOk() (*string, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *PrivateCollectionSearch) SetPublishedSince(v string)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *PrivateCollectionSearch) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetSearchFor

`func (o *PrivateCollectionSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *PrivateCollectionSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *PrivateCollectionSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *PrivateCollectionSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


