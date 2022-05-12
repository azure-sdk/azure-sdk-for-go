# Unreleased

## Additive Changes

### New Funcs

1. *ExpressRouteProviderPort.UnmarshalJSON([]byte) error
1. BaseClient.ExpressRouteProviderPortMethod(context.Context, string) (ExpressRouteProviderPort, error)
1. BaseClient.ExpressRouteProviderPortMethodPreparer(context.Context, string) (*http.Request, error)
1. BaseClient.ExpressRouteProviderPortMethodResponder(*http.Response) (ExpressRouteProviderPort, error)
1. BaseClient.ExpressRouteProviderPortMethodSender(*http.Request) (*http.Response, error)
1. ExpressRouteProviderPort.MarshalJSON() ([]byte, error)
1. ExpressRouteProviderPortListResult.MarshalJSON() ([]byte, error)
1. ExpressRouteProviderPortProperties.MarshalJSON() ([]byte, error)
1. ExpressRouteProviderPortsLocationClient.List(context.Context, string) (ExpressRouteProviderPortListResult, error)
1. ExpressRouteProviderPortsLocationClient.ListPreparer(context.Context, string) (*http.Request, error)
1. ExpressRouteProviderPortsLocationClient.ListResponder(*http.Response) (ExpressRouteProviderPortListResult, error)
1. ExpressRouteProviderPortsLocationClient.ListSender(*http.Request) (*http.Response, error)
1. NewExpressRouteProviderPortsLocationClient(string) ExpressRouteProviderPortsLocationClient
1. NewExpressRouteProviderPortsLocationClientWithBaseURI(string, string) ExpressRouteProviderPortsLocationClient

### Struct Changes

#### New Structs

1. ExpressRouteProviderPort
1. ExpressRouteProviderPortListResult
1. ExpressRouteProviderPortProperties
1. ExpressRouteProviderPortsLocationClient
