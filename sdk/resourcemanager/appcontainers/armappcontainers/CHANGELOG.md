# Release History

## 2.0.0 (2022-06-24)
### Breaking Changes

- Function `*CertificatesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *CertificatesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, Certificate, *CertificatesClientCreateOrUpdateOptions)`
- Field `CertificateEnvelope` of struct `CertificatesClientCreateOrUpdateOptions` has been removed


## 1.0.0 (2022-05-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).