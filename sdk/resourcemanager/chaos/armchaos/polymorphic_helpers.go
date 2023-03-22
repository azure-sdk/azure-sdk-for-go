//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos

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
	switch m["type"] {
	case "continuous":
		b = &ContinuousAction{}
	case "delay":
		b = &DelayAction{}
	case "discrete":
		b = &DiscreteAction{}
	default:
		b = &Action{}
	}
	return b, json.Unmarshal(rawMsg, b)
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

func unmarshalFilterClassification(rawMsg json.RawMessage) (FilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FilterClassification
	switch m["type"] {
	case "Simple":
		b = &SimpleFilter{}
	default:
		b = &Filter{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSelectorClassification(rawMsg json.RawMessage) (SelectorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SelectorClassification
	switch m["type"] {
	case string(SelectorTypeList):
		b = &ListSelector{}
	case string(SelectorTypeQuery):
		b = &QuerySelector{}
	default:
		b = &Selector{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSelectorClassificationArray(rawMsg json.RawMessage) ([]SelectorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SelectorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSelectorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
