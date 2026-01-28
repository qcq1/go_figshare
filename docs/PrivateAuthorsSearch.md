# PrivateAuthorsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int64** | Return only authors in this group or subgroups of the group | [optional] 
**InstitutionId** | Pointer to **int64** | Return only authors associated to this institution | [optional] 
**IsActive** | Pointer to **bool** | Return only active authors if True | [optional] 
**IsPublic** | Pointer to **bool** | Return only authors that have published items if True | [optional] 
**Limit** | Pointer to **int64** | Number of results included on a page. Used for pagination with query | [optional] 
**Offset** | Pointer to **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | [optional] 
**Orcid** | Pointer to **string** | Orcid of author | [optional] 
**Page** | Pointer to **int64** | Page number. Used for pagination with page_size | [optional] 
**PageSize** | Pointer to **int64** | The number of results included on a page. Used for pagination with page | [optional] [default to 10]
**SearchFor** | Pointer to **string** | Search term | [optional] 

## Methods

### NewPrivateAuthorsSearch

`func NewPrivateAuthorsSearch() *PrivateAuthorsSearch`

NewPrivateAuthorsSearch instantiates a new PrivateAuthorsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateAuthorsSearchWithDefaults

`func NewPrivateAuthorsSearchWithDefaults() *PrivateAuthorsSearch`

NewPrivateAuthorsSearchWithDefaults instantiates a new PrivateAuthorsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *PrivateAuthorsSearch) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PrivateAuthorsSearch) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PrivateAuthorsSearch) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PrivateAuthorsSearch) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetInstitutionId

`func (o *PrivateAuthorsSearch) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *PrivateAuthorsSearch) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *PrivateAuthorsSearch) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *PrivateAuthorsSearch) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### GetIsActive

`func (o *PrivateAuthorsSearch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PrivateAuthorsSearch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PrivateAuthorsSearch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PrivateAuthorsSearch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsPublic

`func (o *PrivateAuthorsSearch) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PrivateAuthorsSearch) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PrivateAuthorsSearch) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *PrivateAuthorsSearch) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLimit

`func (o *PrivateAuthorsSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PrivateAuthorsSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PrivateAuthorsSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PrivateAuthorsSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PrivateAuthorsSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PrivateAuthorsSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PrivateAuthorsSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PrivateAuthorsSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrcid

`func (o *PrivateAuthorsSearch) GetOrcid() string`

GetOrcid returns the Orcid field if non-nil, zero value otherwise.

### GetOrcidOk

`func (o *PrivateAuthorsSearch) GetOrcidOk() (*string, bool)`

GetOrcidOk returns a tuple with the Orcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcid

`func (o *PrivateAuthorsSearch) SetOrcid(v string)`

SetOrcid sets Orcid field to given value.

### HasOrcid

`func (o *PrivateAuthorsSearch) HasOrcid() bool`

HasOrcid returns a boolean if a field has been set.

### GetPage

`func (o *PrivateAuthorsSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PrivateAuthorsSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PrivateAuthorsSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *PrivateAuthorsSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PrivateAuthorsSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PrivateAuthorsSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PrivateAuthorsSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PrivateAuthorsSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSearchFor

`func (o *PrivateAuthorsSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *PrivateAuthorsSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *PrivateAuthorsSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *PrivateAuthorsSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


