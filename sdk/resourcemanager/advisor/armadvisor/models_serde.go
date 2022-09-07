//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armadvisor

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ARMErrorResponseBody.
func (a ARMErrorResponseBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", a.Code)
	populate(objectMap, "message", a.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ARMErrorResponseBody.
func (a *ARMErrorResponseBody) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &a.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &a.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ArmErrorResponse.
func (a ArmErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", a.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArmErrorResponse.
func (a *ArmErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &a.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConfigData.
func (c ConfigData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigData.
func (c *ConfigData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDataProperties.
func (c ConfigDataProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "digests", c.Digests)
	populate(objectMap, "duration", c.Duration)
	populate(objectMap, "exclude", c.Exclude)
	populate(objectMap, "lowCpuThreshold", c.LowCPUThreshold)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigDataProperties.
func (c *ConfigDataProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "digests":
			err = unpopulate(val, "Digests", &c.Digests)
			delete(rawMsg, key)
		case "duration":
			err = unpopulate(val, "Duration", &c.Duration)
			delete(rawMsg, key)
		case "exclude":
			err = unpopulate(val, "Exclude", &c.Exclude)
			delete(rawMsg, key)
		case "lowCpuThreshold":
			err = unpopulate(val, "LowCPUThreshold", &c.LowCPUThreshold)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationListResult.
func (c ConfigurationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigurationListResult.
func (c *ConfigurationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &c.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &c.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DigestConfig.
func (d DigestConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionGroupResourceId", d.ActionGroupResourceID)
	populate(objectMap, "categories", d.Categories)
	populate(objectMap, "frequency", d.Frequency)
	populate(objectMap, "language", d.Language)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "state", d.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DigestConfig.
func (d *DigestConfig) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionGroupResourceId":
			err = unpopulate(val, "ActionGroupResourceID", &d.ActionGroupResourceID)
			delete(rawMsg, key)
		case "categories":
			err = unpopulate(val, "Categories", &d.Categories)
			delete(rawMsg, key)
		case "frequency":
			err = unpopulate(val, "Frequency", &d.Frequency)
			delete(rawMsg, key)
		case "language":
			err = unpopulate(val, "Language", &d.Language)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &d.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataEntity.
func (m MetadataEntity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataEntity.
func (m *MetadataEntity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataEntityListResult.
func (m MetadataEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataEntityListResult.
func (m *MetadataEntityListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &m.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataEntityProperties.
func (m MetadataEntityProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicableScenarios", m.ApplicableScenarios)
	populate(objectMap, "dependsOn", m.DependsOn)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "supportedValues", m.SupportedValues)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataEntityProperties.
func (m *MetadataEntityProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "applicableScenarios":
			err = unpopulate(val, "ApplicableScenarios", &m.ApplicableScenarios)
			delete(rawMsg, key)
		case "dependsOn":
			err = unpopulate(val, "DependsOn", &m.DependsOn)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &m.DisplayName)
			delete(rawMsg, key)
		case "supportedValues":
			err = unpopulate(val, "SupportedValues", &m.SupportedValues)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataSupportedValueDetail.
func (m MetadataSupportedValueDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "id", m.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataSupportedValueDetail.
func (m *MetadataSupportedValueDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &m.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplayInfo.
func (o OperationDisplayInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplayInfo.
func (o *OperationDisplayInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationEntity.
func (o OperationEntity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationEntity.
func (o *OperationEntity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationEntityListResult.
func (o OperationEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationEntityListResult.
func (o *OperationEntityListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PredictionRequest.
func (p PredictionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PredictionRequest.
func (p *PredictionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PredictionRequestProperties.
func (p PredictionRequestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedProperties", &p.ExtendedProperties)
	populate(objectMap, "predictionType", p.PredictionType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PredictionRequestProperties.
func (p *PredictionRequestProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "extendedProperties":
			err = unpopulate(val, "ExtendedProperties", &p.ExtendedProperties)
			delete(rawMsg, key)
		case "predictionType":
			err = unpopulate(val, "PredictionType", &p.PredictionType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PredictionResponse.
func (p PredictionResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PredictionResponse.
func (p *PredictionResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PredictionResponseProperties.
func (p PredictionResponseProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", p.Category)
	populate(objectMap, "extendedProperties", &p.ExtendedProperties)
	populate(objectMap, "impact", p.Impact)
	populate(objectMap, "impactedField", p.ImpactedField)
	populateTimeRFC3339(objectMap, "lastUpdated", p.LastUpdated)
	populate(objectMap, "predictionType", p.PredictionType)
	populate(objectMap, "shortDescription", p.ShortDescription)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PredictionResponseProperties.
func (p *PredictionResponseProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
			err = unpopulate(val, "Category", &p.Category)
			delete(rawMsg, key)
		case "extendedProperties":
			err = unpopulate(val, "ExtendedProperties", &p.ExtendedProperties)
			delete(rawMsg, key)
		case "impact":
			err = unpopulate(val, "Impact", &p.Impact)
			delete(rawMsg, key)
		case "impactedField":
			err = unpopulate(val, "ImpactedField", &p.ImpactedField)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, "LastUpdated", &p.LastUpdated)
			delete(rawMsg, key)
		case "predictionType":
			err = unpopulate(val, "PredictionType", &p.PredictionType)
			delete(rawMsg, key)
		case "shortDescription":
			err = unpopulate(val, "ShortDescription", &p.ShortDescription)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecommendationProperties.
func (r RecommendationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", r.Actions)
	populate(objectMap, "category", r.Category)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "exposedMetadataProperties", r.ExposedMetadataProperties)
	populate(objectMap, "extendedProperties", r.ExtendedProperties)
	populate(objectMap, "impact", r.Impact)
	populate(objectMap, "impactedField", r.ImpactedField)
	populate(objectMap, "impactedValue", r.ImpactedValue)
	populate(objectMap, "label", r.Label)
	populateTimeRFC3339(objectMap, "lastUpdated", r.LastUpdated)
	populate(objectMap, "learnMoreLink", r.LearnMoreLink)
	populate(objectMap, "metadata", r.Metadata)
	populate(objectMap, "potentialBenefits", r.PotentialBenefits)
	populate(objectMap, "recommendationTypeId", r.RecommendationTypeID)
	populate(objectMap, "remediation", r.Remediation)
	populate(objectMap, "resourceMetadata", r.ResourceMetadata)
	populate(objectMap, "shortDescription", r.ShortDescription)
	populate(objectMap, "suppressionIds", r.SuppressionIDs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecommendationProperties.
func (r *RecommendationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			err = unpopulate(val, "Actions", &r.Actions)
			delete(rawMsg, key)
		case "category":
			err = unpopulate(val, "Category", &r.Category)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "exposedMetadataProperties":
			err = unpopulate(val, "ExposedMetadataProperties", &r.ExposedMetadataProperties)
			delete(rawMsg, key)
		case "extendedProperties":
			err = unpopulate(val, "ExtendedProperties", &r.ExtendedProperties)
			delete(rawMsg, key)
		case "impact":
			err = unpopulate(val, "Impact", &r.Impact)
			delete(rawMsg, key)
		case "impactedField":
			err = unpopulate(val, "ImpactedField", &r.ImpactedField)
			delete(rawMsg, key)
		case "impactedValue":
			err = unpopulate(val, "ImpactedValue", &r.ImpactedValue)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, "Label", &r.Label)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, "LastUpdated", &r.LastUpdated)
			delete(rawMsg, key)
		case "learnMoreLink":
			err = unpopulate(val, "LearnMoreLink", &r.LearnMoreLink)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, "Metadata", &r.Metadata)
			delete(rawMsg, key)
		case "potentialBenefits":
			err = unpopulate(val, "PotentialBenefits", &r.PotentialBenefits)
			delete(rawMsg, key)
		case "recommendationTypeId":
			err = unpopulate(val, "RecommendationTypeID", &r.RecommendationTypeID)
			delete(rawMsg, key)
		case "remediation":
			err = unpopulate(val, "Remediation", &r.Remediation)
			delete(rawMsg, key)
		case "resourceMetadata":
			err = unpopulate(val, "ResourceMetadata", &r.ResourceMetadata)
			delete(rawMsg, key)
		case "shortDescription":
			err = unpopulate(val, "ShortDescription", &r.ShortDescription)
			delete(rawMsg, key)
		case "suppressionIds":
			err = unpopulate(val, "SuppressionIDs", &r.SuppressionIDs)
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
	objectMap := make(map[string]interface{})
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

// MarshalJSON implements the json.Marshaller interface for type ResourceMetadata.
func (r ResourceMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "action", r.Action)
	populate(objectMap, "plural", r.Plural)
	populate(objectMap, "resourceId", r.ResourceID)
	populate(objectMap, "singular", r.Singular)
	populate(objectMap, "source", r.Source)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceMetadata.
func (r *ResourceMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "action":
			err = unpopulate(val, "Action", &r.Action)
			delete(rawMsg, key)
		case "plural":
			err = unpopulate(val, "Plural", &r.Plural)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &r.ResourceID)
			delete(rawMsg, key)
		case "singular":
			err = unpopulate(val, "Singular", &r.Singular)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &r.Source)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceRecommendationBase.
func (r ResourceRecommendationBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceRecommendationBase.
func (r *ResourceRecommendationBase) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type ResourceRecommendationBaseListResult.
func (r ResourceRecommendationBaseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceRecommendationBaseListResult.
func (r *ResourceRecommendationBaseListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type ShortDescription.
func (s ShortDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "problem", s.Problem)
	populate(objectMap, "solution", s.Solution)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ShortDescription.
func (s *ShortDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "problem":
			err = unpopulate(val, "Problem", &s.Problem)
			delete(rawMsg, key)
		case "solution":
			err = unpopulate(val, "Solution", &s.Solution)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SuppressionContract.
func (s SuppressionContract) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SuppressionContract.
func (s *SuppressionContract) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SuppressionContractListResult.
func (s SuppressionContractListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SuppressionContractListResult.
func (s *SuppressionContractListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SuppressionProperties.
func (s SuppressionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "expirationTimeStamp", s.ExpirationTimeStamp)
	populate(objectMap, "suppressionId", s.SuppressionID)
	populate(objectMap, "ttl", s.TTL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SuppressionProperties.
func (s *SuppressionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTimeStamp":
			err = unpopulateTimeRFC3339(val, "ExpirationTimeStamp", &s.ExpirationTimeStamp)
			delete(rawMsg, key)
		case "suppressionId":
			err = unpopulate(val, "SuppressionID", &s.SuppressionID)
			delete(rawMsg, key)
		case "ttl":
			err = unpopulate(val, "TTL", &s.TTL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
