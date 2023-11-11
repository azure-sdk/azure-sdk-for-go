# Release History

## 0.2.0 (2023-11-11)
### Breaking Changes

- Function `timeRFC3339.MarshalText` has been removed
- Function `*timeRFC3339.Parse` has been removed
- Function `*timeRFC3339.UnmarshalText` has been removed

### Features Added

- New function `dateTimeRFC3339.MarshalText() ([]byte, error)`
- New function `*dateTimeRFC3339.Parse(string) error`
- New function `*dateTimeRFC3339.UnmarshalText([]byte) error`


## 0.1.0 (2023-07-28)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).