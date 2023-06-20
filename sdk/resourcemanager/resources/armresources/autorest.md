### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/resources/resource-manager/readme.md
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/resources/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module: github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources
module-version: 2.0.0
package-resources: true
azcore-version: 1.7.0-beta.2
generate-fakes: true
inject-spans: true
```
