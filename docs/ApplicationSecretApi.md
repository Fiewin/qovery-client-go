# \ApplicationSecretApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationSecret**](ApplicationSecretApi.md#CreateApplicationSecret) | **Post** /application/{applicationId}/secret | Add a secret to the application
[**CreateApplicationSecretAlias**](ApplicationSecretApi.md#CreateApplicationSecretAlias) | **Post** /application/{applicationId}/secret/{secretId}/alias | Create a secret alias at the application level
[**CreateApplicationSecretOverride**](ApplicationSecretApi.md#CreateApplicationSecretOverride) | **Post** /application/{applicationId}/secret/{secretId}/override | Create a secret override at the application level
[**DeleteApplicationSecret**](ApplicationSecretApi.md#DeleteApplicationSecret) | **Delete** /application/{applicationId}/secret/{secretId} | Delete a secret from an application
[**EditApplicationSecret**](ApplicationSecretApi.md#EditApplicationSecret) | **Put** /application/{applicationId}/secret/{secretId} | Edit a secret belonging to the application
[**ListApplicationSecrets**](ApplicationSecretApi.md#ListApplicationSecrets) | **Get** /application/{applicationId}/secret | List application secrets



## CreateApplicationSecret

> SecretResponse CreateApplicationSecret(ctx, applicationId).SecretRequest(secretRequest).Execute()

Add a secret to the application



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
    secretRequest := *openapiclient.NewSecretRequest("Key_example", "Value_example") // SecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationSecretApi.CreateApplicationSecret(context.Background(), applicationId).SecretRequest(secretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.CreateApplicationSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSecretApi.CreateApplicationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretRequest** | [**SecretRequest**](SecretRequest.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationSecretAlias

> SecretResponse CreateApplicationSecretAlias(ctx, applicationId, secretId).Key(key).Execute()

Create a secret alias at the application level



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
    secretId := TODO // string | Secret ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationSecretApi.CreateApplicationSecretAlias(context.Background(), applicationId, secretId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.CreateApplicationSecretAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationSecretAlias`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSecretApi.CreateApplicationSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationSecretAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationSecretOverride

> SecretResponse CreateApplicationSecretOverride(ctx, applicationId, secretId).Value(value).Execute()

Create a secret override at the application level



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
    secretId := TODO // string | Secret ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationSecretApi.CreateApplicationSecretOverride(context.Background(), applicationId, secretId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.CreateApplicationSecretOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationSecretOverride`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSecretApi.CreateApplicationSecretOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationSecretOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationSecret

> DeleteApplicationSecret(ctx, applicationId, secretId).Execute()

Delete a secret from an application



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
    secretId := TODO // string | Secret ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationSecretApi.DeleteApplicationSecret(context.Background(), applicationId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.DeleteApplicationSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationSecretRequest struct via the builder pattern


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


## EditApplicationSecret

> SecretResponse EditApplicationSecret(ctx, applicationId, secretId).SecretEditRequest(secretEditRequest).Execute()

Edit a secret belonging to the application



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
    secretId := TODO // string | Secret ID
    secretEditRequest := *openapiclient.NewSecretEditRequest() // SecretEditRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationSecretApi.EditApplicationSecret(context.Background(), applicationId, secretId).SecretEditRequest(secretEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.EditApplicationSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSecretApi.EditApplicationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretEditRequest** | [**SecretEditRequest**](SecretEditRequest.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationSecrets

> SecretResponseList ListApplicationSecrets(ctx, applicationId).Execute()

List application secrets



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
    resp, r, err := api_client.ApplicationSecretApi.ListApplicationSecrets(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecretApi.ListApplicationSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationSecrets`: SecretResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSecretApi.ListApplicationSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretResponseList**](SecretResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

