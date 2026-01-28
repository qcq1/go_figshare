# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doi** | **string** | Collection DOI | 
**Handle** | **string** | Collection Handle | 
**Id** | **int64** | Collection id | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Title** | **string** | Collection title | 
**Url** | **string** | Api endpoint | 

## Methods

### NewCollection

`func NewCollection(doi string, handle string, id int64, timeline Timeline, title string, url string, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoi

`func (o *Collection) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *Collection) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *Collection) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *Collection) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Collection) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Collection) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Collection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collection) SetId(v int64)`

SetId sets Id field to given value.


### GetTimeline

`func (o *Collection) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *Collection) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *Collection) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetTitle

`func (o *Collection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Collection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Collection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *Collection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Collection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Collection) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


