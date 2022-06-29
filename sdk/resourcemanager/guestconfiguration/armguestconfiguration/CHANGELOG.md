# Release History

## 1.1.0 (2022-06-29)
### Features Added

- New function `*AssignmentsClient.VMHealthCheck(context.Context, string, string, *AssignmentsClientVMHealthCheckOptions) (AssignmentsClientVMHealthCheckResponse, error)`
- New struct `AssignmentsClientVMHealthCheckOptions`
- New struct `AssignmentsClientVMHealthCheckResponse`
- New struct `HealthCheckAssignment`
- New struct `VMHealthCheckProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).