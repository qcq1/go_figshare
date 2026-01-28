# ShortAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Account id | 
**FirstName** | **string** | First Name | 
**LastName** | **string** | Last Name | 
**InstitutionId** | **int64** | Account institution | 
**Email** | **string** | User email | 
**Active** | **int64** | Account activity status | 
**InstitutionUserId** | **string** | Account institution user id | 
**Quota** | **int64** | Total storage available to account, in bytes | 
**UsedQuota** | **int64** | Storage used by the account, in bytes | 
**UserId** | **int64** | User id associated with account, useful for example for adding the account as an author to an item | 
**OrcidId** | **string** | ORCID iD associated to account | 
**SymplecticUserId** | **string** | Symplectic ID associated to account | 

## Methods

### NewShortAccount

`func NewShortAccount(id int64, firstName string, lastName string, institutionId int64, email string, active int64, institutionUserId string, quota int64, usedQuota int64, userId int64, orcidId string, symplecticUserId string, ) *ShortAccount`

NewShortAccount instantiates a new ShortAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortAccountWithDefaults

`func NewShortAccountWithDefaults() *ShortAccount`

NewShortAccountWithDefaults instantiates a new ShortAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShortAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortAccount) SetId(v int64)`

SetId sets Id field to given value.


### GetFirstName

`func (o *ShortAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ShortAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ShortAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ShortAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ShortAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ShortAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetInstitutionId

`func (o *ShortAccount) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *ShortAccount) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *ShortAccount) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.


### GetEmail

`func (o *ShortAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ShortAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ShortAccount) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetActive

`func (o *ShortAccount) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ShortAccount) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ShortAccount) SetActive(v int64)`

SetActive sets Active field to given value.


### GetInstitutionUserId

`func (o *ShortAccount) GetInstitutionUserId() string`

GetInstitutionUserId returns the InstitutionUserId field if non-nil, zero value otherwise.

### GetInstitutionUserIdOk

`func (o *ShortAccount) GetInstitutionUserIdOk() (*string, bool)`

GetInstitutionUserIdOk returns a tuple with the InstitutionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionUserId

`func (o *ShortAccount) SetInstitutionUserId(v string)`

SetInstitutionUserId sets InstitutionUserId field to given value.


### GetQuota

`func (o *ShortAccount) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ShortAccount) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ShortAccount) SetQuota(v int64)`

SetQuota sets Quota field to given value.


### GetUsedQuota

`func (o *ShortAccount) GetUsedQuota() int64`

GetUsedQuota returns the UsedQuota field if non-nil, zero value otherwise.

### GetUsedQuotaOk

`func (o *ShortAccount) GetUsedQuotaOk() (*int64, bool)`

GetUsedQuotaOk returns a tuple with the UsedQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuota

`func (o *ShortAccount) SetUsedQuota(v int64)`

SetUsedQuota sets UsedQuota field to given value.


### GetUserId

`func (o *ShortAccount) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShortAccount) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShortAccount) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetOrcidId

`func (o *ShortAccount) GetOrcidId() string`

GetOrcidId returns the OrcidId field if non-nil, zero value otherwise.

### GetOrcidIdOk

`func (o *ShortAccount) GetOrcidIdOk() (*string, bool)`

GetOrcidIdOk returns a tuple with the OrcidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcidId

`func (o *ShortAccount) SetOrcidId(v string)`

SetOrcidId sets OrcidId field to given value.


### GetSymplecticUserId

`func (o *ShortAccount) GetSymplecticUserId() string`

GetSymplecticUserId returns the SymplecticUserId field if non-nil, zero value otherwise.

### GetSymplecticUserIdOk

`func (o *ShortAccount) GetSymplecticUserIdOk() (*string, bool)`

GetSymplecticUserIdOk returns a tuple with the SymplecticUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymplecticUserId

`func (o *ShortAccount) SetSymplecticUserId(v string)`

SetSymplecticUserId sets SymplecticUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


