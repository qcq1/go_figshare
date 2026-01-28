# PublicFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedMd5** | **string** | File computed md5 | 
**DownloadUrl** | **string** | Url for file download | 
**Id** | **int64** | File id | 
**IsLinkOnly** | **bool** | True if file is hosted somewhere else | 
**Mimetype** | Pointer to **string** | MIME Type of the file, it defaults to an empty string | [optional] 
**Name** | **string** | File name | 
**Size** | **int64** | File size | 
**SuppliedMd5** | **string** | File supplied md5 | 

## Methods

### NewPublicFile

`func NewPublicFile(computedMd5 string, downloadUrl string, id int64, isLinkOnly bool, name string, size int64, suppliedMd5 string, ) *PublicFile`

NewPublicFile instantiates a new PublicFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFileWithDefaults

`func NewPublicFileWithDefaults() *PublicFile`

NewPublicFileWithDefaults instantiates a new PublicFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedMd5

`func (o *PublicFile) GetComputedMd5() string`

GetComputedMd5 returns the ComputedMd5 field if non-nil, zero value otherwise.

### GetComputedMd5Ok

`func (o *PublicFile) GetComputedMd5Ok() (*string, bool)`

GetComputedMd5Ok returns a tuple with the ComputedMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedMd5

`func (o *PublicFile) SetComputedMd5(v string)`

SetComputedMd5 sets ComputedMd5 field to given value.


### GetDownloadUrl

`func (o *PublicFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PublicFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PublicFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetId

`func (o *PublicFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicFile) SetId(v int64)`

SetId sets Id field to given value.


### GetIsLinkOnly

`func (o *PublicFile) GetIsLinkOnly() bool`

GetIsLinkOnly returns the IsLinkOnly field if non-nil, zero value otherwise.

### GetIsLinkOnlyOk

`func (o *PublicFile) GetIsLinkOnlyOk() (*bool, bool)`

GetIsLinkOnlyOk returns a tuple with the IsLinkOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkOnly

`func (o *PublicFile) SetIsLinkOnly(v bool)`

SetIsLinkOnly sets IsLinkOnly field to given value.


### GetMimetype

`func (o *PublicFile) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *PublicFile) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *PublicFile) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *PublicFile) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetName

`func (o *PublicFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicFile) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *PublicFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PublicFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PublicFile) SetSize(v int64)`

SetSize sets Size field to given value.


### GetSuppliedMd5

`func (o *PublicFile) GetSuppliedMd5() string`

GetSuppliedMd5 returns the SuppliedMd5 field if non-nil, zero value otherwise.

### GetSuppliedMd5Ok

`func (o *PublicFile) GetSuppliedMd5Ok() (*string, bool)`

GetSuppliedMd5Ok returns a tuple with the SuppliedMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedMd5

`func (o *PublicFile) SetSuppliedMd5(v string)`

SetSuppliedMd5 sets SuppliedMd5 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


