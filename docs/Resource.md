# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doi** | Pointer to **string** | DOI of resource item | [optional] [default to ""]
**Id** | Pointer to **string** | ID of resource item | [optional] [default to ""]
**Link** | Pointer to **string** | Link of resource item | [optional] [default to ""]
**Status** | Pointer to **string** | Status of resource item | [optional] [default to ""]
**Title** | Pointer to **string** | Title of resource item | [optional] [default to ""]
**Version** | Pointer to **int64** | Version of resource item | [optional] [default to 0]

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoi

`func (o *Resource) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *Resource) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *Resource) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *Resource) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *Resource) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Resource) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Resource) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Resource) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetStatus

`func (o *Resource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Resource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Resource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Resource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *Resource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Resource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Resource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Resource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *Resource) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Resource) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Resource) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Resource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


