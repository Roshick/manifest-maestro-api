# \KustomizeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostKustomizeRenderKustomizationAction**](KustomizeAPI.md#PostKustomizeRenderKustomizationAction) | **Post** /rest/api/v1/kustomize/actions/render-kustomization | Render Kustomize Kustomization



## PostKustomizeRenderKustomizationAction

> KustomizeRenderKustomizationActionResponse PostKustomizeRenderKustomizationAction(ctx).KustomizeRenderKustomizationAction(kustomizeRenderKustomizationAction).Execute()

Render Kustomize Kustomization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Roshick/manifest-maestro-api"
)

func main() {
	kustomizeRenderKustomizationAction := *openapiclient.NewKustomizeRenderKustomizationAction(openapiclient.KustomizationReference{GitRepositoryPathReference: openapiclient.NewGitRepositoryPathReference("RepositoryURL_example", "Reference_example")}) // KustomizeRenderKustomizationAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KustomizeAPI.PostKustomizeRenderKustomizationAction(context.Background()).KustomizeRenderKustomizationAction(kustomizeRenderKustomizationAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KustomizeAPI.PostKustomizeRenderKustomizationAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKustomizeRenderKustomizationAction`: KustomizeRenderKustomizationActionResponse
	fmt.Fprintf(os.Stdout, "Response from `KustomizeAPI.PostKustomizeRenderKustomizationAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKustomizeRenderKustomizationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kustomizeRenderKustomizationAction** | [**KustomizeRenderKustomizationAction**](KustomizeRenderKustomizationAction.md) |  | 

### Return type

[**KustomizeRenderKustomizationActionResponse**](KustomizeRenderKustomizationActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

