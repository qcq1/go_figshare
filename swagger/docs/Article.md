# Article

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **string** | Date when article was created | 
**DefinedType** | **int64** | Type of article identifier | 
**DefinedTypeName** | **string** | Name of the article type identifier | 
**Doi** | **string** | DOI | 
**Handle** | **string** | Handle | 
**Id** | **int64** | Unique identifier for article | 
**ResourceDoi** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [default to ""]
**ResourceTitle** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [default to ""]
**Thumb** | **string** | Thumbnail image | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Title** | **string** | Title of article | 
**Url** | **string** | Api endpoint for article | 
**UrlPrivateApi** | **string** | Private Api endpoint for article | 
**UrlPrivateHtml** | **string** | Private site endpoint for article | 
**UrlPublicApi** | **string** | Public Api endpoint for article | 
**UrlPublicHtml** | **string** | Public site endpoint for article | 

## Methods

### NewArticle

`func NewArticle(createdDate string, definedType int64, definedTypeName string, doi string, handle string, id int64, resourceDoi string, resourceTitle string, thumb string, timeline Timeline, title string, url string, urlPrivateApi string, urlPrivateHtml string, urlPublicApi string, urlPublicHtml string, ) *Article`

NewArticle instantiates a new Article object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleWithDefaults

`func NewArticleWithDefaults() *Article`

NewArticleWithDefaults instantiates a new Article object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *Article) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Article) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Article) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDefinedType

`func (o *Article) GetDefinedType() int64`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *Article) GetDefinedTypeOk() (*int64, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *Article) SetDefinedType(v int64)`

SetDefinedType sets DefinedType field to given value.


### GetDefinedTypeName

`func (o *Article) GetDefinedTypeName() string`

GetDefinedTypeName returns the DefinedTypeName field if non-nil, zero value otherwise.

### GetDefinedTypeNameOk

`func (o *Article) GetDefinedTypeNameOk() (*string, bool)`

GetDefinedTypeNameOk returns a tuple with the DefinedTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypeName

`func (o *Article) SetDefinedTypeName(v string)`

SetDefinedTypeName sets DefinedTypeName field to given value.


### GetDoi

`func (o *Article) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *Article) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *Article) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *Article) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Article) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Article) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Article) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Article) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Article) SetId(v int64)`

SetId sets Id field to given value.


### GetResourceDoi

`func (o *Article) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *Article) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *Article) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *Article) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *Article) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *Article) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetThumb

`func (o *Article) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *Article) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *Article) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetTimeline

`func (o *Article) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *Article) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *Article) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetTitle

`func (o *Article) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Article) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Article) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *Article) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Article) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Article) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlPrivateApi

`func (o *Article) GetUrlPrivateApi() string`

GetUrlPrivateApi returns the UrlPrivateApi field if non-nil, zero value otherwise.

### GetUrlPrivateApiOk

`func (o *Article) GetUrlPrivateApiOk() (*string, bool)`

GetUrlPrivateApiOk returns a tuple with the UrlPrivateApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateApi

`func (o *Article) SetUrlPrivateApi(v string)`

SetUrlPrivateApi sets UrlPrivateApi field to given value.


### GetUrlPrivateHtml

`func (o *Article) GetUrlPrivateHtml() string`

GetUrlPrivateHtml returns the UrlPrivateHtml field if non-nil, zero value otherwise.

### GetUrlPrivateHtmlOk

`func (o *Article) GetUrlPrivateHtmlOk() (*string, bool)`

GetUrlPrivateHtmlOk returns a tuple with the UrlPrivateHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateHtml

`func (o *Article) SetUrlPrivateHtml(v string)`

SetUrlPrivateHtml sets UrlPrivateHtml field to given value.


### GetUrlPublicApi

`func (o *Article) GetUrlPublicApi() string`

GetUrlPublicApi returns the UrlPublicApi field if non-nil, zero value otherwise.

### GetUrlPublicApiOk

`func (o *Article) GetUrlPublicApiOk() (*string, bool)`

GetUrlPublicApiOk returns a tuple with the UrlPublicApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicApi

`func (o *Article) SetUrlPublicApi(v string)`

SetUrlPublicApi sets UrlPublicApi field to given value.


### GetUrlPublicHtml

`func (o *Article) GetUrlPublicHtml() string`

GetUrlPublicHtml returns the UrlPublicHtml field if non-nil, zero value otherwise.

### GetUrlPublicHtmlOk

`func (o *Article) GetUrlPublicHtmlOk() (*string, bool)`

GetUrlPublicHtmlOk returns a tuple with the UrlPublicHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicHtml

`func (o *Article) SetUrlPublicHtml(v string)`

SetUrlPublicHtml sets UrlPublicHtml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


