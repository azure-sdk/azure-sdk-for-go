//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import "encoding/json"

func unmarshalActionClassification(rawMsg json.RawMessage) (ActionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ActionClassification
	switch m["actionType"] {
	case string(ActionTypeAddActionGroups):
		b = &AddActionGroups{}
	case string(ActionTypeCorrelateAlerts):
		b = &CorrelateAlerts{}
	case string(ActionTypeRemoveAllActionGroups):
		b = &RemoveAllActionGroups{}
	default:
		b = &Action{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalActionClassificationArray(rawMsg json.RawMessage) ([]ActionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAlertsMetaDataPropertiesClassification(rawMsg json.RawMessage) (AlertsMetaDataPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AlertsMetaDataPropertiesClassification
	switch m["metadataIdentifier"] {
	case string(MetadataIdentifierMonitorServiceList):
		b = &MonitorServiceList{}
	default:
		b = &AlertsMetaDataProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRecurrenceClassification(rawMsg json.RawMessage) (RecurrenceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecurrenceClassification
	switch m["recurrenceType"] {
	case string(RecurrenceTypeDaily):
		b = &DailyRecurrence{}
	case string(RecurrenceTypeMonthly):
		b = &MonthlyRecurrence{}
	case string(RecurrenceTypeWeekly):
		b = &WeeklyRecurrence{}
	default:
		b = &Recurrence{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRecurrenceClassificationArray(rawMsg json.RawMessage) ([]RecurrenceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RecurrenceClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRecurrenceClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
