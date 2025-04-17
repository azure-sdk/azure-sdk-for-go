// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

import "encoding/json"

func unmarshalMeterDetailsClassification(rawMsg json.RawMessage) (MeterDetailsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MeterDetailsClassification
	switch m["billingType"] {
	case string(BillingTypePav2):
		b = &Pav2MeterDetails{}
	case string(BillingTypePurchase):
		b = &PurchaseMeterDetails{}
	default:
		b = &MeterDetails{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
