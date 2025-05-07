# Release History

## 0.9.0 (2025-05-07)
### Breaking Changes

- Function `*WorkspacesClient.UpdateTags` parameter(s) have been changed from `(context.Context, string, string, TagsObject, *WorkspacesClientUpdateTagsOptions)` to `(context.Context, string, string, WorkspaceTagsUpdate, *WorkspacesClientUpdateTagsOptions)`
- Type of `Provider.ProvisioningState` has been changed from `*Status` to `*ProviderStatus`
- Type of `Workspace.Identity` has been changed from `*WorkspaceIdentity` to `*ManagedServiceIdentity`
- Type of `WorkspaceResourceProperties.ProvisioningState` has been changed from `*ProvisioningStatus` to `*WorkspaceProvisioningStatus`
- Enum `ProvisioningStatus` has been removed
- Enum `ResourceIdentityType` has been removed
- Enum `Status` has been removed
- Function `*ClientFactory.NewWorkspaceClient` has been removed
- Function `NewWorkspaceClient` has been removed
- Function `*WorkspaceClient.CheckNameAvailability` has been removed
- Function `*WorkspaceClient.ListKeys` has been removed
- Function `*WorkspaceClient.RegenerateKeys` has been removed
- Struct `CheckNameAvailabilityParameters` has been removed
- Struct `CheckNameAvailabilityResult` has been removed
- Struct `OperationsList` has been removed
- Struct `TagsObject` has been removed
- Struct `WorkspaceIdentity` has been removed
- Field `OperationsList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CheckNameAvailabilityReason` with values `CheckNameAvailabilityReasonAlreadyExists`, `CheckNameAvailabilityReasonInvalid`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `ProviderStatus` with values `ProviderStatusDeleted`, `ProviderStatusDeleting`, `ProviderStatusFailed`, `ProviderStatusLaunching`, `ProviderStatusSucceeded`, `ProviderStatusUpdating`
- New enum type `WorkspaceProvisioningStatus` with values `WorkspaceProvisioningStatusCanceled`, `WorkspaceProvisioningStatusFailed`, `WorkspaceProvisioningStatusProviderDeleting`, `WorkspaceProvisioningStatusProviderLaunching`, `WorkspaceProvisioningStatusProviderProvisioning`, `WorkspaceProvisioningStatusProviderUpdating`, `WorkspaceProvisioningStatusSucceeded`
- New function `*WorkspacesClient.CheckNameAvailability(context.Context, string, CheckNameAvailabilityRequest, *WorkspacesClientCheckNameAvailabilityOptions) (WorkspacesClientCheckNameAvailabilityResponse, error)`
- New function `*WorkspacesClient.ListKeys(context.Context, string, string, *WorkspacesClientListKeysOptions) (WorkspacesClientListKeysResponse, error)`
- New function `*WorkspacesClient.RegenerateKeys(context.Context, string, string, APIKeys, *WorkspacesClientRegenerateKeysOptions) (WorkspacesClientRegenerateKeysResponse, error)`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New struct `ManagedOnBehalfOfConfiguration`
- New struct `ManagedServiceIdentity`
- New struct `MoboBrokerResource`
- New struct `OperationListResult`
- New struct `UserAssignedIdentity`
- New struct `WorkspaceTagsUpdate`
- New field `ActionType`, `Origin` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `ManagedOnBehalfOfConfiguration`, `ManagedStorageAccount` in struct `WorkspaceResourceProperties`


## 0.8.0 (2024-03-22)
### Features Added

- New enum type `KeyType` with values `KeyTypePrimary`, `KeyTypeSecondary`
- New function `*WorkspaceClient.ListKeys(context.Context, string, string, *WorkspaceClientListKeysOptions) (WorkspaceClientListKeysResponse, error)`
- New function `*WorkspaceClient.RegenerateKeys(context.Context, string, string, APIKeys, *WorkspaceClientRegenerateKeysOptions) (WorkspaceClientRegenerateKeysResponse, error)`
- New struct `APIKey`
- New struct `APIKeys`
- New struct `ListKeysResult`
- New field `APIKeyEnabled` in struct `WorkspaceResourceProperties`


## 0.7.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).