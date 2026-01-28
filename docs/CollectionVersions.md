# CollectionVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funding** | [**[]FundingInformation**](FundingInformation.md) | Full Collection funding information | 
**Url** | **string** | Api endpoint for the collection version | 
**Version** | **int64** | Version number | 

## Methods

### NewCollectionVersions

`func NewCollectionVersions(funding []FundingInformation, url string, version int64, ) *CollectionVersions`

NewCollectionVersions instantiates a new CollectionVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionVersionsWithDefaults

`func NewCollectionVersionsWithDefaults() *CollectionVersions`

NewCollectionVersionsWithDefaults instantiates a new CollectionVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunding

`func (o *CollectionVersions) GetFunding() []FundingInformation`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CollectionVersions) GetFundingOk() (*[]FundingInformation, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CollectionVersions) SetFunding(v []FundingInformation)`

SetFunding sets Funding field to given value.


### GetUrl

`func (o *CollectionVersions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionVersions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionVersions) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *CollectionVersions) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionVersions) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionVersions) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


