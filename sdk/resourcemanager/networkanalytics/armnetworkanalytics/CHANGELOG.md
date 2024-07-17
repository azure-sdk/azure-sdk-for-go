# Release History

## 2.0.0 (2024-07-17)
### Breaking Changes

- Function `*DataProductsClient.ListRolesAssignments` parameter(s) have been changed from `(context.Context, string, string, any, *DataProductsClientListRolesAssignmentsOptions)` to `(context.Context, string, string, ListRolesAssignmentsRequest, *DataProductsClientListRolesAssignmentsOptions)`
- Function `*DataTypesClient.BeginDeleteData` parameter(s) have been changed from `(context.Context, string, string, string, any, *DataTypesClientBeginDeleteDataOptions)` to `(context.Context, string, string, string, DeleteDataRequest, *DataTypesClientBeginDeleteDataOptions)`
- Type of `DataProduct.Identity` has been changed from `*ManagedServiceIdentity` to `*ManagedServiceIdentityV4`
- Type of `DataProductUpdate.Identity` has been changed from `*ManagedServiceIdentity` to `*ManagedServiceIdentityV4`
- `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Struct `ManagedServiceIdentity` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAndUserAssigned` added to enum type `ManagedServiceIdentityType`
- New struct `DeleteDataRequest`
- New struct `ListRolesAssignmentsRequest`
- New struct `ManagedServiceIdentityV4`


## 1.0.0 (2024-01-26)
### Other Changes

- Release stable version.


## 0.1.0 (2023-11-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).