# UploadInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | token received after initializing a file upload | [optional] 
**Md5** | Pointer to **string** | md5 provided on upload initialization | [optional] 
**Size** | Pointer to **int64** | size of file in bytes | [optional] 
**Name** | Pointer to **string** | name of file on upload server | [optional] 
**Status** | Pointer to **string** | Upload status | [optional] 
**Parts** | Pointer to [**[]UploadFilePart**](UploadFilePart.md) | Uploads parts | [optional] 

## Methods

### NewUploadInfo

`func NewUploadInfo() *UploadInfo`

NewUploadInfo instantiates a new UploadInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadInfoWithDefaults

`func NewUploadInfoWithDefaults() *UploadInfo`

NewUploadInfoWithDefaults instantiates a new UploadInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UploadInfo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UploadInfo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UploadInfo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UploadInfo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetMd5

`func (o *UploadInfo) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *UploadInfo) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *UploadInfo) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *UploadInfo) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSize

`func (o *UploadInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UploadInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetName

`func (o *UploadInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UploadInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParts

`func (o *UploadInfo) GetParts() []UploadFilePart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *UploadInfo) GetPartsOk() (*[]UploadFilePart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *UploadInfo) SetParts(v []UploadFilePart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *UploadInfo) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


