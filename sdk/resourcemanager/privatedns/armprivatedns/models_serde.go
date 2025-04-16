// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ARecord.
func (a ARecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "ipv4Address", a.IPv4Address)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ARecord.
func (a *ARecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "ipv4Address":
			err = unpopulate(val, "IPv4Address", &a.IPv4Address)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AaaaRecord.
func (a AaaaRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "ipv6Address", a.IPv6Address)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AaaaRecord.
func (a *AaaaRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "ipv6Address":
			err = unpopulate(val, "IPv6Address", &a.IPv6Address)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CnameRecord.
func (c CnameRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "cname", c.Cname)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CnameRecord.
func (c *CnameRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cname":
			err = unpopulate(val, "Cname", &c.Cname)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MxRecord.
func (m MxRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "exchange", m.Exchange)
	populate(objectMap, "preference", m.Preference)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MxRecord.
func (m *MxRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "exchange":
			err = unpopulate(val, "Exchange", &m.Exchange)
			delete(rawMsg, key)
		case "preference":
			err = unpopulate(val, "Preference", &m.Preference)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateZone.
func (p PrivateZone) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateZone.
func (p *PrivateZone) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "etag":
			err = unpopulate(val, "Etag", &p.Etag)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &p.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &p.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateZoneListResult.
func (p PrivateZoneListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateZoneListResult.
func (p *PrivateZoneListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &p.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &p.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateZoneProperties.
func (p PrivateZoneProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "internalId", p.InternalID)
	populate(objectMap, "maxNumberOfRecordSets", p.MaxNumberOfRecordSets)
	populate(objectMap, "maxNumberOfVirtualNetworkLinks", p.MaxNumberOfVirtualNetworkLinks)
	populate(objectMap, "maxNumberOfVirtualNetworkLinksWithRegistration", p.MaxNumberOfVirtualNetworkLinksWithRegistration)
	populate(objectMap, "numberOfRecordSets", p.NumberOfRecordSets)
	populate(objectMap, "numberOfVirtualNetworkLinks", p.NumberOfVirtualNetworkLinks)
	populate(objectMap, "numberOfVirtualNetworkLinksWithRegistration", p.NumberOfVirtualNetworkLinksWithRegistration)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateZoneProperties.
func (p *PrivateZoneProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "internalId":
			err = unpopulate(val, "InternalID", &p.InternalID)
			delete(rawMsg, key)
		case "maxNumberOfRecordSets":
			err = unpopulate(val, "MaxNumberOfRecordSets", &p.MaxNumberOfRecordSets)
			delete(rawMsg, key)
		case "maxNumberOfVirtualNetworkLinks":
			err = unpopulate(val, "MaxNumberOfVirtualNetworkLinks", &p.MaxNumberOfVirtualNetworkLinks)
			delete(rawMsg, key)
		case "maxNumberOfVirtualNetworkLinksWithRegistration":
			err = unpopulate(val, "MaxNumberOfVirtualNetworkLinksWithRegistration", &p.MaxNumberOfVirtualNetworkLinksWithRegistration)
			delete(rawMsg, key)
		case "numberOfRecordSets":
			err = unpopulate(val, "NumberOfRecordSets", &p.NumberOfRecordSets)
			delete(rawMsg, key)
		case "numberOfVirtualNetworkLinks":
			err = unpopulate(val, "NumberOfVirtualNetworkLinks", &p.NumberOfVirtualNetworkLinks)
			delete(rawMsg, key)
		case "numberOfVirtualNetworkLinksWithRegistration":
			err = unpopulate(val, "NumberOfVirtualNetworkLinksWithRegistration", &p.NumberOfVirtualNetworkLinksWithRegistration)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProxyResource.
func (p ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProxyResource.
func (p *ProxyResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PtrRecord.
func (p PtrRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "ptrdname", p.Ptrdname)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PtrRecord.
func (p *PtrRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "ptrdname":
			err = unpopulate(val, "Ptrdname", &p.Ptrdname)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecordSet.
func (r RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecordSet.
func (r *RecordSet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "etag":
			err = unpopulate(val, "Etag", &r.Etag)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecordSetListResult.
func (r RecordSetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecordSetListResult.
func (r *RecordSetListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecordSetProperties.
func (r RecordSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aRecords", r.ARecords)
	populate(objectMap, "aaaaRecords", r.AaaaRecords)
	populate(objectMap, "cnameRecord", r.CnameRecord)
	populate(objectMap, "fqdn", r.Fqdn)
	populate(objectMap, "isAutoRegistered", r.IsAutoRegistered)
	populate(objectMap, "metadata", r.Metadata)
	populate(objectMap, "mxRecords", r.MxRecords)
	populate(objectMap, "ptrRecords", r.PtrRecords)
	populate(objectMap, "soaRecord", r.SoaRecord)
	populate(objectMap, "srvRecords", r.SrvRecords)
	populate(objectMap, "ttl", r.TTL)
	populate(objectMap, "txtRecords", r.TxtRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecordSetProperties.
func (r *RecordSetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aRecords":
			err = unpopulate(val, "ARecords", &r.ARecords)
			delete(rawMsg, key)
		case "aaaaRecords":
			err = unpopulate(val, "AaaaRecords", &r.AaaaRecords)
			delete(rawMsg, key)
		case "cnameRecord":
			err = unpopulate(val, "CnameRecord", &r.CnameRecord)
			delete(rawMsg, key)
		case "fqdn":
			err = unpopulate(val, "Fqdn", &r.Fqdn)
			delete(rawMsg, key)
		case "isAutoRegistered":
			err = unpopulate(val, "IsAutoRegistered", &r.IsAutoRegistered)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, "Metadata", &r.Metadata)
			delete(rawMsg, key)
		case "mxRecords":
			err = unpopulate(val, "MxRecords", &r.MxRecords)
			delete(rawMsg, key)
		case "ptrRecords":
			err = unpopulate(val, "PtrRecords", &r.PtrRecords)
			delete(rawMsg, key)
		case "soaRecord":
			err = unpopulate(val, "SoaRecord", &r.SoaRecord)
			delete(rawMsg, key)
		case "srvRecords":
			err = unpopulate(val, "SrvRecords", &r.SrvRecords)
			delete(rawMsg, key)
		case "ttl":
			err = unpopulate(val, "TTL", &r.TTL)
			delete(rawMsg, key)
		case "txtRecords":
			err = unpopulate(val, "TxtRecords", &r.TxtRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SoaRecord.
func (s SoaRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "email", s.Email)
	populate(objectMap, "expireTime", s.ExpireTime)
	populate(objectMap, "host", s.Host)
	populate(objectMap, "minimumTtl", s.MinimumTTL)
	populate(objectMap, "refreshTime", s.RefreshTime)
	populate(objectMap, "retryTime", s.RetryTime)
	populate(objectMap, "serialNumber", s.SerialNumber)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SoaRecord.
func (s *SoaRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "email":
			err = unpopulate(val, "Email", &s.Email)
			delete(rawMsg, key)
		case "expireTime":
			err = unpopulate(val, "ExpireTime", &s.ExpireTime)
			delete(rawMsg, key)
		case "host":
			err = unpopulate(val, "Host", &s.Host)
			delete(rawMsg, key)
		case "minimumTtl":
			err = unpopulate(val, "MinimumTTL", &s.MinimumTTL)
			delete(rawMsg, key)
		case "refreshTime":
			err = unpopulate(val, "RefreshTime", &s.RefreshTime)
			delete(rawMsg, key)
		case "retryTime":
			err = unpopulate(val, "RetryTime", &s.RetryTime)
			delete(rawMsg, key)
		case "serialNumber":
			err = unpopulate(val, "SerialNumber", &s.SerialNumber)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SrvRecord.
func (s SrvRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "port", s.Port)
	populate(objectMap, "priority", s.Priority)
	populate(objectMap, "target", s.Target)
	populate(objectMap, "weight", s.Weight)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SrvRecord.
func (s *SrvRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "port":
			err = unpopulate(val, "Port", &s.Port)
			delete(rawMsg, key)
		case "priority":
			err = unpopulate(val, "Priority", &s.Priority)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &s.Target)
			delete(rawMsg, key)
		case "weight":
			err = unpopulate(val, "Weight", &s.Weight)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SubResource.
func (s SubResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubResource.
func (s *SubResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrackedResource.
func (t *TrackedResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &t.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TxtRecord.
func (t TxtRecord) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TxtRecord.
func (t *TxtRecord) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &t.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLink.
func (v VirtualNetworkLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "etag", v.Etag)
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkLink.
func (v *VirtualNetworkLink) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "etag":
			err = unpopulate(val, "Etag", &v.Etag)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &v.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &v.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &v.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &v.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &v.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkListResult.
func (v VirtualNetworkLinkListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkLinkListResult.
func (v *VirtualNetworkLinkListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &v.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &v.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkProperties.
func (v VirtualNetworkLinkProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "registrationEnabled", v.RegistrationEnabled)
	populate(objectMap, "resolutionPolicy", v.ResolutionPolicy)
	populate(objectMap, "virtualNetwork", v.VirtualNetwork)
	populate(objectMap, "virtualNetworkLinkState", v.VirtualNetworkLinkState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkLinkProperties.
func (v *VirtualNetworkLinkProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &v.ProvisioningState)
			delete(rawMsg, key)
		case "registrationEnabled":
			err = unpopulate(val, "RegistrationEnabled", &v.RegistrationEnabled)
			delete(rawMsg, key)
		case "resolutionPolicy":
			err = unpopulate(val, "ResolutionPolicy", &v.ResolutionPolicy)
			delete(rawMsg, key)
		case "virtualNetwork":
			err = unpopulate(val, "VirtualNetwork", &v.VirtualNetwork)
			delete(rawMsg, key)
		case "virtualNetworkLinkState":
			err = unpopulate(val, "VirtualNetworkLinkState", &v.VirtualNetworkLinkState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
