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

// CustomDomainApiService CustomDomainApi service
type CustomDomainApiService service

type ApiCreateApplicationCustomDomainRequest struct {
	ctx                 _context.Context
	ApiService          *CustomDomainApiService
	applicationId       string
	customDomainRequest *CustomDomainRequest
}

func (r ApiCreateApplicationCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiCreateApplicationCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiCreateApplicationCustomDomainRequest) Execute() (CustomDomainResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateApplicationCustomDomainExecute(r)
}

/*
 * CreateApplicationCustomDomain Add custom domain to the application.
 * Add a custom domain to this application in order not to use qovery autogenerated domain
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param applicationId Application ID
 * @return ApiCreateApplicationCustomDomainRequest
 */
func (a *CustomDomainApiService) CreateApplicationCustomDomain(ctx _context.Context, applicationId string) ApiCreateApplicationCustomDomainRequest {
	return ApiCreateApplicationCustomDomainRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

/*
 * Execute executes the request
 * @return CustomDomainResponse
 */
func (a *CustomDomainApiService) CreateApplicationCustomDomainExecute(r ApiCreateApplicationCustomDomainRequest) (CustomDomainResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CustomDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.CreateApplicationCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.customDomainRequest
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

type ApiDeleteCustomDomainRequest struct {
	ctx            _context.Context
	ApiService     *CustomDomainApiService
	applicationId  string
	customDomainId string
}

func (r ApiDeleteCustomDomainRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCustomDomainExecute(r)
}

/*
 * DeleteCustomDomain Delete a Custom Domain
 * To delete an CustomDomain you must have the project user permission
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param applicationId Application ID
 * @param customDomainId Custom Domain ID
 * @return ApiDeleteCustomDomainRequest
 */
func (a *CustomDomainApiService) DeleteCustomDomain(ctx _context.Context, applicationId string, customDomainId string) ApiDeleteCustomDomainRequest {
	return ApiDeleteCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		applicationId:  applicationId,
		customDomainId: customDomainId,
	}
}

/*
 * Execute executes the request
 */
func (a *CustomDomainApiService) DeleteCustomDomainExecute(r ApiDeleteCustomDomainRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.DeleteCustomDomain")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", _neturl.PathEscape(parameterToString(r.customDomainId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditCustomDomainRequest struct {
	ctx                 _context.Context
	ApiService          *CustomDomainApiService
	applicationId       string
	customDomainId      string
	customDomainRequest *CustomDomainRequest
}

func (r ApiEditCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiEditCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiEditCustomDomainRequest) Execute() (CustomDomainResponse, *_nethttp.Response, error) {
	return r.ApiService.EditCustomDomainExecute(r)
}

/*
 * EditCustomDomain Edit a Custom Domain
 * To edit a Custom Domain  you must have the project user permission
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param applicationId Application ID
 * @param customDomainId Custom Domain ID
 * @return ApiEditCustomDomainRequest
 */
func (a *CustomDomainApiService) EditCustomDomain(ctx _context.Context, applicationId string, customDomainId string) ApiEditCustomDomainRequest {
	return ApiEditCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		applicationId:  applicationId,
		customDomainId: customDomainId,
	}
}

/*
 * Execute executes the request
 * @return CustomDomainResponse
 */
func (a *CustomDomainApiService) EditCustomDomainExecute(r ApiEditCustomDomainRequest) (CustomDomainResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CustomDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.EditCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", _neturl.PathEscape(parameterToString(r.customDomainId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.customDomainRequest
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

type ApiListApplicationCustomDomainRequest struct {
	ctx           _context.Context
	ApiService    *CustomDomainApiService
	applicationId string
}

func (r ApiListApplicationCustomDomainRequest) Execute() (CustomDomainResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListApplicationCustomDomainExecute(r)
}

/*
 * ListApplicationCustomDomain List application custom domains
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param applicationId Application ID
 * @return ApiListApplicationCustomDomainRequest
 */
func (a *CustomDomainApiService) ListApplicationCustomDomain(ctx _context.Context, applicationId string) ApiListApplicationCustomDomainRequest {
	return ApiListApplicationCustomDomainRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

/*
 * Execute executes the request
 * @return CustomDomainResponseList
 */
func (a *CustomDomainApiService) ListApplicationCustomDomainExecute(r ApiListApplicationCustomDomainRequest) (CustomDomainResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CustomDomainResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.ListApplicationCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
