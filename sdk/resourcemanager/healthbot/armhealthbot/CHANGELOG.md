# Release History

## 1.1.0 (2022-09-06)
### Features Added

- New function `*BotsClient.ListSecrets(context.Context, string, string, *BotsClientListSecretsOptions) (BotsClientListSecretsResponse, error)`
- New function `*BotsClient.RegenerateAPIJwtSecret(context.Context, string, string, *BotsClientRegenerateAPIJwtSecretOptions) (BotsClientRegenerateAPIJwtSecretResponse, error)`
- New struct `BotsClientListSecretsOptions`
- New struct `BotsClientListSecretsResponse`
- New struct `BotsClientRegenerateAPIJwtSecretOptions`
- New struct `BotsClientRegenerateAPIJwtSecretResponse`
- New struct `Key`
- New struct `KeyVaultProperties`
- New struct `KeysResponse`
- New field `Properties` in struct `UpdateParameters`
- New field `KeyVaultProperties` in struct `Properties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).