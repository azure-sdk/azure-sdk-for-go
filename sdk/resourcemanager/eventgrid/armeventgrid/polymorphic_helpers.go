//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import "encoding/json"

func unmarshalAdvancedFilterClassification(rawMsg json.RawMessage) (AdvancedFilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AdvancedFilterClassification
	switch m["operatorType"] {
	case string(AdvancedFilterOperatorTypeBoolEquals):
		b = &BoolEqualsAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeIsNotNull):
		b = &IsNotNullAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeIsNullOrUndefined):
		b = &IsNullOrUndefinedAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberGreaterThan):
		b = &NumberGreaterThanAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberGreaterThanOrEquals):
		b = &NumberGreaterThanOrEqualsAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberIn):
		b = &NumberInAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberInRange):
		b = &NumberInRangeAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberLessThan):
		b = &NumberLessThanAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberLessThanOrEquals):
		b = &NumberLessThanOrEqualsAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberNotIn):
		b = &NumberNotInAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeNumberNotInRange):
		b = &NumberNotInRangeAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringBeginsWith):
		b = &StringBeginsWithAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringContains):
		b = &StringContainsAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringEndsWith):
		b = &StringEndsWithAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringIn):
		b = &StringInAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringNotBeginsWith):
		b = &StringNotBeginsWithAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringNotContains):
		b = &StringNotContainsAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringNotEndsWith):
		b = &StringNotEndsWithAdvancedFilter{}
	case string(AdvancedFilterOperatorTypeStringNotIn):
		b = &StringNotInAdvancedFilter{}
	default:
		b = &AdvancedFilter{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAdvancedFilterClassificationArray(rawMsg json.RawMessage) ([]AdvancedFilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AdvancedFilterClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAdvancedFilterClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDeadLetterDestinationClassification(rawMsg json.RawMessage) (DeadLetterDestinationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DeadLetterDestinationClassification
	switch m["endpointType"] {
	case string(DeadLetterEndPointTypeStorageBlob):
		b = &StorageBlobDeadLetterDestination{}
	default:
		b = &DeadLetterDestination{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDeliveryAttributeMappingClassification(rawMsg json.RawMessage) (DeliveryAttributeMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DeliveryAttributeMappingClassification
	switch m["type"] {
	case string(DeliveryAttributeMappingTypeDynamic):
		b = &DynamicDeliveryAttributeMapping{}
	case string(DeliveryAttributeMappingTypeStatic):
		b = &StaticDeliveryAttributeMapping{}
	default:
		b = &DeliveryAttributeMapping{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDeliveryAttributeMappingClassificationArray(rawMsg json.RawMessage) ([]DeliveryAttributeMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DeliveryAttributeMappingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDeliveryAttributeMappingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalEventSubscriptionDestinationClassification(rawMsg json.RawMessage) (EventSubscriptionDestinationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventSubscriptionDestinationClassification
	switch m["endpointType"] {
	case string(EndpointTypeAzureFunction):
		b = &AzureFunctionEventSubscriptionDestination{}
	case string(EndpointTypeEventHub):
		b = &EventHubEventSubscriptionDestination{}
	case string(EndpointTypeHybridConnection):
		b = &HybridConnectionEventSubscriptionDestination{}
	case string(EndpointTypePartnerDestination):
		b = &PartnerEventSubscriptionDestination{}
	case string(EndpointTypeServiceBusQueue):
		b = &ServiceBusQueueEventSubscriptionDestination{}
	case string(EndpointTypeServiceBusTopic):
		b = &ServiceBusTopicEventSubscriptionDestination{}
	case string(EndpointTypeStorageQueue):
		b = &StorageQueueEventSubscriptionDestination{}
	case string(EndpointTypeWebHook):
		b = &WebHookEventSubscriptionDestination{}
	default:
		b = &EventSubscriptionDestination{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFilterClassification(rawMsg json.RawMessage) (FilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FilterClassification
	switch m["operatorType"] {
	case string(FilterOperatorTypeBoolEquals):
		b = &BoolEqualsFilter{}
	case string(FilterOperatorTypeIsNotNull):
		b = &IsNotNullFilter{}
	case string(FilterOperatorTypeIsNullOrUndefined):
		b = &IsNullOrUndefinedFilter{}
	case string(FilterOperatorTypeNumberGreaterThan):
		b = &NumberGreaterThanFilter{}
	case string(FilterOperatorTypeNumberGreaterThanOrEquals):
		b = &NumberGreaterThanOrEqualsFilter{}
	case string(FilterOperatorTypeNumberIn):
		b = &NumberInFilter{}
	case string(FilterOperatorTypeNumberInRange):
		b = &NumberInRangeFilter{}
	case string(FilterOperatorTypeNumberLessThan):
		b = &NumberLessThanFilter{}
	case string(FilterOperatorTypeNumberLessThanOrEquals):
		b = &NumberLessThanOrEqualsFilter{}
	case string(FilterOperatorTypeNumberNotIn):
		b = &NumberNotInFilter{}
	case string(FilterOperatorTypeNumberNotInRange):
		b = &NumberNotInRangeFilter{}
	case string(FilterOperatorTypeStringBeginsWith):
		b = &StringBeginsWithFilter{}
	case string(FilterOperatorTypeStringContains):
		b = &StringContainsFilter{}
	case string(FilterOperatorTypeStringEndsWith):
		b = &StringEndsWithFilter{}
	case string(FilterOperatorTypeStringIn):
		b = &StringInFilter{}
	case string(FilterOperatorTypeStringNotBeginsWith):
		b = &StringNotBeginsWithFilter{}
	case string(FilterOperatorTypeStringNotContains):
		b = &StringNotContainsFilter{}
	case string(FilterOperatorTypeStringNotEndsWith):
		b = &StringNotEndsWithFilter{}
	case string(FilterOperatorTypeStringNotIn):
		b = &StringNotInFilter{}
	default:
		b = &Filter{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFilterClassificationArray(rawMsg json.RawMessage) ([]FilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FilterClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFilterClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalInputSchemaMappingClassification(rawMsg json.RawMessage) (InputSchemaMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InputSchemaMappingClassification
	switch m["inputSchemaMappingType"] {
	case string(InputSchemaMappingTypeJSON):
		b = &JSONInputSchemaMapping{}
	default:
		b = &InputSchemaMapping{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPartnerClientAuthenticationClassification(rawMsg json.RawMessage) (PartnerClientAuthenticationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PartnerClientAuthenticationClassification
	switch m["clientAuthenticationType"] {
	case string(PartnerClientAuthenticationTypeAzureAD):
		b = &AzureADPartnerClientAuthentication{}
	default:
		b = &PartnerClientAuthentication{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPartnerDestinationInfoClassification(rawMsg json.RawMessage) (PartnerDestinationInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PartnerDestinationInfoClassification
	switch m["endpointType"] {
	case string(PartnerEndpointTypeWebHook):
		b = &WebhookPartnerDestinationInfo{}
	default:
		b = &PartnerDestinationInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPartnerUpdateDestinationInfoClassification(rawMsg json.RawMessage) (PartnerUpdateDestinationInfoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PartnerUpdateDestinationInfoClassification
	switch m["endpointType"] {
	case string(PartnerEndpointTypeWebHook):
		b = &WebhookUpdatePartnerDestinationInfo{}
	default:
		b = &PartnerUpdateDestinationInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
