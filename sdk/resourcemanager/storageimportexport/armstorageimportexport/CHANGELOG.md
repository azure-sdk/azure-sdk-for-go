# Release History

## 1.0.0 (2023-04-03)
### Features Added

- New function `NewClientFactory(string, *string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewBitLockerKeysClient() *BitLockerKeysClient`
- New function `*ClientFactory.NewJobsClient() *JobsClient`
- New function `*ClientFactory.NewLocationsClient() *LocationsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New struct `ClientFactory`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).