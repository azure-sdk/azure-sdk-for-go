# Release History

## 1.1.0 (2022-11-16)
### Features Added

- New const `ManagedServiceIdentityTypeSystemAssigned`
- New const `ManagedServiceIdentityTypeNone`
- New const `CloudTieringLowDiskModeStateDisabled`
- New const `ManagedServiceIdentityTypeUserAssigned`
- New const `CloudTieringLowDiskModeStateEnabled`
- New const `ManagedServiceIdentityTypeSystemAssignedUserAssigned`
- New type alias `ManagedServiceIdentityType`
- New type alias `CloudTieringLowDiskModeState`
- New function `*RegisteredServersClient.BeginUpdate(context.Context, string, string, string, RegisteredServerUpdateParameters, *RegisteredServersClientBeginUpdateOptions) (*runtime.Poller[RegisteredServersClientUpdateResponse], error)`
- New function `PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType`
- New function `*CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys(context.Context, string, string, string, string, *CloudEndpointsClientAfsShareMetadataCertificatePublicKeysOptions) (CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse, error)`
- New function `PossibleCloudTieringLowDiskModeStateValues() []CloudTieringLowDiskModeState`
- New struct `CloudEndpointAfsShareMetadataCertificatePublicKeys`
- New struct `CloudEndpointsClientAfsShareMetadataCertificatePublicKeysOptions`
- New struct `CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse`
- New struct `CloudTieringLowDiskMode`
- New struct `ManagedServiceIdentity`
- New struct `RegisteredServerUpdateParameters`
- New struct `RegisteredServerUpdateParametersProperties`
- New struct `RegisteredServersClientBeginUpdateOptions`
- New struct `RegisteredServersClientUpdateResponse`
- New struct `UserAssignedIdentity`
- New field `ApplicationID` in struct `RegisteredServerCreateParametersProperties`
- New field `IdentityID` in struct `RegisteredServerProperties`
- New field `IdentityType` in struct `RegisteredServerProperties`
- New field `IdentityType` in struct `ServiceUpdateProperties`
- New field `IdentityID` in struct `ServiceUpdateProperties`
- New field `IdentityType` in struct `ServiceCreateParametersProperties`
- New field `IdentityID` in struct `ServiceCreateParametersProperties`
- New field `IdentityType` in struct `ServiceProperties`
- New field `IdentityID` in struct `ServiceProperties`
- New field `Identity` in struct `ServiceProperties`
- New field `LowDiskMode` in struct `ServerEndpointCloudTieringStatus`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).