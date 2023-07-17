### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/compute/resource-manager/readme.md
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/compute/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module: github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5
module-version: 6.0.0
azcore-version: 1.7.0-beta.2
generate-fakes: true
inject-spans: true
```
