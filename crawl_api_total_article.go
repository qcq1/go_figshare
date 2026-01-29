package go_figshare

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (a *ArticlesAPIService) TotalArticle(ctx context.Context, articleId int64) ApiTotalArticleRequest {
	return ApiTotalArticleRequest{
		ApiService: a,
		ctx:        ctx,
		articleId:  articleId,
	}
}

type ApiTotalArticleRequest struct {
	ctx        context.Context
	ApiService *ArticlesAPIService
	articleId  int64
}

func (r ApiTotalArticleRequest) Execute() (*TotalArticle, *http.Response, error) {
	return r.ApiService.TotalArticleExecute(r)
}

func (a *ArticlesAPIService) TotalArticleExecute(r ApiTotalArticleRequest) (*TotalArticle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TotalArticle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArticlesAPIService.TotalArticle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/total/articles/{article_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"article_id"+"}", url.PathEscape(parameterValueToString(r.articleId, "articleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.articleId < 1 {
		return localVarReturnValue, nil, reportError("articleId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
