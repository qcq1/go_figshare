# ArticleSearch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int32** | only return collections from this group | [optional] [default to null]
**Institution** | **int32** | only return collections from this institution | [optional] [default to null]
**Limit** | **int64** | Number of results included on a page. Used for pagination with query | [optional] [default to null]
**ModifiedSince** | **string** | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] [default to null]
**Offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | [optional] [default to null]
**OrderDirection** | **string** | Direction of ordering | [optional] [default to null]
**Page** | **int64** | Page number. Used for pagination with page_size | [optional] [default to null]
**PageSize** | **int64** | The number of results included on a page. Used for pagination with page | [optional] [default to 10]
**PublishedSince** | **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | [optional] [default to null]
**SearchFor** | **string** | Search term | [optional] [default to null]
**Doi** | **string** | Only return articles with this doi | [optional] [default to null]
**Handle** | **string** | Only return articles with this handle | [optional] [default to null]
**ItemType** | **int64** | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | [optional] [default to null]
**Order** | **string** | The field by which to order | [optional] [default to null]
**ProjectId** | **int64** | Only return articles in this project | [optional] [default to null]
**ResourceDoi** | **string** | Only return articles with this resource_doi | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


