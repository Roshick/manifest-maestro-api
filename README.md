# Go API client for openapi

Renders Kubernetes manifests with the help of various tools such as Helm and Kustomize.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/Roshick](https://github.com/Roshick)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/Roshick/manifest-maestro-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HelmAPI* | [**PostHelmGetChartMetadataAction**](docs/HelmAPI.md#posthelmgetchartmetadataaction) | **Post** /rest/api/v1/helm/actions/get-chart-metadata | Get Helm Chart Metadata
*HelmAPI* | [**PostHelmListChartAction**](docs/HelmAPI.md#posthelmlistchartaction) | **Post** /rest/api/v1/helm/actions/list-charts | List all Helm Charts
*HelmAPI* | [**PostHelmRenderAction**](docs/HelmAPI.md#posthelmrenderaction) | **Post** /rest/api/v1/helm/actions/render-chart | Render Helm Chart
*KustomizeAPI* | [**PostKustomizeRenderKustomizationAction**](docs/KustomizeAPI.md#postkustomizerenderkustomizationaction) | **Post** /rest/api/v1/kustomize/actions/render-kustomization | Render Kustomize Kustomization
*ManagementAPI* | [**GetLivenessHealth**](docs/ManagementAPI.md#getlivenesshealth) | **Get** /health/liveness | 
*ManagementAPI* | [**GetReadinessHealth**](docs/ManagementAPI.md#getreadinesshealth) | **Get** /health/readiness | 


## Documentation For Models

 - [ErrorResponse](docs/ErrorResponse.md)
 - [GitRepositoryPathReference](docs/GitRepositoryPathReference.md)
 - [GitRepositoryReference](docs/GitRepositoryReference.md)
 - [HealthResponse](docs/HealthResponse.md)
 - [HelmChartReference](docs/HelmChartReference.md)
 - [HelmChartRepositoryChartReference](docs/HelmChartRepositoryChartReference.md)
 - [HelmChartRepositoryReference](docs/HelmChartRepositoryReference.md)
 - [HelmGetChartMetadataAction](docs/HelmGetChartMetadataAction.md)
 - [HelmGetChartMetadataActionResponse](docs/HelmGetChartMetadataActionResponse.md)
 - [HelmListChartsAction](docs/HelmListChartsAction.md)
 - [HelmListChartsActionResponse](docs/HelmListChartsActionResponse.md)
 - [HelmRenderChartAction](docs/HelmRenderChartAction.md)
 - [HelmRenderChartActionResponse](docs/HelmRenderChartActionResponse.md)
 - [HelmRenderMetadata](docs/HelmRenderMetadata.md)
 - [HelmRenderParameters](docs/HelmRenderParameters.md)
 - [HelmRepositoryReference](docs/HelmRepositoryReference.md)
 - [KustomizationReference](docs/KustomizationReference.md)
 - [KustomizeManifestInjection](docs/KustomizeManifestInjection.md)
 - [KustomizeRenderKustomizationAction](docs/KustomizeRenderKustomizationAction.md)
 - [KustomizeRenderKustomizationActionResponse](docs/KustomizeRenderKustomizationActionResponse.md)
 - [KustomizeRenderParameters](docs/KustomizeRenderParameters.md)
 - [Manifest](docs/Manifest.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

e.rieb@posteo.de

