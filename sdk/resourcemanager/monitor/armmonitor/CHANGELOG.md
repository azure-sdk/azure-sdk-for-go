# Release History

## 0.8.0 (2022-06-08)
### Breaking Changes

- Field `AzureAppPushReceivers` of struct `NotificationRequestBody` has been removed
- Field `AzureFunctionReceivers` of struct `NotificationRequestBody` has been removed
- Field `EventHubReceivers` of struct `NotificationRequestBody` has been removed
- Field `SmsReceivers` of struct `NotificationRequestBody` has been removed
- Field `AlertType` of struct `NotificationRequestBody` has been removed
- Field `ItsmReceivers` of struct `NotificationRequestBody` has been removed
- Field `VoiceReceivers` of struct `NotificationRequestBody` has been removed
- Field `AutomationRunbookReceivers` of struct `NotificationRequestBody` has been removed
- Field `EmailReceivers` of struct `NotificationRequestBody` has been removed
- Field `LogicAppReceivers` of struct `NotificationRequestBody` has been removed
- Field `ArmRoleReceivers` of struct `NotificationRequestBody` has been removed
- Field `WebhookReceivers` of struct `NotificationRequestBody` has been removed

### Features Added

- New struct `NotificationRequest`
- New field `Properties` in struct `NotificationRequestBody`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).