//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import "encoding/json"

func unmarshalResourceDetailsClassification(rawMsg json.RawMessage) (ResourceDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceDetailsClassification
	switch m["source"] {
	case string(SourceAzure):
		b = &AzureResourceDetails{}
	case string(SourceOnPremise):
		b = &OnPremiseResourceDetails{}
	case string(SourceOnPremiseSQL):
		b = &OnPremiseSQLResourceDetails{}
	default:
		b = &ResourceDetails{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
