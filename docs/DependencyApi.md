# \DependencyApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationDependency**](DependencyApi.md#CreateApplicationDependency) | **Post** /application/{applicationId}/dependency/{targetApplicationId} | Add application dependency to this application.
[**ListApplicationDependency**](DependencyApi.md#ListApplicationDependency) | **Get** /application/{applicationId}/dependency | List application dependencies
[**RemoveApplicationDependency**](DependencyApi.md#RemoveApplicationDependency) | **Delete** /application/{applicationId}/dependency/{targetApplicationId} | Remove application dependency to this application.



## CreateApplicationDependency

> ApplicationResponse CreateApplicationDependency(ctx, applicationId, targetApplicationId).Execute()

Add application dependency to this application.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | Application ID
    targetApplicationId := TODO // string | Target application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DependencyApi.CreateApplicationDependency(context.Background(), applicationId, targetApplicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependencyApi.CreateApplicationDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationDependency`: ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `DependencyApi.CreateApplicationDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**targetApplicationId** | [**string**](.md) | Target application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationDependency

> ApplicationResponseList ListApplicationDependency(ctx, applicationId).Execute()

List application dependencies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | Application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DependencyApi.ListApplicationDependency(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependencyApi.ListApplicationDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationDependency`: ApplicationResponseList
    fmt.Fprintf(os.Stdout, "Response from `DependencyApi.ListApplicationDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponseList**](ApplicationResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveApplicationDependency

> RemoveApplicationDependency(ctx, applicationId, targetApplicationId).Execute()

Remove application dependency to this application.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | Application ID
    targetApplicationId := TODO // string | Target application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DependencyApi.RemoveApplicationDependency(context.Background(), applicationId, targetApplicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependencyApi.RemoveApplicationDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**targetApplicationId** | [**string**](.md) | Target application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApplicationDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

