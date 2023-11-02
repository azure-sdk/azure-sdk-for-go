# Release History

## 2.0.0-beta.1 (2023-11-02)
### Breaking Changes

- Enum `PrivateLinkServiceConnectionProvisioningState` has been removed
- Function `*ClientFactory.NewManagementClient` has been removed
- Function `*ClientFactory.NewUsagesClient` has been removed
- Function `NewManagementClient` has been removed
- Function `*ManagementClient.UsageBySubscriptionSKU` has been removed
- Function `NewUsagesClient` has been removed
- Function `*UsagesClient.NewListBySubscriptionPager` has been removed
- Struct `QuotaUsageResult` has been removed
- Struct `QuotaUsageResultName` has been removed
- Struct `QuotaUsagesListResult` has been removed
- Field `GroupID`, `ProvisioningState` of struct `PrivateEndpointConnectionProperties` has been removed

### Features Added

- New value `IdentityTypeSystemAssignedUserAssigned`, `IdentityTypeUserAssigned` added to enum type `IdentityType`
- New value `SearchServiceStatusStopped` added to enum type `SearchServiceStatus`
- New enum type `SearchBypass` with values `SearchBypassAzurePortal`, `SearchBypassNone`
- New enum type `SearchDisabledDataExfiltrationOption` with values `SearchDisabledDataExfiltrationOptionAll`
- New struct `OperationAvailability`
- New struct `OperationLogsSpecification`
- New struct `OperationMetricDimension`
- New struct `OperationMetricsSpecification`
- New struct `OperationProperties`
- New struct `OperationServiceSpecification`
- New struct `UserAssignedManagedIdentity`
- New field `UserAssignedIdentities` in struct `Identity`
- New field `Bypass` in struct `NetworkRuleSet`
- New field `IsDataAction`, `Origin`, `Properties` in struct `Operation`
- New field `DisabledDataExfiltrationOptions`, `ETag` in struct `ServiceProperties`


## 1.2.0 (2023-10-27)
### Features Added

- New enum type `AADAuthFailureMode` with values `AADAuthFailureModeHttp401WithBearerChallenge`, `AADAuthFailureModeHttp403`
- New enum type `PrivateLinkServiceConnectionProvisioningState` with values `PrivateLinkServiceConnectionProvisioningStateCanceled`, `PrivateLinkServiceConnectionProvisioningStateDeleting`, `PrivateLinkServiceConnectionProvisioningStateFailed`, `PrivateLinkServiceConnectionProvisioningStateIncomplete`, `PrivateLinkServiceConnectionProvisioningStateSucceeded`, `PrivateLinkServiceConnectionProvisioningStateUpdating`
- New enum type `SearchEncryptionComplianceStatus` with values `SearchEncryptionComplianceStatusCompliant`, `SearchEncryptionComplianceStatusNonCompliant`
- New enum type `SearchEncryptionWithCmk` with values `SearchEncryptionWithCmkDisabled`, `SearchEncryptionWithCmkEnabled`, `SearchEncryptionWithCmkUnspecified`
- New enum type `SearchSemanticSearch` with values `SearchSemanticSearchDisabled`, `SearchSemanticSearchFree`, `SearchSemanticSearchStandard`
- New function `*ClientFactory.NewManagementClient() *ManagementClient`
- New function `*ClientFactory.NewUsagesClient() *UsagesClient`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `*ManagementClient.UsageBySubscriptionSKU(context.Context, string, string, *SearchManagementRequestOptions, *ManagementClientUsageBySubscriptionSKUOptions) (ManagementClientUsageBySubscriptionSKUResponse, error)`
- New function `NewUsagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UsagesClient, error)`
- New function `*UsagesClient.NewListBySubscriptionPager(string, *SearchManagementRequestOptions, *UsagesClientListBySubscriptionOptions) *runtime.Pager[UsagesClientListBySubscriptionResponse]`
- New struct `DataPlaneAADOrAPIKeyAuthOption`
- New struct `DataPlaneAuthOptions`
- New struct `EncryptionWithCmk`
- New struct `QuotaUsageResult`
- New struct `QuotaUsageResultName`
- New struct `QuotaUsagesListResult`
- New field `GroupID`, `ProvisioningState` in struct `PrivateEndpointConnectionProperties`
- New field `AuthOptions`, `DisableLocalAuth`, `EncryptionWithCmk`, `SemanticSearch` in struct `ServiceProperties`


## 1.1.0 (2023-04-03)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).