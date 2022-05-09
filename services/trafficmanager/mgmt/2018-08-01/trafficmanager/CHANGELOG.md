# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. EndpointsClient.CreateOrUpdate
	- Params
		- From: context.Context, string, string, string, string, Endpoint
		- To: context.Context, string, string, EndpointType, string, Endpoint
1. EndpointsClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context, string, string, string, string, Endpoint
		- To: context.Context, string, string, EndpointType, string, Endpoint
1. EndpointsClient.Delete
	- Params
		- From: context.Context, string, string, string, string
		- To: context.Context, string, string, EndpointType, string
1. EndpointsClient.DeletePreparer
	- Params
		- From: context.Context, string, string, string, string
		- To: context.Context, string, string, EndpointType, string
1. EndpointsClient.Get
	- Params
		- From: context.Context, string, string, string, string
		- To: context.Context, string, string, EndpointType, string
1. EndpointsClient.GetPreparer
	- Params
		- From: context.Context, string, string, string, string
		- To: context.Context, string, string, EndpointType, string
1. EndpointsClient.Update
	- Params
		- From: context.Context, string, string, string, string, Endpoint
		- To: context.Context, string, string, EndpointType, string, Endpoint
1. EndpointsClient.UpdatePreparer
	- Params
		- From: context.Context, string, string, string, string, Endpoint
		- To: context.Context, string, string, EndpointType, string, Endpoint

## Additive Changes

### New Constants

1. EndpointType.EndpointTypeAzureEndpoints
1. EndpointType.EndpointTypeExternalEndpoints
1. EndpointType.EndpointTypeNestedEndpoints

### New Funcs

1. PossibleEndpointTypeValues() []EndpointType
