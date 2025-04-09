# Release History

## 1.1.0-beta.1 (2025-04-09)
### Features Added

- New enum type `InstanceFeatureMode` with values `InstanceFeatureModeDisabled`, `InstanceFeatureModePreview`, `InstanceFeatureModeStable`
- New enum type `RetainmentPolicyMode` with values `RetainmentPolicyModeAll`, `RetainmentPolicyModeCustom`, `RetainmentPolicyModeNone`
- New function `*CustomStateStoreRetainmentPolicy.GetStateStoreRetainmentPolicy() *StateStoreRetainmentPolicy`
- New function `*CustomSubscriberQueueRetainmentPolicy.GetSubscriberQueueRetainmentPolicy() *SubscriberQueueRetainmentPolicy`
- New function `*CustomTopicRetainmentPolicy.GetTopicRetainmentPolicy() *TopicRetainmentPolicy`
- New function `*StateStoreRetainmentPolicy.GetStateStoreRetainmentPolicy() *StateStoreRetainmentPolicy`
- New function `*SubscriberQueueRetainmentPolicy.GetSubscriberQueueRetainmentPolicy() *SubscriberQueueRetainmentPolicy`
- New function `*TopicRetainmentPolicy.GetTopicRetainmentPolicy() *TopicRetainmentPolicy`
- New struct `CustomStateStoreRetainmentPolicy`
- New struct `CustomSubscriberQueueRetainmentPolicy`
- New struct `CustomTopicRetainmentPolicy`
- New struct `DynamicPersistenceSettings`
- New struct `DynamicRetainmentSettings`
- New struct `InstanceFeature`
- New struct `Persistence`
- New struct `StateStoreRetainmentResources`
- New struct `StateStoreRetainmentSettings`
- New struct `SubscriberQueueRetainmentSettings`
- New struct `TopicRetainmentSettings`
- New struct `VolumeClaimResourceRequirementsClaims`
- New field `Persistence` in struct `BrokerProperties`
- New field `Features` in struct `InstanceProperties`
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