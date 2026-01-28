# InstitutionAccountsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFor** | Pointer to **string** | Search term | [optional] 
**IsActive** | Pointer to **int64** | Filter by active status | [optional] 
**Page** | Pointer to **int64** | Page number. Used for pagination with page_size | [optional] 
**PageSize** | Pointer to **int64** | The number of results included on a page. Used for pagination with page | [optional] [default to 10]
**Limit** | Pointer to **int64** | Number of results included on a page. Used for pagination with query | [optional] 
**Offset** | Pointer to **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | [optional] 
**InstitutionUserId** | Pointer to **string** | filter by institution_user_id | [optional] 
**Email** | Pointer to **string** | filter by email | [optional] 

## Methods

### NewInstitutionAccountsSearch

`func NewInstitutionAccountsSearch() *InstitutionAccountsSearch`

NewInstitutionAccountsSearch instantiates a new InstitutionAccountsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionAccountsSearchWithDefaults

`func NewInstitutionAccountsSearchWithDefaults() *InstitutionAccountsSearch`

NewInstitutionAccountsSearchWithDefaults instantiates a new InstitutionAccountsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFor

`func (o *InstitutionAccountsSearch) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *InstitutionAccountsSearch) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *InstitutionAccountsSearch) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.

### HasSearchFor

`func (o *InstitutionAccountsSearch) HasSearchFor() bool`

HasSearchFor returns a boolean if a field has been set.

### GetIsActive

`func (o *InstitutionAccountsSearch) GetIsActive() int64`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *InstitutionAccountsSearch) GetIsActiveOk() (*int64, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *InstitutionAccountsSearch) SetIsActive(v int64)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *InstitutionAccountsSearch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPage

`func (o *InstitutionAccountsSearch) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InstitutionAccountsSearch) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InstitutionAccountsSearch) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *InstitutionAccountsSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *InstitutionAccountsSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InstitutionAccountsSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InstitutionAccountsSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *InstitutionAccountsSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetLimit

`func (o *InstitutionAccountsSearch) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InstitutionAccountsSearch) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InstitutionAccountsSearch) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InstitutionAccountsSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *InstitutionAccountsSearch) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InstitutionAccountsSearch) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InstitutionAccountsSearch) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *InstitutionAccountsSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetInstitutionUserId

`func (o *InstitutionAccountsSearch) GetInstitutionUserId() string`

GetInstitutionUserId returns the InstitutionUserId field if non-nil, zero value otherwise.

### GetInstitutionUserIdOk

`func (o *InstitutionAccountsSearch) GetInstitutionUserIdOk() (*string, bool)`

GetInstitutionUserIdOk returns a tuple with the InstitutionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionUserId

`func (o *InstitutionAccountsSearch) SetInstitutionUserId(v string)`

SetInstitutionUserId sets InstitutionUserId field to given value.

### HasInstitutionUserId

`func (o *InstitutionAccountsSearch) HasInstitutionUserId() bool`

HasInstitutionUserId returns a boolean if a field has been set.

### GetEmail

`func (o *InstitutionAccountsSearch) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InstitutionAccountsSearch) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InstitutionAccountsSearch) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InstitutionAccountsSearch) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


