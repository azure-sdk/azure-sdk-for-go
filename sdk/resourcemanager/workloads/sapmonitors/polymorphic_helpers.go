//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package sapmonitors

import "encoding/json"

func unmarshalProviderSpecificPropertiesClassification(rawMsg json.RawMessage) (ProviderSpecificPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProviderSpecificPropertiesClassification
	switch m["providerType"] {
	case "Db2":
		b = &Db2ProviderInstanceProperties{}
	case "MsSqlServer":
		b = &MsSQLServerProviderInstanceProperties{}
	case "PrometheusHaCluster":
		b = &PrometheusHaClusterProviderInstanceProperties{}
	case "PrometheusOS":
		b = &PrometheusOsProviderInstanceProperties{}
	case "SapHana":
		b = &HanaDbProviderInstanceProperties{}
	case "SapNetWeaver":
		b = &SapNetWeaverProviderInstanceProperties{}
	default:
		b = &ProviderSpecificProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
