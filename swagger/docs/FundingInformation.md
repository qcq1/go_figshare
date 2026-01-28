# FundingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunderName** | **string** | Funder&#39;s name | 
**GrantCode** | **string** | The grant code | 
**Id** | **int64** | Funding id | 
**IsUserDefined** | **int64** | Return 1 whether the grant has been introduced manually, 0 otherwise | 
**Title** | **string** | The funding name | 
**Url** | **string** | The grant url | 

## Methods

### NewFundingInformation

`func NewFundingInformation(funderName string, grantCode string, id int64, isUserDefined int64, title string, url string, ) *FundingInformation`

NewFundingInformation instantiates a new FundingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingInformationWithDefaults

`func NewFundingInformationWithDefaults() *FundingInformation`

NewFundingInformationWithDefaults instantiates a new FundingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunderName

`func (o *FundingInformation) GetFunderName() string`

GetFunderName returns the FunderName field if non-nil, zero value otherwise.

### GetFunderNameOk

`func (o *FundingInformation) GetFunderNameOk() (*string, bool)`

GetFunderNameOk returns a tuple with the FunderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunderName

`func (o *FundingInformation) SetFunderName(v string)`

SetFunderName sets FunderName field to given value.


### GetGrantCode

`func (o *FundingInformation) GetGrantCode() string`

GetGrantCode returns the GrantCode field if non-nil, zero value otherwise.

### GetGrantCodeOk

`func (o *FundingInformation) GetGrantCodeOk() (*string, bool)`

GetGrantCodeOk returns a tuple with the GrantCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantCode

`func (o *FundingInformation) SetGrantCode(v string)`

SetGrantCode sets GrantCode field to given value.


### GetId

`func (o *FundingInformation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FundingInformation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FundingInformation) SetId(v int64)`

SetId sets Id field to given value.


### GetIsUserDefined

`func (o *FundingInformation) GetIsUserDefined() int64`

GetIsUserDefined returns the IsUserDefined field if non-nil, zero value otherwise.

### GetIsUserDefinedOk

`func (o *FundingInformation) GetIsUserDefinedOk() (*int64, bool)`

GetIsUserDefinedOk returns a tuple with the IsUserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserDefined

`func (o *FundingInformation) SetIsUserDefined(v int64)`

SetIsUserDefined sets IsUserDefined field to given value.


### GetTitle

`func (o *FundingInformation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FundingInformation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FundingInformation) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *FundingInformation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FundingInformation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FundingInformation) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


