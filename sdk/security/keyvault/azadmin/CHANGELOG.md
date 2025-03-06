# Release History

## 1.3.1-beta.1 (2025-03-06)
### Other Changes


## 2.0.0-beta.1 (2025-03-06)
### Breaking Changes

- Function `*Client.CreateOrUpdateRoleDefinition` parameter(s) have been changed from `(context.Context, RoleScope, string, RoleDefinitionCreateParameters, *CreateOrUpdateRoleDefinitionOptions)` to `(context.Context, string, string, RoleDefinitionCreateParameters, *CreateOrUpdateRoleDefinitionOptions)`
- Function `*Client.CreateRoleAssignment` parameter(s) have been changed from `(context.Context, RoleScope, string, RoleAssignmentCreateParameters, *CreateRoleAssignmentOptions)` to `(context.Context, string, string, RoleAssignmentCreateParameters, *CreateRoleAssignmentOptions)`
- Function `*Client.DeleteRoleAssignment` parameter(s) have been changed from `(context.Context, RoleScope, string, *DeleteRoleAssignmentOptions)` to `(context.Context, string, string, *DeleteRoleAssignmentOptions)`
- Function `*Client.DeleteRoleDefinition` parameter(s) have been changed from `(context.Context, RoleScope, string, *DeleteRoleDefinitionOptions)` to `(context.Context, string, string, *DeleteRoleDefinitionOptions)`
- Function `*Client.GetRoleAssignment` parameter(s) have been changed from `(context.Context, RoleScope, string, *GetRoleAssignmentOptions)` to `(context.Context, string, string, *GetRoleAssignmentOptions)`
- Function `*Client.GetRoleDefinition` parameter(s) have been changed from `(context.Context, RoleScope, string, *GetRoleDefinitionOptions)` to `(context.Context, string, string, *GetRoleDefinitionOptions)`
- Function `*Client.NewListRoleAssignmentsPager` parameter(s) have been changed from `(RoleScope, *ListRoleAssignmentsOptions)` to `(string, *ListRoleAssignmentsOptions)`
- Function `*Client.NewListRoleDefinitionsPager` parameter(s) have been changed from `(RoleScope, *ListRoleDefinitionsOptions)` to `(string, *ListRoleDefinitionsOptions)`


## 2.0.0-beta.1 (2025-03-06)
### Breaking Changes

- Type of `FullBackupOperation.Error` has been changed from `*ErrorInfo` to `*FullBackupOperationError`
- Type of `FullBackupOperation.Status` has been changed from `*string` to `*OperationStatus`
- Type of `RestoreOperation.Error` has been changed from `*ErrorInfo` to `*FullBackupOperationError`
- Type of `RestoreOperation.Status` has been changed from `*string` to `*OperationStatus`
- Type of `SelectiveKeyRestoreOperation.Error` has been changed from `*ErrorInfo` to `*FullBackupOperationError`
- Type of `SelectiveKeyRestoreOperation.Status` has been changed from `*string` to `*OperationStatus`

### Features Added

- New enum type `OperationStatus` with values `OperationStatusCanceled`, `OperationStatusFailed`, `OperationStatusInProgress`, `OperationStatusSucceeded`
- New function `*Client.BeginPreFullBackup(context.Context, PreBackupOperationParameters, *BeginPreFullBackupOptions) (*runtime.Poller[PreFullBackupResponse], error)`
- New function `*Client.BeginPreFullRestore(context.Context, PreRestoreOperationParameters, *BeginPreFullRestoreOptions) (*runtime.Poller[PreFullRestoreResponse], error)`
- New struct `FullBackupOperationError`
- New struct `PreBackupOperationParameters`
- New struct `PreRestoreOperationParameters`


### 1.3.1 (2025-02-13)

#### Other Changes
* Upgraded dependencies

### 1.3.0 (2024-11-13)

#### Features Added
* Added API Version support. Users can now change the default API Version by setting ClientOptions.APIVersion

### 1.2.0 (2024-10-21)

#### Features Added
* Added CAE support
* Client requests tokens from the Vault's tenant, overriding any credential default
  (thanks @francescomari)

### 1.1.0 (2024-02-13)

#### Other Changes
* Upgraded to API service version `7.5`
* Upgraded dependencies

### 1.1.0-beta.1 (2023-11-09)

#### Features Added
* Managed Identity can now be used in place of a SAS token to access the blob storage resource when performing backup and restore operations.

#### Other Changes
* Upgraded service version to `7.5-preview.1`
* Updated to latest version of `azcore`.
* Enabled spans for distributed tracing.

### 1.0.1 (2023-08-24)

#### Other Changes
* Upgraded dependencies 

### 1.0.0 (2023-07-17)

#### Features Added
* First stable release of the azadmin module

### 0.3.0 (2023-06-08)

### Breatking Changes
* Renamed `SASTokenParameter` to `SASTokenParameters`
* Renamed `RestoreOperationParameters.SasTokenParameters` to `RestoreOperationParameters.SASTokenParameters`

### Other Changes
* Updated dependencies

### 0.2.0 (2023-04-13)

#### Breaking Changes
* Renamed `ServerError` to `ErrorInfo`

### 0.1.0 (2023-03-13)
* This is the initial release of the `azadmin` library
