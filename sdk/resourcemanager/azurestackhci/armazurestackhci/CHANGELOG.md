# Release History

## 1.1.0 (2022-09-02)
### Features Added

- New const `ClusterNodeTypeFirstParty`
- New const `SoftwareAssuranceIntentEnable`
- New const `ClusterNodeTypeThirdParty`
- New const `SoftwareAssuranceStatusEnabled`
- New const `SoftwareAssuranceStatusDisabled`
- New const `SoftwareAssuranceIntentDisable`
- New type alias `SoftwareAssuranceIntent`
- New type alias `ClusterNodeType`
- New type alias `SoftwareAssuranceStatus`
- New function `PossibleClusterNodeTypeValues() []ClusterNodeType`
- New function `PossibleSoftwareAssuranceStatusValues() []SoftwareAssuranceStatus`
- New function `PossibleSoftwareAssuranceIntentValues() []SoftwareAssuranceIntent`
- New function `*ClustersClient.BeginExtendSoftwareAssuranceBenefit(context.Context, string, string, SoftwareAssuranceChangeRequest, *ClustersClientBeginExtendSoftwareAssuranceBenefitOptions) (*runtime.Poller[ClustersClientExtendSoftwareAssuranceBenefitResponse], error)`
- New struct `ClustersClientBeginExtendSoftwareAssuranceBenefitOptions`
- New struct `ClustersClientExtendSoftwareAssuranceBenefitResponse`
- New struct `SoftwareAssuranceChangeRequest`
- New struct `SoftwareAssuranceChangeRequestProperties`
- New struct `SoftwareAssuranceProperties`
- New field `SoftwareAssuranceProperties` in struct `ClusterProperties`
- New field `NodeType` in struct `ClusterNode`
- New field `OSDisplayVersion` in struct `ClusterNode`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).