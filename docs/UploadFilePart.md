# UploadFilePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartNo** | Pointer to **int64** | File part id | [optional] 
**StartOffset** | Pointer to **int64** | Indexes on byte range. zero-based and inclusive | [optional] 
**EndOffset** | Pointer to **int64** | Indexes on byte range. zero-based and inclusive | [optional] 
**Status** | Pointer to **string** | part status | [optional] 
**Locked** | Pointer to **bool** | When a part is being uploaded it is being locked, by setting the locked flag to true. No changes/uploads can happen on this part from other requests. | [optional] 

## Methods

### NewUploadFilePart

`func NewUploadFilePart() *UploadFilePart`

NewUploadFilePart instantiates a new UploadFilePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFilePartWithDefaults

`func NewUploadFilePartWithDefaults() *UploadFilePart`

NewUploadFilePartWithDefaults instantiates a new UploadFilePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartNo

`func (o *UploadFilePart) GetPartNo() int64`

GetPartNo returns the PartNo field if non-nil, zero value otherwise.

### GetPartNoOk

`func (o *UploadFilePart) GetPartNoOk() (*int64, bool)`

GetPartNoOk returns a tuple with the PartNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNo

`func (o *UploadFilePart) SetPartNo(v int64)`

SetPartNo sets PartNo field to given value.

### HasPartNo

`func (o *UploadFilePart) HasPartNo() bool`

HasPartNo returns a boolean if a field has been set.

### GetStartOffset

`func (o *UploadFilePart) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *UploadFilePart) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *UploadFilePart) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *UploadFilePart) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### GetEndOffset

`func (o *UploadFilePart) GetEndOffset() int64`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *UploadFilePart) GetEndOffsetOk() (*int64, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *UploadFilePart) SetEndOffset(v int64)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *UploadFilePart) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetStatus

`func (o *UploadFilePart) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadFilePart) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadFilePart) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadFilePart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocked

`func (o *UploadFilePart) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UploadFilePart) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UploadFilePart) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UploadFilePart) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


