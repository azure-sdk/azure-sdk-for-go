# Release History

## 2.0.0-beta.1 (2024-09-10)
### Breaking Changes

- Function `NewClient` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`
- Type of `Encryption.CustomerManagedKeyEncryption` has been changed from `*CustomerManagedKeyEncryption` to `*EncryptionCustomerManagedKeyEncryption`
- `KindGen1` from enum `Kind` has been removed
- `NameS0`, `NameS1` from enum `Name` has been removed
- Enum `IdentityType` has been removed
- Function `*Client.NewListSubscriptionOperationsPager` has been removed
- Struct `CustomerManagedKeyEncryption` has been removed
- Struct `CustomerManagedKeyEncryptionKeyIdentity` has been removed
- Struct `Dimension` has been removed
- Struct `MetricSpecification` has been removed
- Struct `OperationDetail` has been removed
- Struct `OperationProperties` has been removed
- Struct `Operations` has been removed
- Struct `ServiceSpecification` has been removed
- Field `Operations` of struct `ClientListOperationsResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `EncryptionCustomerManagedKeyEncryptionKeyIdentityType` with values `EncryptionCustomerManagedKeyEncryptionKeyIdentityTypeDelegatedResourceIdentity`, `EncryptionCustomerManagedKeyEncryptionKeyIdentityTypeSystemAssignedIdentity`, `EncryptionCustomerManagedKeyEncryptionKeyIdentityTypeUserAssignedIdentity`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New struct `EncryptionCustomerManagedKeyEncryption`
- New struct `EncryptionCustomerManagedKeyEncryptionKeyIdentity`
- New struct `LocationsItem`
- New struct `Operation`
- New struct `OperationListResult`
- New field `Locations` in struct `AccountProperties`
- New anonymous field `OperationListResult` in struct `ClientListOperationsResponse`
- New field `ConsumedStorageUnitSizeInBytes`, `TotalStorageUnitSizeInBytes` in struct `CreatorProperties`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.0.0 (2023-08-25)
### Breaking Changes

- Type of `ManagedServiceIdentity.Type` has been changed from `*ResourceIdentityType` to `*ManagedServiceIdentityType`
- Type of `ManagedServiceIdentity.UserAssignedIdentities` has been changed from `map[string]*Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties` to `map[string]*UserAssignedIdentity`
- Enum `ResourceIdentityType` has been removed
- Struct `Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed

### Features Added

- New value `SigningKeyManagedIdentity` added to enum type `SigningKey`
- New enum type `IdentityType` with values `IdentityTypeDelegatedResourceIdentity`, `IdentityTypeSystemAssignedIdentity`, `IdentityTypeUserAssignedIdentity`
- New enum type `InfrastructureEncryption` with values `InfrastructureEncryptionDisabled`, `InfrastructureEncryptionEnabled`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New struct `CustomerManagedKeyEncryption`
- New struct `CustomerManagedKeyEncryptionKeyIdentity`
- New struct `Encryption`
- New struct `UserAssignedIdentity`
- New field `Encryption` in struct `AccountProperties`
- New field `LockAggregationType`, `SourceMdmNamespace`, `SupportedAggregationTypes` in struct `MetricSpecification`


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).