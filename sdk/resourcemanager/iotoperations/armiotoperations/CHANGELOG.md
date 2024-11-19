# Release History

## 0.2.0 (2024-11-19)
### Breaking Changes

- Function `*BrokerAuthenticationClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, BrokerAuthenticationResource, *BrokerAuthenticationClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, AuthenticationResource, *BrokerAuthenticationClientBeginCreateOrUpdateOptions)`
- Function `*BrokerAuthorizationClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, BrokerAuthorizationResource, *BrokerAuthorizationClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, AuthorizationResource, *BrokerAuthorizationClientBeginCreateOrUpdateOptions)`
- Function `*BrokerListenerClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, BrokerListenerResource, *BrokerListenerClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, ListenerResource, *BrokerListenerClientBeginCreateOrUpdateOptions)`
- Type of `DataflowEndpointResource.Properties` has been changed from `*DataflowEndpointProperties` to `DataflowEndpointPropertiesClassification`
- `ManagedServiceIdentityTypeSystemAndUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Enum `BrokerAuthenticationMethod` has been removed
- Struct `BrokerAuthenticationProperties` has been removed
- Struct `BrokerAuthenticationResource` has been removed
- Struct `BrokerAuthenticationResourceListResult` has been removed
- Struct `BrokerAuthenticatorMethodCustom` has been removed
- Struct `BrokerAuthenticatorMethodSat` has been removed
- Struct `BrokerAuthenticatorMethodX509` has been removed
- Struct `BrokerAuthenticatorMethods` has been removed
- Struct `BrokerAuthorizationProperties` has been removed
- Struct `BrokerAuthorizationResource` has been removed
- Struct `BrokerAuthorizationResourceListResult` has been removed
- Struct `BrokerListenerProperties` has been removed
- Struct `BrokerListenerResource` has been removed
- Struct `BrokerListenerResourceListResult` has been removed
- Struct `DataflowEndpointDataExplorer` has been removed
- Struct `DataflowEndpointDataLakeStorage` has been removed
- Struct `DataflowEndpointFabricOneLake` has been removed
- Struct `DataflowEndpointKafka` has been removed
- Struct `DataflowEndpointLocalStorage` has been removed
- Struct `DataflowEndpointMqtt` has been removed
- Field `BrokerAuthenticationResource` of struct `BrokerAuthenticationClientCreateOrUpdateResponse` has been removed
- Field `BrokerAuthenticationResource` of struct `BrokerAuthenticationClientGetResponse` has been removed
- Field `BrokerAuthenticationResourceListResult` of struct `BrokerAuthenticationClientListByResourceGroupResponse` has been removed
- Field `BrokerAuthorizationResource` of struct `BrokerAuthorizationClientCreateOrUpdateResponse` has been removed
- Field `BrokerAuthorizationResource` of struct `BrokerAuthorizationClientGetResponse` has been removed
- Field `BrokerAuthorizationResourceListResult` of struct `BrokerAuthorizationClientListByResourceGroupResponse` has been removed
- Field `BrokerListenerResource` of struct `BrokerListenerClientCreateOrUpdateResponse` has been removed
- Field `BrokerListenerResource` of struct `BrokerListenerClientGetResponse` has been removed
- Field `BrokerListenerResourceListResult` of struct `BrokerListenerClientListByResourceGroupResponse` has been removed
- Field `DataExplorerSettings`, `DataLakeStorageSettings`, `FabricOneLakeSettings`, `KafkaSettings`, `LocalStorageSettings`, `MqttSettings` of struct `DataflowEndpointProperties` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAssignedUserAssigned` added to enum type `ManagedServiceIdentityType`
- New enum type `AuthenticationMethod` with values `AuthenticationMethodCustom`, `AuthenticationMethodServiceAccountToken`, `AuthenticationMethodX509`
- New function `*DataExplorerEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*DataLakeStorageEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*DataflowEndpointProperties.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*FabricOneLakeEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*KafkaEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*LocalStorageEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New function `*MqttEndpoint.GetDataflowEndpointProperties() *DataflowEndpointProperties`
- New struct `AuthenticationProperties`
- New struct `AuthenticationResource`
- New struct `AuthenticationResourceListResult`
- New struct `AuthenticatorMethodCustom`
- New struct `AuthenticatorMethodSat`
- New struct `AuthenticatorMethodX509`
- New struct `AuthenticatorMethods`
- New struct `AuthorizationProperties`
- New struct `AuthorizationResource`
- New struct `AuthorizationResourceListResult`
- New struct `DataExplorerEndpoint`
- New struct `DataExplorerSettings`
- New struct `DataLakeStorageEndpoint`
- New struct `DataLakeStorageSettings`
- New struct `FabricOneLakeEndpoint`
- New struct `FabricOneLakeSettings`
- New struct `KafkaEndpoint`
- New struct `KafkaSettings`
- New struct `ListenerProperties`
- New struct `ListenerResource`
- New struct `ListenerResourceListResult`
- New struct `LocalStorageEndpoint`
- New struct `LocalStorageSettings`
- New struct `MqttEndpoint`
- New struct `MqttSettings`
- New anonymous field `AuthenticationResource` in struct `BrokerAuthenticationClientCreateOrUpdateResponse`
- New anonymous field `AuthenticationResource` in struct `BrokerAuthenticationClientGetResponse`
- New anonymous field `AuthenticationResourceListResult` in struct `BrokerAuthenticationClientListByResourceGroupResponse`
- New anonymous field `AuthorizationResource` in struct `BrokerAuthorizationClientCreateOrUpdateResponse`
- New anonymous field `AuthorizationResource` in struct `BrokerAuthorizationClientGetResponse`
- New anonymous field `AuthorizationResourceListResult` in struct `BrokerAuthorizationClientListByResourceGroupResponse`
- New anonymous field `ListenerResource` in struct `BrokerListenerClientCreateOrUpdateResponse`
- New anonymous field `ListenerResource` in struct `BrokerListenerClientGetResponse`
- New anonymous field `ListenerResourceListResult` in struct `BrokerListenerClientListByResourceGroupResponse`


## 0.1.0 (2024-10-24)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).