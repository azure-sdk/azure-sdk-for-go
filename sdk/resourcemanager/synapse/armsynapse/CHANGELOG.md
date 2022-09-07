# Release History

## 0.6.0 (2022-09-07)
### Breaking Changes

- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.DesiredState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredState` to `*DesiredState`
- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.ActualState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualState` to `*ActualState`
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateDisabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateDisabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateEnabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateEnabling` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateEnabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateUnknown` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateDisabling` has been removed
- Type alias `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualState` has been removed
- Type alias `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredState` has been removed
- Function `PossibleManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateValues` has been removed
- Function `PossibleManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateValues` has been removed

### Features Added

- New const `ActualStateDisabled`
- New const `ActualStateDisabling`
- New const `ActualStateEnabled`
- New const `ActualStateUnknown`
- New const `DesiredStateEnabled`
- New const `DesiredStateDisabled`
- New const `ActualStateEnabling`
- New type alias `ActualState`
- New type alias `DesiredState`
- New function `PossibleDesiredStateValues() []DesiredState`
- New function `PossibleActualStateValues() []ActualState`
- New field `AADObjectID` in struct `DatabasePrincipalProperties`
- New field `IsAutotuneEnabled` in struct `BigDataPoolResourceProperties`
- New field `ConfigMergeRule` in struct `SparkConfigurationInfo`
- New field `AADObjectID` in struct `ClusterPrincipalProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).