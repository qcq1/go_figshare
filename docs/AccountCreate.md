# AccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of account | 
**FirstName** | Pointer to **string** | First Name | [optional] [default to ""]
**GroupId** | Pointer to **int64** | Not applicable to regular users. This field is reserved to institutions/publishers with access to assign to specific groups | [optional] 
**InstitutionUserId** | Pointer to **string** | Institution user id | [optional] [default to ""]
**IsActive** | Pointer to **bool** | Is account active | [optional] 
**LastName** | **string** | Last Name | [default to ""]
**Quota** | Pointer to **int64** | Account quota | [optional] 
**SymplecticUserId** | Pointer to **string** | Symplectic user id | [optional] [default to ""]

## Methods

### NewAccountCreate

`func NewAccountCreate(email string, lastName string, ) *AccountCreate`

NewAccountCreate instantiates a new AccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateWithDefaults

`func NewAccountCreateWithDefaults() *AccountCreate`

NewAccountCreateWithDefaults instantiates a new AccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *AccountCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountCreate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroupId

`func (o *AccountCreate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AccountCreate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AccountCreate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AccountCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetInstitutionUserId

`func (o *AccountCreate) GetInstitutionUserId() string`

GetInstitutionUserId returns the InstitutionUserId field if non-nil, zero value otherwise.

### GetInstitutionUserIdOk

`func (o *AccountCreate) GetInstitutionUserIdOk() (*string, bool)`

GetInstitutionUserIdOk returns a tuple with the InstitutionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionUserId

`func (o *AccountCreate) SetInstitutionUserId(v string)`

SetInstitutionUserId sets InstitutionUserId field to given value.

### HasInstitutionUserId

`func (o *AccountCreate) HasInstitutionUserId() bool`

HasInstitutionUserId returns a boolean if a field has been set.

### GetIsActive

`func (o *AccountCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AccountCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastName

`func (o *AccountCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetQuota

`func (o *AccountCreate) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *AccountCreate) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *AccountCreate) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *AccountCreate) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetSymplecticUserId

`func (o *AccountCreate) GetSymplecticUserId() string`

GetSymplecticUserId returns the SymplecticUserId field if non-nil, zero value otherwise.

### GetSymplecticUserIdOk

`func (o *AccountCreate) GetSymplecticUserIdOk() (*string, bool)`

GetSymplecticUserIdOk returns a tuple with the SymplecticUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymplecticUserId

`func (o *AccountCreate) SetSymplecticUserId(v string)`

SetSymplecticUserId sets SymplecticUserId field to given value.

### HasSymplecticUserId

`func (o *AccountCreate) HasSymplecticUserId() bool`

HasSymplecticUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


