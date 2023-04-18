# Release History

## 0.8.0 (2023-04-18)
### Breaking Changes

- Field `Interface` of struct `ClientResourcesHistoryResponse` has been removed

### Features Added

- New enum type `ChangeCategory` with values `ChangeCategorySystem`, `ChangeCategoryUser`
- New enum type `ChangeType` with values `ChangeTypeCreate`, `ChangeTypeDelete`, `ChangeTypeUpdate`
- New enum type `PropertyChangeType` with values `PropertyChangeTypeInsert`, `PropertyChangeTypeRemove`, `PropertyChangeTypeUpdate`
- New function `*Client.ResourceChangeDetails(context.Context, ResourceChangeDetailsRequestParameters, *ClientResourceChangeDetailsOptions) (ClientResourceChangeDetailsResponse, error)`
- New function `*Client.ResourceChanges(context.Context, ResourceChangesRequestParameters, *ClientResourceChangesOptions) (ClientResourceChangesResponse, error)`
- New struct `ResourceChangeData`
- New struct `ResourceChangeDataAfterSnapshot`
- New struct `ResourceChangeDataBeforeSnapshot`
- New struct `ResourceChangeDetailsRequestParameters`
- New struct `ResourceChangeList`
- New struct `ResourceChangesRequestParameters`
- New struct `ResourceChangesRequestParametersInterval`
- New struct `ResourcePropertyChange`
- New struct `ResourceSnapshotData`
- New field `Value` in struct `ClientResourcesHistoryResponse`


## 0.7.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.7.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.6.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).