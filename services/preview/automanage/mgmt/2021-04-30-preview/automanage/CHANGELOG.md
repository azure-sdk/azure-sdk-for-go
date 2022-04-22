# Unreleased

## Additive Changes

### New Funcs

1. *ServicePrincipal.UnmarshalJSON([]byte) error
1. NewServicePrincipalsClient(string) ServicePrincipalsClient
1. NewServicePrincipalsClientWithBaseURI(string, string) ServicePrincipalsClient
1. ServicePrincipal.MarshalJSON() ([]byte, error)
1. ServicePrincipalProperties.MarshalJSON() ([]byte, error)
1. ServicePrincipalsClient.Get(context.Context) (ServicePrincipal, error)
1. ServicePrincipalsClient.GetPreparer(context.Context) (*http.Request, error)
1. ServicePrincipalsClient.GetResponder(*http.Response) (ServicePrincipal, error)
1. ServicePrincipalsClient.GetSender(*http.Request) (*http.Response, error)
1. ServicePrincipalsClient.ListBySubscription(context.Context) (ServicePrincipalListResult, error)
1. ServicePrincipalsClient.ListBySubscriptionPreparer(context.Context) (*http.Request, error)
1. ServicePrincipalsClient.ListBySubscriptionResponder(*http.Response) (ServicePrincipalListResult, error)
1. ServicePrincipalsClient.ListBySubscriptionSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. ServicePrincipal
1. ServicePrincipalListResult
1. ServicePrincipalProperties
1. ServicePrincipalsClient
