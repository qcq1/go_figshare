# GroupEmbargoOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Embargo option id | 
**IpName** | **string** | IP range name; value appears if type is ip_range | 
**Type** | **string** | Embargo permission type | 

## Methods

### NewGroupEmbargoOptions

`func NewGroupEmbargoOptions(id int64, ipName string, type_ string, ) *GroupEmbargoOptions`

NewGroupEmbargoOptions instantiates a new GroupEmbargoOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupEmbargoOptionsWithDefaults

`func NewGroupEmbargoOptionsWithDefaults() *GroupEmbargoOptions`

NewGroupEmbargoOptionsWithDefaults instantiates a new GroupEmbargoOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupEmbargoOptions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupEmbargoOptions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupEmbargoOptions) SetId(v int64)`

SetId sets Id field to given value.


### GetIpName

`func (o *GroupEmbargoOptions) GetIpName() string`

GetIpName returns the IpName field if non-nil, zero value otherwise.

### GetIpNameOk

`func (o *GroupEmbargoOptions) GetIpNameOk() (*string, bool)`

GetIpNameOk returns a tuple with the IpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpName

`func (o *GroupEmbargoOptions) SetIpName(v string)`

SetIpName sets IpName field to given value.


### GetType

`func (o *GroupEmbargoOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupEmbargoOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupEmbargoOptions) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


