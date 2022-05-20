# Unreleased

## Additive Changes

### New Constants

1. DataEncryptionType.DataEncryptionTypeAzureKeyVault
1. DataEncryptionType.DataEncryptionTypeSystemManaged
1. ManagedServiceIdentityType.ManagedServiceIdentityTypeUserAssigned

### New Funcs

1. BackupsClient.Put(context.Context, string, string, string) (ServerBackup, error)
1. BackupsClient.PutPreparer(context.Context, string, string, string) (*http.Request, error)
1. BackupsClient.PutResponder(*http.Response) (ServerBackup, error)
1. BackupsClient.PutSender(*http.Request) (*http.Response, error)
1. Identity.MarshalJSON() ([]byte, error)
1. PossibleDataEncryptionTypeValues() []DataEncryptionType
1. PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType
1. UserAssignedIdentity.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. DataEncryption
1. Identity
1. UserAssignedIdentity

#### New Struct Fields

1. Server.Identity
1. ServerForUpdate.Identity
1. ServerProperties.DataEncryption
1. ServerPropertiesForUpdate.DataEncryption
