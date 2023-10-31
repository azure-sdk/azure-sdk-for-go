//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

// AdvancedFilterClassification provides polymorphic access to related types.
// Call the interface's GetAdvancedFilter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AdvancedFilter, *BoolEqualsAdvancedFilter, *IsNotNullAdvancedFilter, *IsNullOrUndefinedAdvancedFilter, *NumberGreaterThanAdvancedFilter,
// - *NumberGreaterThanOrEqualsAdvancedFilter, *NumberInAdvancedFilter, *NumberInRangeAdvancedFilter, *NumberLessThanAdvancedFilter,
// - *NumberLessThanOrEqualsAdvancedFilter, *NumberNotInAdvancedFilter, *NumberNotInRangeAdvancedFilter, *StringBeginsWithAdvancedFilter,
// - *StringContainsAdvancedFilter, *StringEndsWithAdvancedFilter, *StringInAdvancedFilter, *StringNotBeginsWithAdvancedFilter,
// - *StringNotContainsAdvancedFilter, *StringNotEndsWithAdvancedFilter, *StringNotInAdvancedFilter
type AdvancedFilterClassification interface {
	// GetAdvancedFilter returns the AdvancedFilter content of the underlying type.
	GetAdvancedFilter() *AdvancedFilter
}

// DeadLetterDestinationClassification provides polymorphic access to related types.
// Call the interface's GetDeadLetterDestination() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DeadLetterDestination, *StorageBlobDeadLetterDestination
type DeadLetterDestinationClassification interface {
	// GetDeadLetterDestination returns the DeadLetterDestination content of the underlying type.
	GetDeadLetterDestination() *DeadLetterDestination
}

// DeliveryAttributeMappingClassification provides polymorphic access to related types.
// Call the interface's GetDeliveryAttributeMapping() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DeliveryAttributeMapping, *DynamicDeliveryAttributeMapping, *StaticDeliveryAttributeMapping
type DeliveryAttributeMappingClassification interface {
	// GetDeliveryAttributeMapping returns the DeliveryAttributeMapping content of the underlying type.
	GetDeliveryAttributeMapping() *DeliveryAttributeMapping
}

// EventSubscriptionDestinationClassification provides polymorphic access to related types.
// Call the interface's GetEventSubscriptionDestination() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureFunctionEventSubscriptionDestination, *EventHubEventSubscriptionDestination, *EventSubscriptionDestination, *HybridConnectionEventSubscriptionDestination,
// - *MonitorAlertEventSubscriptionDestination, *NamespaceTopicEventSubscriptionDestination, *PartnerEventSubscriptionDestination,
// - *ServiceBusQueueEventSubscriptionDestination, *ServiceBusTopicEventSubscriptionDestination, *StorageQueueEventSubscriptionDestination,
// - *WebHookEventSubscriptionDestination
type EventSubscriptionDestinationClassification interface {
	// GetEventSubscriptionDestination returns the EventSubscriptionDestination content of the underlying type.
	GetEventSubscriptionDestination() *EventSubscriptionDestination
}

// FilterClassification provides polymorphic access to related types.
// Call the interface's GetFilter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BoolEqualsFilter, *Filter, *IsNotNullFilter, *IsNullOrUndefinedFilter, *NumberGreaterThanFilter, *NumberGreaterThanOrEqualsFilter,
// - *NumberInFilter, *NumberInRangeFilter, *NumberLessThanFilter, *NumberLessThanOrEqualsFilter, *NumberNotInFilter, *NumberNotInRangeFilter,
// - *StringBeginsWithFilter, *StringContainsFilter, *StringEndsWithFilter, *StringInFilter, *StringNotBeginsWithFilter, *StringNotContainsFilter,
// - *StringNotEndsWithFilter, *StringNotInFilter
type FilterClassification interface {
	// GetFilter returns the Filter content of the underlying type.
	GetFilter() *Filter
}

// InputSchemaMappingClassification provides polymorphic access to related types.
// Call the interface's GetInputSchemaMapping() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *InputSchemaMapping, *JSONInputSchemaMapping
type InputSchemaMappingClassification interface {
	// GetInputSchemaMapping returns the InputSchemaMapping content of the underlying type.
	GetInputSchemaMapping() *InputSchemaMapping
}

// PartnerClientAuthenticationClassification provides polymorphic access to related types.
// Call the interface's GetPartnerClientAuthentication() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureADPartnerClientAuthentication, *PartnerClientAuthentication
type PartnerClientAuthenticationClassification interface {
	// GetPartnerClientAuthentication returns the PartnerClientAuthentication content of the underlying type.
	GetPartnerClientAuthentication() *PartnerClientAuthentication
}

// PartnerDestinationInfoClassification provides polymorphic access to related types.
// Call the interface's GetPartnerDestinationInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PartnerDestinationInfo, *WebhookPartnerDestinationInfo
type PartnerDestinationInfoClassification interface {
	// GetPartnerDestinationInfo returns the PartnerDestinationInfo content of the underlying type.
	GetPartnerDestinationInfo() *PartnerDestinationInfo
}

// PartnerUpdateDestinationInfoClassification provides polymorphic access to related types.
// Call the interface's GetPartnerUpdateDestinationInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PartnerUpdateDestinationInfo, *WebhookUpdatePartnerDestinationInfo
type PartnerUpdateDestinationInfoClassification interface {
	// GetPartnerUpdateDestinationInfo returns the PartnerUpdateDestinationInfo content of the underlying type.
	GetPartnerUpdateDestinationInfo() *PartnerUpdateDestinationInfo
}
