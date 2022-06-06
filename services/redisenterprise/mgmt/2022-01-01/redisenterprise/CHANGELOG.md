# Unreleased

## Additive Changes

### New Constants

1. Name.NameEnterpriseE10
1. Name.NameEnterpriseE100
1. Name.NameEnterpriseE20
1. Name.NameEnterpriseE50
1. Name.NameEnterpriseFlashF1500
1. Name.NameEnterpriseFlashF300
1. Name.NameEnterpriseFlashF700

### New Funcs

1. NewSkusClient(string) SkusClient
1. NewSkusClientWithBaseURI(string, string) SkusClient
1. PossibleNameValues() []Name
1. SkusClient.List(context.Context, string) (RegionSkuDetails, error)
1. SkusClient.ListPreparer(context.Context, string) (*http.Request, error)
1. SkusClient.ListResponder(*http.Response) (RegionSkuDetails, error)
1. SkusClient.ListSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. Capability
1. LocationInfo
1. RegionSkuDetail
1. RegionSkuDetails
1. SkuDetail
1. SkusClient
