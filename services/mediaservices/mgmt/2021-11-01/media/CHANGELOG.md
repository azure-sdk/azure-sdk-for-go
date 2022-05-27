# Unreleased

## Additive Changes

### New Constants

1. AsyncOperationStatus.AsyncOperationStatusFailed
1. AsyncOperationStatus.AsyncOperationStatusInProgress
1. AsyncOperationStatus.AsyncOperationStatusSucceeded

### New Funcs

1. LiveEventsClient.AsyncOperation(context.Context, string, string, string) (AsyncOperationResult, error)
1. LiveEventsClient.AsyncOperationPreparer(context.Context, string, string, string) (*http.Request, error)
1. LiveEventsClient.AsyncOperationResponder(*http.Response) (AsyncOperationResult, error)
1. LiveEventsClient.AsyncOperationSender(*http.Request) (*http.Response, error)
1. LiveEventsClient.OperationLocation(context.Context, string, string, string, string) (LiveEvent, error)
1. LiveEventsClient.OperationLocationPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. LiveEventsClient.OperationLocationResponder(*http.Response) (LiveEvent, error)
1. LiveEventsClient.OperationLocationSender(*http.Request) (*http.Response, error)
1. LiveOutputsClient.AsyncOperation(context.Context, string, string, string) (AsyncOperationResult, error)
1. LiveOutputsClient.AsyncOperationPreparer(context.Context, string, string, string) (*http.Request, error)
1. LiveOutputsClient.AsyncOperationResponder(*http.Response) (AsyncOperationResult, error)
1. LiveOutputsClient.AsyncOperationSender(*http.Request) (*http.Response, error)
1. LiveOutputsClient.OperationLocation(context.Context, string, string, string, string, string) (LiveOutput, error)
1. LiveOutputsClient.OperationLocationPreparer(context.Context, string, string, string, string, string) (*http.Request, error)
1. LiveOutputsClient.OperationLocationResponder(*http.Response) (LiveOutput, error)
1. LiveOutputsClient.OperationLocationSender(*http.Request) (*http.Response, error)
1. PossibleAsyncOperationStatusValues() []AsyncOperationStatus
1. StreamingEndpointsClient.AsyncOperation(context.Context, string, string, string) (AsyncOperationResult, error)
1. StreamingEndpointsClient.AsyncOperationPreparer(context.Context, string, string, string) (*http.Request, error)
1. StreamingEndpointsClient.AsyncOperationResponder(*http.Response) (AsyncOperationResult, error)
1. StreamingEndpointsClient.AsyncOperationSender(*http.Request) (*http.Response, error)
1. StreamingEndpointsClient.OperationLocation(context.Context, string, string, string, string) (StreamingEndpoint, error)
1. StreamingEndpointsClient.OperationLocationPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. StreamingEndpointsClient.OperationLocationResponder(*http.Response) (StreamingEndpoint, error)
1. StreamingEndpointsClient.OperationLocationSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. AsyncOperationErrorDetail
1. AsyncOperationResult
