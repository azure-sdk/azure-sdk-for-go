# Release History

## 2.0.0 (2024-07-18)
### Breaking Changes

- Function `*DataProductsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, DataProductUpdate, *DataProductsClientBeginUpdateOptions)` to `(context.Context, string, string, DataProduct, *DataProductsClientBeginUpdateOptions)`
- Function `*DataProductsClient.ListRolesAssignments` parameter(s) have been changed from `(context.Context, string, string, any, *DataProductsClientListRolesAssignmentsOptions)` to `(context.Context, string, string, ListRolesAssignmentsRequest, *DataProductsClientListRolesAssignmentsOptions)`
- Function `*DataTypesClient.BeginDeleteData` parameter(s) have been changed from `(context.Context, string, string, string, any, *DataTypesClientBeginDeleteDataOptions)` to `(context.Context, string, string, string, DeleteDataRequest, *DataTypesClientBeginDeleteDataOptions)`
- Function `*DataTypesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, DataTypeUpdate, *DataTypesClientBeginUpdateOptions)` to `(context.Context, string, string, string, DataType, *DataTypesClientBeginUpdateOptions)`
- Type of `DataProduct.Identity` has been changed from `*ManagedServiceIdentity` to `*ManagedServiceIdentityV4`
- `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Function `*DataTypesClient.NewListByDataProductPager` has been removed
- Struct `DataProductUpdate` has been removed
- Struct `DataProductUpdateProperties` has been removed
- Struct `DataTypeUpdate` has been removed
- Struct `DataTypeUpdateProperties` has been removed
- Struct `ManagedServiceIdentity` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAndUserAssigned` added to enum type `ManagedServiceIdentityType`
- New function `*DataTypesClient.NewListByParentPager(string, string, *DataTypesClientListByParentOptions) *runtime.Pager[DataTypesClientListByParentResponse]`
- New struct `DeleteDataRequest`
- New struct `ListRolesAssignmentsRequest`
- New struct `ManagedServiceIdentityV4`


## 1.0.0 (2024-01-26)
### Other Changes

- Release stable version.


## 0.1.0 (2023-11-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).