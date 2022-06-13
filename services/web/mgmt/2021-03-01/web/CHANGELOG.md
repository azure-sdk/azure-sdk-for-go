# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. AppServiceCertificateOrderPatchResourceProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason
1. AppServiceCertificateOrderProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason

## Additive Changes

### New Constants

1. ResourceNotRenewableReason.ResourceNotRenewableReasonExpirationNotInRenewalTimeRange
1. ResourceNotRenewableReason.ResourceNotRenewableReasonRegistrationStatusNotSupportedForRenewal
1. ResourceNotRenewableReason.ResourceNotRenewableReasonSubscriptionNotActive

### New Funcs

1. AppsClient.CreateOneDeployOperation(context.Context, string, string) (SetObject, error)
1. AppsClient.CreateOneDeployOperationPreparer(context.Context, string, string) (*http.Request, error)
1. AppsClient.CreateOneDeployOperationResponder(*http.Response) (SetObject, error)
1. AppsClient.CreateOneDeployOperationSender(*http.Request) (*http.Response, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlot(context.Context, string, string, string) (SiteAuthSettingsV2, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotPreparer(context.Context, string, string, string) (*http.Request, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotResponder(*http.Response) (SiteAuthSettingsV2, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotSender(*http.Request) (*http.Response, error)
1. AppsClient.GetOneDeployStatus(context.Context, string, string) (SetObject, error)
1. AppsClient.GetOneDeployStatusPreparer(context.Context, string, string) (*http.Request, error)
1. AppsClient.GetOneDeployStatusResponder(*http.Response) (SetObject, error)
1. AppsClient.GetOneDeployStatusSender(*http.Request) (*http.Response, error)
1. PossibleResourceNotRenewableReasonValues() []ResourceNotRenewableReason
