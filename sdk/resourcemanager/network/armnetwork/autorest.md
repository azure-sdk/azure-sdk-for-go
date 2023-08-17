### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/network/resource-manager/readme.md
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/network/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module: github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4
module-version: 4.0.1
azcore-version: 1.8.0-beta.1
generate-fakes: true
inject-spans: true
```
