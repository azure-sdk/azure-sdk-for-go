# Unreleased

## Additive Changes

### New Constants

1. RestServiceAuthenticationType.RestServiceAuthenticationTypeOAuth2ClientCredential

### New Funcs

1. *GlobalParameterListResponseIterator.Next() error
1. *GlobalParameterListResponseIterator.NextWithContext(context.Context) error
1. *GlobalParameterListResponsePage.Next() error
1. *GlobalParameterListResponsePage.NextWithContext(context.Context) error
1. GlobalParameterListResponse.IsEmpty() bool
1. GlobalParameterListResponseIterator.NotDone() bool
1. GlobalParameterListResponseIterator.Response() GlobalParameterListResponse
1. GlobalParameterListResponseIterator.Value() GlobalParameterResource
1. GlobalParameterListResponsePage.NotDone() bool
1. GlobalParameterListResponsePage.Response() GlobalParameterListResponse
1. GlobalParameterListResponsePage.Values() []GlobalParameterResource
1. GlobalParameterResource.MarshalJSON() ([]byte, error)
1. GlobalParametersClient.CreateOrUpdate(context.Context, string, string, string, GlobalParameterResource) (GlobalParameterResource, error)
1. GlobalParametersClient.CreateOrUpdatePreparer(context.Context, string, string, string, GlobalParameterResource) (*http.Request, error)
1. GlobalParametersClient.CreateOrUpdateResponder(*http.Response) (GlobalParameterResource, error)
1. GlobalParametersClient.CreateOrUpdateSender(*http.Request) (*http.Response, error)
1. GlobalParametersClient.Delete(context.Context, string, string, string) (autorest.Response, error)
1. GlobalParametersClient.DeletePreparer(context.Context, string, string, string) (*http.Request, error)
1. GlobalParametersClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. GlobalParametersClient.DeleteSender(*http.Request) (*http.Response, error)
1. GlobalParametersClient.Get(context.Context, string, string, string) (GlobalParameterResource, error)
1. GlobalParametersClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. GlobalParametersClient.GetResponder(*http.Response) (GlobalParameterResource, error)
1. GlobalParametersClient.GetSender(*http.Request) (*http.Response, error)
1. GlobalParametersClient.ListByFactory(context.Context, string, string) (GlobalParameterListResponsePage, error)
1. GlobalParametersClient.ListByFactoryComplete(context.Context, string, string) (GlobalParameterListResponseIterator, error)
1. GlobalParametersClient.ListByFactoryPreparer(context.Context, string, string) (*http.Request, error)
1. GlobalParametersClient.ListByFactoryResponder(*http.Response) (GlobalParameterListResponse, error)
1. GlobalParametersClient.ListByFactorySender(*http.Request) (*http.Response, error)
1. NewGlobalParameterListResponseIterator(GlobalParameterListResponsePage) GlobalParameterListResponseIterator
1. NewGlobalParameterListResponsePage(GlobalParameterListResponse, func(context.Context, GlobalParameterListResponse) (GlobalParameterListResponse, error)) GlobalParameterListResponsePage
1. NewGlobalParametersClient(string) GlobalParametersClient
1. NewGlobalParametersClientWithBaseURI(string, string) GlobalParametersClient

### Struct Changes

#### New Structs

1. GlobalParameterListResponse
1. GlobalParameterListResponseIterator
1. GlobalParameterListResponsePage
1. GlobalParameterResource
1. GlobalParametersClient

#### New Struct Fields

1. RestServiceLinkedServiceTypeProperties.ClientID
1. RestServiceLinkedServiceTypeProperties.ClientSecret
1. RestServiceLinkedServiceTypeProperties.Resource
1. RestServiceLinkedServiceTypeProperties.Scope
1. RestServiceLinkedServiceTypeProperties.TokenEndpoint
