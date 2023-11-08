//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import "encoding/json"

func unmarshalApplicationGroupPolicyClassification(rawMsg json.RawMessage) (ApplicationGroupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ApplicationGroupPolicyClassification
	switch m["type"] {
	case string(ApplicationGroupPolicyTypeThrottlingPolicy):
		b = &ThrottlingPolicy{}
	default:
		b = &ApplicationGroupPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalApplicationGroupPolicyClassificationArray(rawMsg json.RawMessage) ([]ApplicationGroupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ApplicationGroupPolicyClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalApplicationGroupPolicyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
