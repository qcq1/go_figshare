# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Account id | 
**FirstName** | **string** | First Name | 
**LastName** | **string** | Last Name | 
**UsedQuotaPrivate** | **int64** | Account used private quota | 
**ModifiedDate** | **string** | Date of last account modification | 
**UsedQuota** | **int64** | Account total used quota | 
**CreatedDate** | **string** | Date when account was created | 
**Quota** | **int64** | Account quota | 
**GroupId** | **int64** | Account group id | 
**InstitutionUserId** | **string** | Account institution user id | 
**InstitutionId** | **int64** | Account institution | 
**Email** | **string** | User email | 
**UsedQuotaPublic** | **int64** | Account public used quota | 
**PendingQuotaRequest** | **bool** | True if a quota request is pending | 
**Active** | **int64** | Account activity status | 
**MaximumFileSize** | **int64** | Maximum upload size for account | 
**UserId** | **int64** | User id associated with account, useful for example for adding the account as an author to an item | 
**OrcidId** | **string** | ORCID iD associated to account | 
**SymplecticUserId** | **string** | Symplectic ID associated to account | 

## Methods

### NewAccount

`func NewAccount(id int64, firstName string, lastName string, usedQuotaPrivate int64, modifiedDate string, usedQuota int64, createdDate string, quota int64, groupId int64, institutionUserId string, institutionId int64, email string, usedQuotaPublic int64, pendingQuotaRequest bool, active int64, maximumFileSize int64, userId int64, orcidId string, symplecticUserId string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int64)`

SetId sets Id field to given value.


### GetFirstName

`func (o *Account) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Account) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Account) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Account) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Account) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Account) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUsedQuotaPrivate

`func (o *Account) GetUsedQuotaPrivate() int64`

GetUsedQuotaPrivate returns the UsedQuotaPrivate field if non-nil, zero value otherwise.

### GetUsedQuotaPrivateOk

`func (o *Account) GetUsedQuotaPrivateOk() (*int64, bool)`

GetUsedQuotaPrivateOk returns a tuple with the UsedQuotaPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaPrivate

`func (o *Account) SetUsedQuotaPrivate(v int64)`

SetUsedQuotaPrivate sets UsedQuotaPrivate field to given value.


### GetModifiedDate

`func (o *Account) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *Account) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *Account) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetUsedQuota

`func (o *Account) GetUsedQuota() int64`

GetUsedQuota returns the UsedQuota field if non-nil, zero value otherwise.

### GetUsedQuotaOk

`func (o *Account) GetUsedQuotaOk() (*int64, bool)`

GetUsedQuotaOk returns a tuple with the UsedQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuota

`func (o *Account) SetUsedQuota(v int64)`

SetUsedQuota sets UsedQuota field to given value.


### GetCreatedDate

`func (o *Account) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Account) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Account) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetQuota

`func (o *Account) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Account) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Account) SetQuota(v int64)`

SetQuota sets Quota field to given value.


### GetGroupId

`func (o *Account) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Account) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Account) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetInstitutionUserId

`func (o *Account) GetInstitutionUserId() string`

GetInstitutionUserId returns the InstitutionUserId field if non-nil, zero value otherwise.

### GetInstitutionUserIdOk

`func (o *Account) GetInstitutionUserIdOk() (*string, bool)`

GetInstitutionUserIdOk returns a tuple with the InstitutionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionUserId

`func (o *Account) SetInstitutionUserId(v string)`

SetInstitutionUserId sets InstitutionUserId field to given value.


### GetInstitutionId

`func (o *Account) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *Account) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *Account) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.


### GetEmail

`func (o *Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Account) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsedQuotaPublic

`func (o *Account) GetUsedQuotaPublic() int64`

GetUsedQuotaPublic returns the UsedQuotaPublic field if non-nil, zero value otherwise.

### GetUsedQuotaPublicOk

`func (o *Account) GetUsedQuotaPublicOk() (*int64, bool)`

GetUsedQuotaPublicOk returns a tuple with the UsedQuotaPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaPublic

`func (o *Account) SetUsedQuotaPublic(v int64)`

SetUsedQuotaPublic sets UsedQuotaPublic field to given value.


### GetPendingQuotaRequest

`func (o *Account) GetPendingQuotaRequest() bool`

GetPendingQuotaRequest returns the PendingQuotaRequest field if non-nil, zero value otherwise.

### GetPendingQuotaRequestOk

`func (o *Account) GetPendingQuotaRequestOk() (*bool, bool)`

GetPendingQuotaRequestOk returns a tuple with the PendingQuotaRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingQuotaRequest

`func (o *Account) SetPendingQuotaRequest(v bool)`

SetPendingQuotaRequest sets PendingQuotaRequest field to given value.


### GetActive

`func (o *Account) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Account) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Account) SetActive(v int64)`

SetActive sets Active field to given value.


### GetMaximumFileSize

`func (o *Account) GetMaximumFileSize() int64`

GetMaximumFileSize returns the MaximumFileSize field if non-nil, zero value otherwise.

### GetMaximumFileSizeOk

`func (o *Account) GetMaximumFileSizeOk() (*int64, bool)`

GetMaximumFileSizeOk returns a tuple with the MaximumFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileSize

`func (o *Account) SetMaximumFileSize(v int64)`

SetMaximumFileSize sets MaximumFileSize field to given value.


### GetUserId

`func (o *Account) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Account) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Account) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetOrcidId

`func (o *Account) GetOrcidId() string`

GetOrcidId returns the OrcidId field if non-nil, zero value otherwise.

### GetOrcidIdOk

`func (o *Account) GetOrcidIdOk() (*string, bool)`

GetOrcidIdOk returns a tuple with the OrcidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcidId

`func (o *Account) SetOrcidId(v string)`

SetOrcidId sets OrcidId field to given value.


### GetSymplecticUserId

`func (o *Account) GetSymplecticUserId() string`

GetSymplecticUserId returns the SymplecticUserId field if non-nil, zero value otherwise.

### GetSymplecticUserIdOk

`func (o *Account) GetSymplecticUserIdOk() (*string, bool)`

GetSymplecticUserIdOk returns a tuple with the SymplecticUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymplecticUserId

`func (o *Account) SetSymplecticUserId(v string)`

SetSymplecticUserId sets SymplecticUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


