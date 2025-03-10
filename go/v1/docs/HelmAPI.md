# \HelmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostHelmGetChartMetadataAction**](HelmAPI.md#PostHelmGetChartMetadataAction) | **Post** /rest/api/v1/helm/actions/get-chart-metadata | Get Helm Chart Metadata
[**PostHelmListChartAction**](HelmAPI.md#PostHelmListChartAction) | **Post** /rest/api/v1/helm/actions/list-charts | List all Helm Charts
[**PostHelmRenderAction**](HelmAPI.md#PostHelmRenderAction) | **Post** /rest/api/v1/helm/actions/render-chart | Render Helm Chart



## PostHelmGetChartMetadataAction

> HelmGetChartMetadataActionResponse PostHelmGetChartMetadataAction(ctx).HelmGetChartMetadataAction(helmGetChartMetadataAction).Execute()

Get Helm Chart Metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Roshick/manifest-maestro-api/v1"
)

func main() {
	helmGetChartMetadataAction := *openapiclient.NewHelmGetChartMetadataAction(openapiclient.HelmChartReference{GitRepositoryPathReference: openapiclient.NewGitRepositoryPathReference("RepositoryType_example", "RepositoryURL_example", "GitReference_example")}) // HelmGetChartMetadataAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmAPI.PostHelmGetChartMetadataAction(context.Background()).HelmGetChartMetadataAction(helmGetChartMetadataAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmAPI.PostHelmGetChartMetadataAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHelmGetChartMetadataAction`: HelmGetChartMetadataActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmAPI.PostHelmGetChartMetadataAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHelmGetChartMetadataActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmGetChartMetadataAction** | [**HelmGetChartMetadataAction**](HelmGetChartMetadataAction.md) |  | 

### Return type

[**HelmGetChartMetadataActionResponse**](HelmGetChartMetadataActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHelmListChartAction

> HelmListChartsActionResponse PostHelmListChartAction(ctx).HelmListChartsAction(helmListChartsAction).Execute()

List all Helm Charts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Roshick/manifest-maestro-api/v1"
)

func main() {
	helmListChartsAction := *openapiclient.NewHelmListChartsAction(openapiclient.HelmRepositoryReference{GitRepositoryReference: openapiclient.NewGitRepositoryReference("RepositoryType_example", "RepositoryURL_example", "GitReference_example")}) // HelmListChartsAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmAPI.PostHelmListChartAction(context.Background()).HelmListChartsAction(helmListChartsAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmAPI.PostHelmListChartAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHelmListChartAction`: HelmListChartsActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmAPI.PostHelmListChartAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHelmListChartActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmListChartsAction** | [**HelmListChartsAction**](HelmListChartsAction.md) |  | 

### Return type

[**HelmListChartsActionResponse**](HelmListChartsActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHelmRenderAction

> HelmRenderChartActionResponse PostHelmRenderAction(ctx).HelmRenderChartAction(helmRenderChartAction).Execute()

Render Helm Chart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Roshick/manifest-maestro-api/v1"
)

func main() {
	helmRenderChartAction := *openapiclient.NewHelmRenderChartAction(openapiclient.HelmChartReference{GitRepositoryPathReference: openapiclient.NewGitRepositoryPathReference("RepositoryType_example", "RepositoryURL_example", "GitReference_example")}) // HelmRenderChartAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmAPI.PostHelmRenderAction(context.Background()).HelmRenderChartAction(helmRenderChartAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmAPI.PostHelmRenderAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHelmRenderAction`: HelmRenderChartActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmAPI.PostHelmRenderAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHelmRenderActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmRenderChartAction** | [**HelmRenderChartAction**](HelmRenderChartAction.md) |  | 

### Return type

[**HelmRenderChartActionResponse**](HelmRenderChartActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

