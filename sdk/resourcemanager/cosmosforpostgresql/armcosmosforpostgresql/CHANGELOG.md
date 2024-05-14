# Release History

## 1.2.0-beta.2 (2024-05-14)
### Features Added

- New enum type `AADEnabledEnum` with values `AADEnabledEnumDisabled`, `AADEnabledEnumEnabled`
- New enum type `DataEncryptionType` with values `DataEncryptionTypeAzureKeyVault`, `DataEncryptionTypeSystemAssigned`
- New enum type `IdentityType` with values `IdentityTypeSystemAssigned`, `IdentityTypeUserAssigned`
- New enum type `PasswordEnabledEnum` with values `PasswordEnabledEnumDisabled`, `PasswordEnabledEnumEnabled`
- New struct `DataEncryption`
- New struct `IdentityProperties`
- New struct `UserAssignedIdentity`
- New field `Identity` in struct `Cluster`
- New field `Identity` in struct `ClusterForUpdate`
- New field `AADAuthEnabled`, `DataEncryption`, `PasswordEnabled` in struct `ClusterProperties`


## 1.2.0-beta.1 (2024-03-22)
### Features Added

- New enum type `ActiveDirectoryAuth` with values `ActiveDirectoryAuthDisabled`, `ActiveDirectoryAuthEnabled`
- New enum type `PasswordAuth` with values `PasswordAuthDisabled`, `PasswordAuthEnabled`
- New enum type `PrincipalType` with values `PrincipalTypeGroup`, `PrincipalTypeServicePrincipal`, `PrincipalTypeUser`
- New enum type `RoleType` with values `RoleTypeAdmin`, `RoleTypeUser`
- New struct `AuthConfig`
- New struct `PromoteRequest`
- New struct `RolePropertiesExternalIdentity`
- New field `AuthConfig`, `DatabaseName`, `EnableGeoBackup` in struct `ClusterProperties`
- New field `PromoteRequest` in struct `ClustersClientBeginPromoteReadReplicaOptions`
- New field `ExternalIdentity`, `RoleType` in struct `RoleProperties`


## 1.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.0.0 (2023-09-22)
### Other Changes

- Release stable version.

## 0.1.0 (2023-06-23)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).