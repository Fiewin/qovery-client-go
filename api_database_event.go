/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DatabaseEventApiService DatabaseEventApi service
type DatabaseEventApiService service

type ApiListDatabaseEventRequest struct {
	ctx        _context.Context
	ApiService *DatabaseEventApiService
	databaseId string
	startId    *string
}

func (r ApiListDatabaseEventRequest) StartId(startId string) ApiListDatabaseEventRequest {
	r.startId = &startId
	return r
}

func (r ApiListDatabaseEventRequest) Execute() (EventPaginatedResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListDatabaseEventExecute(r)
}

/*
 * ListDatabaseEvent List database  events
 * By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param databaseId Database ID
 * @return ApiListDatabaseEventRequest
 */
func (a *DatabaseEventApiService) ListDatabaseEvent(ctx _context.Context, databaseId string) ApiListDatabaseEventRequest {
	return ApiListDatabaseEventRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

/*
 * Execute executes the request
 * @return EventPaginatedResponseList
 */
func (a *DatabaseEventApiService) ListDatabaseEventExecute(r ApiListDatabaseEventRequest) (EventPaginatedResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EventPaginatedResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseEventApiService.ListDatabaseEvent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/event"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", _neturl.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startId != nil {
		localVarQueryParams.Add("startId", parameterToString(*r.startId, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
