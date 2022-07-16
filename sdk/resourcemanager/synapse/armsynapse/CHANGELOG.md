# Release History

## 0.6.0 (2022-07-16)
### Breaking Changes

- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.DesiredState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredState` to `*DesiredState`
- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.ActualState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualState` to `*ActualState`
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateEnabling` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateDisabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateEnabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateDisabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateEnabled` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateUnknown` has been removed
- Const `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateDisabling` has been removed
- Function `PossibleManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateValues` has been removed
- Function `PossibleManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualStateValues` has been removed

### Features Added

- New const `ActualStateDisabling`
- New const `DesiredStateDisabled`
- New const `DesiredStateEnabled`
- New const `ActualStateEnabling`
- New const `ActualStateUnknown`
- New const `ActualStateDisabled`
- New const `ActualStateEnabled`
- New function `PossibleDesiredStateValues() []DesiredState`
- New function `PossibleActualStateValues() []ActualState`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).