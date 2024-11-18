# Release History

## 1.3.0 (2024-11-18)
### Features Added

- New enum type `CloudTieringLowDiskModeState` with values `CloudTieringLowDiskModeStateDisabled`, `CloudTieringLowDiskModeStateEnabled`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `ServerAuthType` with values `ServerAuthTypeCertificate`, `ServerAuthTypeManagedIdentity`
- New enum type `ServerProvisioningStatus` with values `ServerProvisioningStatusError`, `ServerProvisioningStatusInProgress`, `ServerProvisioningStatusNotStarted`, `ServerProvisioningStatusReadySyncFunctional`, `ServerProvisioningStatusReadySyncNotFunctional`
- New function `*CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys(context.Context, string, string, string, string, *CloudEndpointsClientAfsShareMetadataCertificatePublicKeysOptions) (CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse, error)`
- New function `*RegisteredServersClient.BeginUpdate(context.Context, string, string, string, RegisteredServerUpdateParameters, *RegisteredServersClientBeginUpdateOptions) (*runtime.Poller[RegisteredServersClientUpdateResponse], error)`
- New struct `CloudEndpointAfsShareMetadataCertificatePublicKeys`
- New struct `CloudTieringLowDiskMode`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `ManagedServiceIdentity`
- New struct `RegisteredServerUpdateParameters`
- New struct `RegisteredServerUpdateProperties`
- New struct `ServerEndpointProvisioningStatus`
- New struct `ServerEndpointProvisioningStepStatus`
- New struct `UserAssignedIdentity`
- New field `LockAggregationType` in struct `OperationResourceMetricSpecification`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `ApplicationID`, `Identity` in struct `RegisteredServerCreateParametersProperties`
- New field `ActiveAuthType`, `ApplicationID`, `Identity`, `LatestApplicationID` in struct `RegisteredServerProperties`
- New field `LowDiskMode` in struct `ServerEndpointCloudTieringStatus`
- New field `ServerEndpointProvisioningStatus` in struct `ServerEndpointProperties`
- New field `Identity` in struct `Service`
- New field `ID`, `Identity`, `Name`, `SystemData`, `Type` in struct `ServiceCreateParameters`
- New field `UseIdentity` in struct `ServiceCreateParametersProperties`
- New field `UseIdentity` in struct `ServiceProperties`
- New field `Identity` in struct `ServiceUpdateParameters`
- New field `UseIdentity` in struct `ServiceUpdateProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).