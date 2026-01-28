# FileCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderPath** | Pointer to **string** | Unix-style directory path of the file; only available if the file was uploaded within a folder structure | [optional] 
**Link** | Pointer to **string** | Url for an existing file that will not be uploaded to Figshare | [optional] 
**Md5** | Pointer to **string** | MD5 sum pre-computed on client side. | [optional] 
**Name** | Pointer to **string** | File name including the extension; can be omitted only for linked files. | [optional] 
**Size** | Pointer to **int64** | File size in bytes; can be omitted only for linked files. | [optional] 

## Methods

### NewFileCreator

`func NewFileCreator() *FileCreator`

NewFileCreator instantiates a new FileCreator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCreatorWithDefaults

`func NewFileCreatorWithDefaults() *FileCreator`

NewFileCreatorWithDefaults instantiates a new FileCreator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderPath

`func (o *FileCreator) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *FileCreator) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *FileCreator) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *FileCreator) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### GetLink

`func (o *FileCreator) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *FileCreator) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *FileCreator) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *FileCreator) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMd5

`func (o *FileCreator) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileCreator) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileCreator) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *FileCreator) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetName

`func (o *FileCreator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileCreator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileCreator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileCreator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *FileCreator) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileCreator) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileCreator) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileCreator) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


