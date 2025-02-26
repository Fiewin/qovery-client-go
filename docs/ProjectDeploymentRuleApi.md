# \ProjectDeploymentRuleApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentRule**](ProjectDeploymentRuleApi.md#CreateDeploymentRule) | **Post** /project/{projectId}/deploymentRule | Create a deployment rule
[**DeleteProjectDeploymentRule**](ProjectDeploymentRuleApi.md#DeleteProjectDeploymentRule) | **Delete** /project/{projectId}/deploymentRule/{deploymentRuleId} | Delete a project deployment rule
[**EditProjectDeployemtnRule**](ProjectDeploymentRuleApi.md#EditProjectDeployemtnRule) | **Put** /project/{projectId}/deploymentRule/{deploymentRuleId} | Edit a project deployment rule
[**GetProjectDeploymentRule**](ProjectDeploymentRuleApi.md#GetProjectDeploymentRule) | **Get** /project/{projectId}/deploymentRule/{deploymentRuleId} | Get project deployment rule
[**ListProjectDeploymentRule**](ProjectDeploymentRuleApi.md#ListProjectDeploymentRule) | **Get** /project/{projectId}/deploymentRule | List project deployment rules



## CreateDeploymentRule

> ProjectDeploymentRuleResponse CreateDeploymentRule(ctx, projectId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()

Create a deployment rule

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
    projectId := TODO // string | Project ID
    projectDeploymentRuleRequest := *openapiclient.NewProjectDeploymentRuleRequest("^feat", "Name_example", "Mode_example", "Cluster_example", false) // ProjectDeploymentRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectDeploymentRuleApi.CreateDeploymentRule(context.Background(), projectId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleApi.CreateDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploymentRule`: ProjectDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleApi.CreateDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectDeploymentRuleRequest** | [**ProjectDeploymentRuleRequest**](ProjectDeploymentRuleRequest.md) |  | 

### Return type

[**ProjectDeploymentRuleResponse**](ProjectDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectDeploymentRule

> DeleteProjectDeploymentRule(ctx, projectId, deploymentRuleId).Execute()

Delete a project deployment rule

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
    projectId := TODO // string | Project ID
    deploymentRuleId := TODO // string | Deployment Rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectDeploymentRuleApi.DeleteProjectDeploymentRule(context.Background(), projectId, deploymentRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleApi.DeleteProjectDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**deploymentRuleId** | [**string**](.md) | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectDeploymentRuleRequest struct via the builder pattern


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


## EditProjectDeployemtnRule

> ProjectDeploymentRuleResponse EditProjectDeployemtnRule(ctx, projectId, deploymentRuleId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()

Edit a project deployment rule

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
    projectId := TODO // string | Project ID
    deploymentRuleId := TODO // string | Deployment Rule ID
    projectDeploymentRuleRequest := *openapiclient.NewProjectDeploymentRuleRequest("^feat", "Name_example", "Mode_example", "Cluster_example", false) // ProjectDeploymentRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectDeploymentRuleApi.EditProjectDeployemtnRule(context.Background(), projectId, deploymentRuleId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleApi.EditProjectDeployemtnRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProjectDeployemtnRule`: ProjectDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleApi.EditProjectDeployemtnRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**deploymentRuleId** | [**string**](.md) | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectDeployemtnRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectDeploymentRuleRequest** | [**ProjectDeploymentRuleRequest**](ProjectDeploymentRuleRequest.md) |  | 

### Return type

[**ProjectDeploymentRuleResponse**](ProjectDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDeploymentRule

> ProjectDeploymentRuleResponse GetProjectDeploymentRule(ctx, projectId, deploymentRuleId).Execute()

Get project deployment rule

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
    projectId := TODO // string | Project ID
    deploymentRuleId := TODO // string | Deployment Rule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectDeploymentRuleApi.GetProjectDeploymentRule(context.Background(), projectId, deploymentRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleApi.GetProjectDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDeploymentRule`: ProjectDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleApi.GetProjectDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**deploymentRuleId** | [**string**](.md) | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDeploymentRuleResponse**](ProjectDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectDeploymentRule

> ProjectDeploymentRuleResponseList ListProjectDeploymentRule(ctx, projectId).Execute()

List project deployment rules

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
    projectId := TODO // string | Project ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectDeploymentRuleApi.ListProjectDeploymentRule(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleApi.ListProjectDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjectDeploymentRule`: ProjectDeploymentRuleResponseList
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleApi.ListProjectDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectDeploymentRuleResponseList**](ProjectDeploymentRuleResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

