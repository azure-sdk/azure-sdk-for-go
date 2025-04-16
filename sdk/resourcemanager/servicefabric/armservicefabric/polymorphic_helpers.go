// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

import "encoding/json"

func unmarshalPartitionSchemeDescriptionClassification(rawMsg json.RawMessage) (PartitionSchemeDescriptionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PartitionSchemeDescriptionClassification
	switch m["partitionScheme"] {
	case string(PartitionSchemeNamed):
		b = &NamedPartitionSchemeDescription{}
	case string(PartitionSchemeSingleton):
		b = &SingletonPartitionSchemeDescription{}
	case string(PartitionSchemeUniformInt64Range):
		b = &UniformInt64RangePartitionSchemeDescription{}
	default:
		b = &PartitionSchemeDescription{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalServicePlacementPolicyDescriptionClassification(rawMsg json.RawMessage) (ServicePlacementPolicyDescriptionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServicePlacementPolicyDescriptionClassification
	switch m["type"] {
	default:
		b = &ServicePlacementPolicyDescription{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalServicePlacementPolicyDescriptionClassificationArray(rawMsg json.RawMessage) ([]ServicePlacementPolicyDescriptionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ServicePlacementPolicyDescriptionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalServicePlacementPolicyDescriptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalServiceResourcePropertiesClassification(rawMsg json.RawMessage) (ServiceResourcePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourcePropertiesClassification
	switch m["serviceKind"] {
	case string(ServiceKindStateful):
		b = &StatefulServiceProperties{}
	case string(ServiceKindStateless):
		b = &StatelessServiceProperties{}
	default:
		b = &ServiceResourceProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalServiceResourceUpdatePropertiesClassification(rawMsg json.RawMessage) (ServiceResourceUpdatePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourceUpdatePropertiesClassification
	switch m["serviceKind"] {
	case string(ServiceKindStateful):
		b = &StatefulServiceUpdateProperties{}
	case string(ServiceKindStateless):
		b = &StatelessServiceUpdateProperties{}
	default:
		b = &ServiceResourceUpdateProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
