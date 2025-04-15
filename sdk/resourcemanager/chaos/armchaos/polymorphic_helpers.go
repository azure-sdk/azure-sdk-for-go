// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armchaos

import "encoding/json"

func unmarshalExperimentActionClassification(rawMsg json.RawMessage) (ExperimentActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ExperimentActionClassification
	switch m["type"] {
	case string(ExperimentActionTypeContinuous):
		b = &ContinuousAction{}
	case string(ExperimentActionTypeDelay):
		b = &DelayAction{}
	case string(ExperimentActionTypeDiscrete):
		b = &DiscreteAction{}
	default:
		b = &ExperimentAction{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalExperimentActionClassificationArray(rawMsg json.RawMessage) ([]ExperimentActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ExperimentActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalExperimentActionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTargetFilterClassification(rawMsg json.RawMessage) (TargetFilterClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetFilterClassification
	switch m["type"] {
	case string(FilterTypeSimple):
		b = &TargetSimpleFilter{}
	default:
		b = &TargetFilter{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTargetSelectorClassification(rawMsg json.RawMessage) (TargetSelectorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetSelectorClassification
	switch m["type"] {
	case string(SelectorTypeList):
		b = &TargetListSelector{}
	case string(SelectorTypeQuery):
		b = &TargetQuerySelector{}
	default:
		b = &TargetSelector{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTargetSelectorClassificationArray(rawMsg json.RawMessage) ([]TargetSelectorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TargetSelectorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTargetSelectorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
