//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ArmDisasterRecoveryListResult.
func (a ArmDisasterRecoveryListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CorrelationFilter.
func (c CorrelationFilter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", c.ContentType)
	populate(objectMap, "correlationId", c.CorrelationID)
	populate(objectMap, "label", c.Label)
	populate(objectMap, "messageId", c.MessageID)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "replyTo", c.ReplyTo)
	populate(objectMap, "replyToSessionId", c.ReplyToSessionID)
	populate(objectMap, "requiresPreprocessing", c.RequiresPreprocessing)
	populate(objectMap, "sessionId", c.SessionID)
	populate(objectMap, "to", c.To)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Encryption.
func (e Encryption) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "keySource", e.KeySource)
	populate(objectMap, "keyVaultProperties", e.KeyVaultProperties)
	populate(objectMap, "requireInfrastructureEncryption", e.RequireInfrastructureEncryption)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseError.
func (e ErrorResponseError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MigrationConfigListResult.
func (m MigrationConfigListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSetListResult.
func (n NetworkRuleSetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", n.NextLink)
	populate(objectMap, "value", n.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSetProperties.
func (n NetworkRuleSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultAction", n.DefaultAction)
	populate(objectMap, "ipRules", n.IPRules)
	populate(objectMap, "publicNetworkAccess", n.PublicNetworkAccess)
	populate(objectMap, "trustedServiceAccessEnabled", n.TrustedServiceAccessEnabled)
	populate(objectMap, "virtualNetworkRules", n.VirtualNetworkRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourcesListResult.
func (p PrivateLinkResourcesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceNamespacePatch.
func (r ResourceNamespacePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RuleListResult.
func (r RuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBAuthorizationRuleListResult.
func (s SBAuthorizationRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBAuthorizationRuleProperties.
func (s SBAuthorizationRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "rights", s.Rights)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBNamespace.
func (s SBNamespace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBNamespaceListResult.
func (s SBNamespaceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBNamespaceProperties.
func (s SBNamespaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alternateName", s.AlternateName)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "disableLocalAuth", s.DisableLocalAuth)
	populate(objectMap, "encryption", s.Encryption)
	populate(objectMap, "metricId", s.MetricID)
	populate(objectMap, "minimumTlsVersion", s.MinimumTLSVersion)
	populate(objectMap, "privateEndpointConnections", s.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "serviceBusEndpoint", s.ServiceBusEndpoint)
	populate(objectMap, "status", s.Status)
	populateTimeRFC3339(objectMap, "updatedAt", s.UpdatedAt)
	populate(objectMap, "zoneRedundant", s.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SBNamespaceProperties.
func (s *SBNamespaceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alternateName":
			err = unpopulate(val, &s.AlternateName)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, &s.DisableLocalAuth)
			delete(rawMsg, key)
		case "encryption":
			err = unpopulate(val, &s.Encryption)
			delete(rawMsg, key)
		case "metricId":
			err = unpopulate(val, &s.MetricID)
			delete(rawMsg, key)
		case "minimumTlsVersion":
			err = unpopulate(val, &s.MinimumTLSVersion)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &s.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "serviceBusEndpoint":
			err = unpopulate(val, &s.ServiceBusEndpoint)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, &s.UpdatedAt)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, &s.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SBNamespaceUpdateParameters.
func (s SBNamespaceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBNamespaceUpdateProperties.
func (s SBNamespaceUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alternateName", s.AlternateName)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "disableLocalAuth", s.DisableLocalAuth)
	populate(objectMap, "encryption", s.Encryption)
	populate(objectMap, "metricId", s.MetricID)
	populate(objectMap, "privateEndpointConnections", s.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "serviceBusEndpoint", s.ServiceBusEndpoint)
	populate(objectMap, "status", s.Status)
	populateTimeRFC3339(objectMap, "updatedAt", s.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SBNamespaceUpdateProperties.
func (s *SBNamespaceUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alternateName":
			err = unpopulate(val, &s.AlternateName)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, &s.DisableLocalAuth)
			delete(rawMsg, key)
		case "encryption":
			err = unpopulate(val, &s.Encryption)
			delete(rawMsg, key)
		case "metricId":
			err = unpopulate(val, &s.MetricID)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &s.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "serviceBusEndpoint":
			err = unpopulate(val, &s.ServiceBusEndpoint)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, &s.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SBQueueListResult.
func (s SBQueueListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBQueueProperties.
func (s SBQueueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "accessedAt", s.AccessedAt)
	populate(objectMap, "autoDeleteOnIdle", s.AutoDeleteOnIdle)
	populate(objectMap, "countDetails", s.CountDetails)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "deadLetteringOnMessageExpiration", s.DeadLetteringOnMessageExpiration)
	populate(objectMap, "defaultMessageTimeToLive", s.DefaultMessageTimeToLive)
	populate(objectMap, "duplicateDetectionHistoryTimeWindow", s.DuplicateDetectionHistoryTimeWindow)
	populate(objectMap, "enableBatchedOperations", s.EnableBatchedOperations)
	populate(objectMap, "enableExpress", s.EnableExpress)
	populate(objectMap, "enablePartitioning", s.EnablePartitioning)
	populate(objectMap, "forwardDeadLetteredMessagesTo", s.ForwardDeadLetteredMessagesTo)
	populate(objectMap, "forwardTo", s.ForwardTo)
	populate(objectMap, "lockDuration", s.LockDuration)
	populate(objectMap, "maxDeliveryCount", s.MaxDeliveryCount)
	populate(objectMap, "maxMessageSizeInKilobytes", s.MaxMessageSizeInKilobytes)
	populate(objectMap, "maxSizeInMegabytes", s.MaxSizeInMegabytes)
	populate(objectMap, "messageCount", s.MessageCount)
	populate(objectMap, "requiresDuplicateDetection", s.RequiresDuplicateDetection)
	populate(objectMap, "requiresSession", s.RequiresSession)
	populate(objectMap, "sizeInBytes", s.SizeInBytes)
	populate(objectMap, "status", s.Status)
	populateTimeRFC3339(objectMap, "updatedAt", s.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SBQueueProperties.
func (s *SBQueueProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessedAt":
			err = unpopulateTimeRFC3339(val, &s.AccessedAt)
			delete(rawMsg, key)
		case "autoDeleteOnIdle":
			err = unpopulate(val, &s.AutoDeleteOnIdle)
			delete(rawMsg, key)
		case "countDetails":
			err = unpopulate(val, &s.CountDetails)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "deadLetteringOnMessageExpiration":
			err = unpopulate(val, &s.DeadLetteringOnMessageExpiration)
			delete(rawMsg, key)
		case "defaultMessageTimeToLive":
			err = unpopulate(val, &s.DefaultMessageTimeToLive)
			delete(rawMsg, key)
		case "duplicateDetectionHistoryTimeWindow":
			err = unpopulate(val, &s.DuplicateDetectionHistoryTimeWindow)
			delete(rawMsg, key)
		case "enableBatchedOperations":
			err = unpopulate(val, &s.EnableBatchedOperations)
			delete(rawMsg, key)
		case "enableExpress":
			err = unpopulate(val, &s.EnableExpress)
			delete(rawMsg, key)
		case "enablePartitioning":
			err = unpopulate(val, &s.EnablePartitioning)
			delete(rawMsg, key)
		case "forwardDeadLetteredMessagesTo":
			err = unpopulate(val, &s.ForwardDeadLetteredMessagesTo)
			delete(rawMsg, key)
		case "forwardTo":
			err = unpopulate(val, &s.ForwardTo)
			delete(rawMsg, key)
		case "lockDuration":
			err = unpopulate(val, &s.LockDuration)
			delete(rawMsg, key)
		case "maxDeliveryCount":
			err = unpopulate(val, &s.MaxDeliveryCount)
			delete(rawMsg, key)
		case "maxMessageSizeInKilobytes":
			err = unpopulate(val, &s.MaxMessageSizeInKilobytes)
			delete(rawMsg, key)
		case "maxSizeInMegabytes":
			err = unpopulate(val, &s.MaxSizeInMegabytes)
			delete(rawMsg, key)
		case "messageCount":
			err = unpopulate(val, &s.MessageCount)
			delete(rawMsg, key)
		case "requiresDuplicateDetection":
			err = unpopulate(val, &s.RequiresDuplicateDetection)
			delete(rawMsg, key)
		case "requiresSession":
			err = unpopulate(val, &s.RequiresSession)
			delete(rawMsg, key)
		case "sizeInBytes":
			err = unpopulate(val, &s.SizeInBytes)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, &s.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SBSubscriptionListResult.
func (s SBSubscriptionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBSubscriptionProperties.
func (s SBSubscriptionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "accessedAt", s.AccessedAt)
	populate(objectMap, "autoDeleteOnIdle", s.AutoDeleteOnIdle)
	populate(objectMap, "clientAffineProperties", s.ClientAffineProperties)
	populate(objectMap, "countDetails", s.CountDetails)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "deadLetteringOnFilterEvaluationExceptions", s.DeadLetteringOnFilterEvaluationExceptions)
	populate(objectMap, "deadLetteringOnMessageExpiration", s.DeadLetteringOnMessageExpiration)
	populate(objectMap, "defaultMessageTimeToLive", s.DefaultMessageTimeToLive)
	populate(objectMap, "duplicateDetectionHistoryTimeWindow", s.DuplicateDetectionHistoryTimeWindow)
	populate(objectMap, "enableBatchedOperations", s.EnableBatchedOperations)
	populate(objectMap, "forwardDeadLetteredMessagesTo", s.ForwardDeadLetteredMessagesTo)
	populate(objectMap, "forwardTo", s.ForwardTo)
	populate(objectMap, "isClientAffine", s.IsClientAffine)
	populate(objectMap, "lockDuration", s.LockDuration)
	populate(objectMap, "maxDeliveryCount", s.MaxDeliveryCount)
	populate(objectMap, "messageCount", s.MessageCount)
	populate(objectMap, "requiresSession", s.RequiresSession)
	populate(objectMap, "status", s.Status)
	populateTimeRFC3339(objectMap, "updatedAt", s.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SBSubscriptionProperties.
func (s *SBSubscriptionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessedAt":
			err = unpopulateTimeRFC3339(val, &s.AccessedAt)
			delete(rawMsg, key)
		case "autoDeleteOnIdle":
			err = unpopulate(val, &s.AutoDeleteOnIdle)
			delete(rawMsg, key)
		case "clientAffineProperties":
			err = unpopulate(val, &s.ClientAffineProperties)
			delete(rawMsg, key)
		case "countDetails":
			err = unpopulate(val, &s.CountDetails)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "deadLetteringOnFilterEvaluationExceptions":
			err = unpopulate(val, &s.DeadLetteringOnFilterEvaluationExceptions)
			delete(rawMsg, key)
		case "deadLetteringOnMessageExpiration":
			err = unpopulate(val, &s.DeadLetteringOnMessageExpiration)
			delete(rawMsg, key)
		case "defaultMessageTimeToLive":
			err = unpopulate(val, &s.DefaultMessageTimeToLive)
			delete(rawMsg, key)
		case "duplicateDetectionHistoryTimeWindow":
			err = unpopulate(val, &s.DuplicateDetectionHistoryTimeWindow)
			delete(rawMsg, key)
		case "enableBatchedOperations":
			err = unpopulate(val, &s.EnableBatchedOperations)
			delete(rawMsg, key)
		case "forwardDeadLetteredMessagesTo":
			err = unpopulate(val, &s.ForwardDeadLetteredMessagesTo)
			delete(rawMsg, key)
		case "forwardTo":
			err = unpopulate(val, &s.ForwardTo)
			delete(rawMsg, key)
		case "isClientAffine":
			err = unpopulate(val, &s.IsClientAffine)
			delete(rawMsg, key)
		case "lockDuration":
			err = unpopulate(val, &s.LockDuration)
			delete(rawMsg, key)
		case "maxDeliveryCount":
			err = unpopulate(val, &s.MaxDeliveryCount)
			delete(rawMsg, key)
		case "messageCount":
			err = unpopulate(val, &s.MessageCount)
			delete(rawMsg, key)
		case "requiresSession":
			err = unpopulate(val, &s.RequiresSession)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, &s.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SBTopicListResult.
func (s SBTopicListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SBTopicProperties.
func (s SBTopicProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "accessedAt", s.AccessedAt)
	populate(objectMap, "autoDeleteOnIdle", s.AutoDeleteOnIdle)
	populate(objectMap, "countDetails", s.CountDetails)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "defaultMessageTimeToLive", s.DefaultMessageTimeToLive)
	populate(objectMap, "duplicateDetectionHistoryTimeWindow", s.DuplicateDetectionHistoryTimeWindow)
	populate(objectMap, "enableBatchedOperations", s.EnableBatchedOperations)
	populate(objectMap, "enableExpress", s.EnableExpress)
	populate(objectMap, "enablePartitioning", s.EnablePartitioning)
	populate(objectMap, "maxMessageSizeInKilobytes", s.MaxMessageSizeInKilobytes)
	populate(objectMap, "maxSizeInMegabytes", s.MaxSizeInMegabytes)
	populate(objectMap, "requiresDuplicateDetection", s.RequiresDuplicateDetection)
	populate(objectMap, "sizeInBytes", s.SizeInBytes)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "subscriptionCount", s.SubscriptionCount)
	populate(objectMap, "supportOrdering", s.SupportOrdering)
	populateTimeRFC3339(objectMap, "updatedAt", s.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SBTopicProperties.
func (s *SBTopicProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessedAt":
			err = unpopulateTimeRFC3339(val, &s.AccessedAt)
			delete(rawMsg, key)
		case "autoDeleteOnIdle":
			err = unpopulate(val, &s.AutoDeleteOnIdle)
			delete(rawMsg, key)
		case "countDetails":
			err = unpopulate(val, &s.CountDetails)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "defaultMessageTimeToLive":
			err = unpopulate(val, &s.DefaultMessageTimeToLive)
			delete(rawMsg, key)
		case "duplicateDetectionHistoryTimeWindow":
			err = unpopulate(val, &s.DuplicateDetectionHistoryTimeWindow)
			delete(rawMsg, key)
		case "enableBatchedOperations":
			err = unpopulate(val, &s.EnableBatchedOperations)
			delete(rawMsg, key)
		case "enableExpress":
			err = unpopulate(val, &s.EnableExpress)
			delete(rawMsg, key)
		case "enablePartitioning":
			err = unpopulate(val, &s.EnablePartitioning)
			delete(rawMsg, key)
		case "maxMessageSizeInKilobytes":
			err = unpopulate(val, &s.MaxMessageSizeInKilobytes)
			delete(rawMsg, key)
		case "maxSizeInMegabytes":
			err = unpopulate(val, &s.MaxSizeInMegabytes)
			delete(rawMsg, key)
		case "requiresDuplicateDetection":
			err = unpopulate(val, &s.RequiresDuplicateDetection)
			delete(rawMsg, key)
		case "sizeInBytes":
			err = unpopulate(val, &s.SizeInBytes)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &s.Status)
			delete(rawMsg, key)
		case "subscriptionCount":
			err = unpopulate(val, &s.SubscriptionCount)
			delete(rawMsg, key)
		case "supportOrdering":
			err = unpopulate(val, &s.SupportOrdering)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, &s.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
