# AccountReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | A unique ID for the AccountRecord | 
**AccountId** | **int64** | The ID of the account which generated this report. | 
**CreatedDate** | **string** | Date when the AccountReport was requested | 
**Status** | **string** | Status of the report | 
**DownloadUrl** | **string** | The download link for the generated XLSX | 
**GroupId** | **int64** | The group ID that was used to filter the report, if any. | 

## Methods

### NewAccountReport

`func NewAccountReport(id int64, accountId int64, createdDate string, status string, downloadUrl string, groupId int64, ) *AccountReport`

NewAccountReport instantiates a new AccountReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountReportWithDefaults

`func NewAccountReportWithDefaults() *AccountReport`

NewAccountReportWithDefaults instantiates a new AccountReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountReport) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountReport) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountReport) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *AccountReport) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountReport) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountReport) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetCreatedDate

`func (o *AccountReport) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AccountReport) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AccountReport) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetStatus

`func (o *AccountReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountReport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDownloadUrl

`func (o *AccountReport) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *AccountReport) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *AccountReport) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetGroupId

`func (o *AccountReport) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AccountReport) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AccountReport) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


