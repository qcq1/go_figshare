# PrivateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAttachedToPublicVersion** | **bool** | True if the file is attached to a public item version | 
**PreviewState** | **string** | File preview state | 
**UploadToken** | **string** | Token for file upload | 
**UploadUrl** | **string** | Upload url for file | 
**ViewerType** | **string** | File viewer type | 
**ComputedMd5** | **string** | File computed md5 | 
**DownloadUrl** | **string** | Url for file download | 
**Id** | **int64** | File id | 
**IsLinkOnly** | **bool** | True if file is hosted somewhere else | 
**Mimetype** | Pointer to **string** | MIME Type of the file, it defaults to an empty string | [optional] 
**Name** | **string** | File name | 
**Size** | **int64** | File size | 
**SuppliedMd5** | **string** | File supplied md5 | 

## Methods

### NewPrivateFile

`func NewPrivateFile(isAttachedToPublicVersion bool, previewState string, uploadToken string, uploadUrl string, viewerType string, computedMd5 string, downloadUrl string, id int64, isLinkOnly bool, name string, size int64, suppliedMd5 string, ) *PrivateFile`

NewPrivateFile instantiates a new PrivateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFileWithDefaults

`func NewPrivateFileWithDefaults() *PrivateFile`

NewPrivateFileWithDefaults instantiates a new PrivateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAttachedToPublicVersion

`func (o *PrivateFile) GetIsAttachedToPublicVersion() bool`

GetIsAttachedToPublicVersion returns the IsAttachedToPublicVersion field if non-nil, zero value otherwise.

### GetIsAttachedToPublicVersionOk

`func (o *PrivateFile) GetIsAttachedToPublicVersionOk() (*bool, bool)`

GetIsAttachedToPublicVersionOk returns a tuple with the IsAttachedToPublicVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttachedToPublicVersion

`func (o *PrivateFile) SetIsAttachedToPublicVersion(v bool)`

SetIsAttachedToPublicVersion sets IsAttachedToPublicVersion field to given value.


### GetPreviewState

`func (o *PrivateFile) GetPreviewState() string`

GetPreviewState returns the PreviewState field if non-nil, zero value otherwise.

### GetPreviewStateOk

`func (o *PrivateFile) GetPreviewStateOk() (*string, bool)`

GetPreviewStateOk returns a tuple with the PreviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewState

`func (o *PrivateFile) SetPreviewState(v string)`

SetPreviewState sets PreviewState field to given value.


### GetUploadToken

`func (o *PrivateFile) GetUploadToken() string`

GetUploadToken returns the UploadToken field if non-nil, zero value otherwise.

### GetUploadTokenOk

`func (o *PrivateFile) GetUploadTokenOk() (*string, bool)`

GetUploadTokenOk returns a tuple with the UploadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadToken

`func (o *PrivateFile) SetUploadToken(v string)`

SetUploadToken sets UploadToken field to given value.


### GetUploadUrl

`func (o *PrivateFile) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *PrivateFile) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *PrivateFile) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.


### GetViewerType

`func (o *PrivateFile) GetViewerType() string`

GetViewerType returns the ViewerType field if non-nil, zero value otherwise.

### GetViewerTypeOk

`func (o *PrivateFile) GetViewerTypeOk() (*string, bool)`

GetViewerTypeOk returns a tuple with the ViewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerType

`func (o *PrivateFile) SetViewerType(v string)`

SetViewerType sets ViewerType field to given value.


### GetComputedMd5

`func (o *PrivateFile) GetComputedMd5() string`

GetComputedMd5 returns the ComputedMd5 field if non-nil, zero value otherwise.

### GetComputedMd5Ok

`func (o *PrivateFile) GetComputedMd5Ok() (*string, bool)`

GetComputedMd5Ok returns a tuple with the ComputedMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedMd5

`func (o *PrivateFile) SetComputedMd5(v string)`

SetComputedMd5 sets ComputedMd5 field to given value.


### GetDownloadUrl

`func (o *PrivateFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PrivateFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PrivateFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetId

`func (o *PrivateFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateFile) SetId(v int64)`

SetId sets Id field to given value.


### GetIsLinkOnly

`func (o *PrivateFile) GetIsLinkOnly() bool`

GetIsLinkOnly returns the IsLinkOnly field if non-nil, zero value otherwise.

### GetIsLinkOnlyOk

`func (o *PrivateFile) GetIsLinkOnlyOk() (*bool, bool)`

GetIsLinkOnlyOk returns a tuple with the IsLinkOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkOnly

`func (o *PrivateFile) SetIsLinkOnly(v bool)`

SetIsLinkOnly sets IsLinkOnly field to given value.


### GetMimetype

`func (o *PrivateFile) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *PrivateFile) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *PrivateFile) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *PrivateFile) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetName

`func (o *PrivateFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateFile) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *PrivateFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrivateFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrivateFile) SetSize(v int64)`

SetSize sets Size field to given value.


### GetSuppliedMd5

`func (o *PrivateFile) GetSuppliedMd5() string`

GetSuppliedMd5 returns the SuppliedMd5 field if non-nil, zero value otherwise.

### GetSuppliedMd5Ok

`func (o *PrivateFile) GetSuppliedMd5Ok() (*string, bool)`

GetSuppliedMd5Ok returns a tuple with the SuppliedMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedMd5

`func (o *PrivateFile) SetSuppliedMd5(v string)`

SetSuppliedMd5 sets SuppliedMd5 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


