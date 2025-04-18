// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

const (
	fullDateJSON = `"2006-01-02"`
	jsonFormat   = `"%04d-%02d-%02d"`
)

type dateType time.Time

func (t dateType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(jsonFormat, time.Time(t).Year(), time.Time(t).Month(), time.Time(t).Day())), nil
}

func (d *dateType) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(fullDateJSON, string(data))
	*d = (dateType)(t)
	return err
}

func populateDateType(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateType)(t)
}

func unpopulateDateType(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux dateType
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
