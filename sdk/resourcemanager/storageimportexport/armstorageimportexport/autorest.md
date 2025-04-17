### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/storageimportexport/resource-manager/readme.md
- /mnt/vss/_work/1/s/azure-rest-api-specs/specification/storageimportexport/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module-version: 0.7.0

directive:
  - from: swagger-document
    where: $.parameters.Accept-Language
    transform: $['x-ms-parameter-location'] = "method"
```
