# Unreleased

## Breaking Changes

### Removed Funcs

1. BaseClient.GenerateGithubAccessTokenForAppserviceCLIAsync(context.Context, AppserviceGithubTokenRequest) (AppserviceGithubToken, error)
1. BaseClient.GenerateGithubAccessTokenForAppserviceCLIAsyncPreparer(context.Context, AppserviceGithubTokenRequest) (*http.Request, error)
1. BaseClient.GenerateGithubAccessTokenForAppserviceCLIAsyncResponder(*http.Response) (AppserviceGithubToken, error)
1. BaseClient.GenerateGithubAccessTokenForAppserviceCLIAsyncSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Struct Fields

1. AppserviceGithubToken.autorest.Response

### Signature Changes

#### Struct Fields

1. AppServiceCertificateOrderPatchResourceProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason
1. AppServiceCertificateOrderProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason

## Additive Changes

### New Constants

1. ResourceNotRenewableReason.ExpirationNotInRenewalTimeRange
1. ResourceNotRenewableReason.RegistrationStatusNotSupportedForRenewal
1. ResourceNotRenewableReason.SubscriptionNotActive

### New Funcs

1. PossibleResourceNotRenewableReasonValues() []ResourceNotRenewableReason

### Struct Changes

#### New Struct Fields

1. SiteConfig.AcrUseManagedIdentityCreds
1. SiteConfig.AcrUserManagedIdentityID
