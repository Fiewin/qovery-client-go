# \EnvironmentSecretApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentSecret**](EnvironmentSecretApi.md#CreateEnvironmentSecret) | **Post** /environment/{environmentId}/secret | Add a secret to the environment
[**CreateEnvironmentSecretAlias**](EnvironmentSecretApi.md#CreateEnvironmentSecretAlias) | **Post** /environment/{environmentId}/secret/{secretId}/alias | Create a secret alias at the environment level
[**CreateEnvironmentSecretOverride**](EnvironmentSecretApi.md#CreateEnvironmentSecretOverride) | **Post** /environment/{environmentId}/secret/{secretId}/override | Create a secret override at the environment level
[**EditEnvironmentSecret**](EnvironmentSecretApi.md#EditEnvironmentSecret) | **Put** /environment/{environmentId}/secret/{secretId} | Edit a secret belonging to the environment
[**EnvironmentEnvironmentIdSecretSecretIdDelete**](EnvironmentSecretApi.md#EnvironmentEnvironmentIdSecretSecretIdDelete) | **Delete** /environment/{environmentId}/secret/{secretId} | Delete a secret from the environment
[**ListEnvironmentSecrets**](EnvironmentSecretApi.md#ListEnvironmentSecrets) | **Get** /environment/{environmentId}/secret | List environment secrets



## CreateEnvironmentSecret

> SecretResponse CreateEnvironmentSecret(ctx, environmentId).SecretRequest(secretRequest).Execute()

Add a secret to the environment



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
    environmentId := TODO // string | Environment ID
    secretRequest := *openapiclient.NewSecretRequest("Key_example", "Value_example") // SecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.CreateEnvironmentSecret(context.Background(), environmentId).SecretRequest(secretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.CreateEnvironmentSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentSecretApi.CreateEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentSecretRequest struct via the builder pattern


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


## CreateEnvironmentSecretAlias

> SecretResponse CreateEnvironmentSecretAlias(ctx, environmentId, secretId).Key(key).Execute()

Create a secret alias at the environment level



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
    environmentId := TODO // string | Environment ID
    secretId := TODO // string | Secret ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.CreateEnvironmentSecretAlias(context.Background(), environmentId, secretId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.CreateEnvironmentSecretAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentSecretAlias`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentSecretApi.CreateEnvironmentSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentSecretAliasRequest struct via the builder pattern


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


## CreateEnvironmentSecretOverride

> SecretResponse CreateEnvironmentSecretOverride(ctx, environmentId, secretId).Value(value).Execute()

Create a secret override at the environment level



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
    environmentId := TODO // string | Environment ID
    secretId := TODO // string | Secret ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.CreateEnvironmentSecretOverride(context.Background(), environmentId, secretId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.CreateEnvironmentSecretOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentSecretOverride`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentSecretApi.CreateEnvironmentSecretOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentSecretOverrideRequest struct via the builder pattern


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


## EditEnvironmentSecret

> SecretResponse EditEnvironmentSecret(ctx, environmentId, secretId).SecretEditRequest(secretEditRequest).Execute()

Edit a secret belonging to the environment



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
    environmentId := TODO // string | Environment ID
    secretId := TODO // string | Secret ID
    secretEditRequest := *openapiclient.NewSecretEditRequest() // SecretEditRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.EditEnvironmentSecret(context.Background(), environmentId, secretId).SecretEditRequest(secretEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.EditEnvironmentSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEnvironmentSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentSecretApi.EditEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentSecretRequest struct via the builder pattern


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


## EnvironmentEnvironmentIdSecretSecretIdDelete

> EnvironmentEnvironmentIdSecretSecretIdDelete(ctx, environmentId, secretId).Execute()

Delete a secret from the environment



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
    environmentId := TODO // string | Environment ID
    secretId := TODO // string | Secret ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.EnvironmentEnvironmentIdSecretSecretIdDelete(context.Background(), environmentId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.EnvironmentEnvironmentIdSecretSecretIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentEnvironmentIdSecretSecretIdDeleteRequest struct via the builder pattern


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


## ListEnvironmentSecrets

> SecretResponseList ListEnvironmentSecrets(ctx, environmentId).Execute()

List environment secrets

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
    environmentId := TODO // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentSecretApi.ListEnvironmentSecrets(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentSecretApi.ListEnvironmentSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentSecrets`: SecretResponseList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentSecretApi.ListEnvironmentSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentSecretsRequest struct via the builder pattern


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

