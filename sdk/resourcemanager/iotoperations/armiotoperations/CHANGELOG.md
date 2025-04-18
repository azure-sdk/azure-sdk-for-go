# Release History

## 1.1.0-beta.1 (2025-04-18)
### Features Added

- New value `EndpointTypeEventGrid`, `EndpointTypeEventhub`, `EndpointTypeFabricRT`, `EndpointTypeLocalMq`, `EndpointTypeOtel` added to enum type `EndpointType`
- New enum type `InstanceFeatureMode` with values `InstanceFeatureModeDisabled`, `InstanceFeatureModePreview`, `InstanceFeatureModeStable`
- New enum type `RemoteSupportAccessLevels` with values `RemoteSupportAccessLevelsDiagnose`, `RemoteSupportAccessLevelsDiagnoseAndRepair`
- New enum type `RemoteSupportActivationState` with values `RemoteSupportActivationStateDisabled`, `RemoteSupportActivationStateEnabled`, `RemoteSupportActivationStateExpired`
- New enum type `RetainmentPolicyMode` with values `RetainmentPolicyModeAll`, `RetainmentPolicyModeCustom`, `RetainmentPolicyModeNone`
- New function `*ClientFactory.NewDiagnosticClient() *DiagnosticClient`
- New function `*CustomStateStoreRetainmentPolicy.GetStateStoreRetainmentPolicy() *StateStoreRetainmentPolicy`
- New function `*CustomSubscriberQueueRetainmentPolicy.GetSubscriberQueueRetainmentPolicy() *SubscriberQueueRetainmentPolicy`
- New function `*CustomTopicRetainmentPolicy.GetTopicRetainmentPolicy() *TopicRetainmentPolicy`
- New function `NewDiagnosticClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DiagnosticClient, error)`
- New function `*DiagnosticClient.BeginCreateOrUpdate(context.Context, string, string, string, DiagnosticResource, *DiagnosticClientBeginCreateOrUpdateOptions) (*runtime.Poller[DiagnosticClientCreateOrUpdateResponse], error)`
- New function `*DiagnosticClient.BeginDelete(context.Context, string, string, string, *DiagnosticClientBeginDeleteOptions) (*runtime.Poller[DiagnosticClientDeleteResponse], error)`
- New function `*DiagnosticClient.Get(context.Context, string, string, string, *DiagnosticClientGetOptions) (DiagnosticClientGetResponse, error)`
- New function `*DiagnosticClient.NewListByResourceGroupPager(string, string, *DiagnosticClientListByResourceGroupOptions) *runtime.Pager[DiagnosticClientListByResourceGroupResponse]`
- New function `*StateStoreRetainmentPolicy.GetStateStoreRetainmentPolicy() *StateStoreRetainmentPolicy`
- New function `*SubscriberQueueRetainmentPolicy.GetSubscriberQueueRetainmentPolicy() *SubscriberQueueRetainmentPolicy`
- New function `*TopicRetainmentPolicy.GetTopicRetainmentPolicy() *TopicRetainmentPolicy`
- New struct `CustomStateStoreRetainmentPolicy`
- New struct `CustomSubscriberQueueRetainmentPolicy`
- New struct `CustomTopicRetainmentPolicy`
- New struct `DataflowEndpointOtel`
- New struct `DiagnosticProperties`
- New struct `DiagnosticResource`
- New struct `DiagnosticResourceListResult`
- New struct `DynamicPersistenceSettings`
- New struct `DynamicRetainmentSettings`
- New struct `InstanceFeature`
- New struct `Persistence`
- New struct `RemoteSupportProperties`
- New struct `StateStoreRetainmentResources`
- New struct `StateStoreRetainmentSettings`
- New struct `SubscriberQueueRetainmentSettings`
- New struct `TopicRetainmentSettings`
- New struct `VolumeClaimResourceRequirementsClaims`
- New field `Persistence` in struct `BrokerProperties`
- New field `OtelSettings` in struct `DataflowEndpointProperties`
- New field `AdrNamespace`, `Features`, `SecretProviderClassRef` in struct `InstanceProperties`
- New field `Claims` in struct `VolumeClaimResourceRequirements`


## 1.0.0 (2024-12-12)
### Breaking Changes

- `ManagedServiceIdentityTypeSystemAndUserAssigned` from enum `ManagedServiceIdentityType` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAssignedUserAssigned` added to enum type `ManagedServiceIdentityType`


## 0.1.0 (2024-10-24)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).