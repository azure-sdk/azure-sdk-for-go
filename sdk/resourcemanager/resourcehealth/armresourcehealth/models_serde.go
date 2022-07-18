//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Event.
func (e Event) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Event.
func (e *Event) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &e.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &e.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventProperties.
func (e EventProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInformation", e.AdditionalInformation)
	populate(objectMap, "article", e.Article)
	populate(objectMap, "description", e.Description)
	populate(objectMap, "duration", e.Duration)
	populate(objectMap, "enableChatWithUs", e.EnableChatWithUs)
	populate(objectMap, "enableMicrosoftSupport", e.EnableMicrosoftSupport)
	populate(objectMap, "eventLevel", e.EventLevel)
	populate(objectMap, "eventSource", e.EventSource)
	populate(objectMap, "eventType", e.EventType)
	populate(objectMap, "faqs", e.Faqs)
	populate(objectMap, "header", e.Header)
	populate(objectMap, "hirStage", e.HirStage)
	populate(objectMap, "impact", e.Impact)
	populateTimeRFC3339(objectMap, "impactMitigationTime", e.ImpactMitigationTime)
	populateTimeRFC3339(objectMap, "impactStartTime", e.ImpactStartTime)
	populate(objectMap, "impactType", e.ImpactType)
	populate(objectMap, "isHIR", e.IsHIR)
	populateTimeRFC3339(objectMap, "lastUpdateTime", e.LastUpdateTime)
	populate(objectMap, "level", e.Level)
	populate(objectMap, "links", e.Links)
	populate(objectMap, "platformInitiated", e.PlatformInitiated)
	populate(objectMap, "priority", e.Priority)
	populate(objectMap, "recommendedActions", e.RecommendedActions)
	populate(objectMap, "status", e.Status)
	populate(objectMap, "summary", e.Summary)
	populate(objectMap, "title", e.Title)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventProperties.
func (e *EventProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInformation":
			err = unpopulate(val, "AdditionalInformation", &e.AdditionalInformation)
			delete(rawMsg, key)
		case "article":
			err = unpopulate(val, "Article", &e.Article)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &e.Description)
			delete(rawMsg, key)
		case "duration":
			err = unpopulate(val, "Duration", &e.Duration)
			delete(rawMsg, key)
		case "enableChatWithUs":
			err = unpopulate(val, "EnableChatWithUs", &e.EnableChatWithUs)
			delete(rawMsg, key)
		case "enableMicrosoftSupport":
			err = unpopulate(val, "EnableMicrosoftSupport", &e.EnableMicrosoftSupport)
			delete(rawMsg, key)
		case "eventLevel":
			err = unpopulate(val, "EventLevel", &e.EventLevel)
			delete(rawMsg, key)
		case "eventSource":
			err = unpopulate(val, "EventSource", &e.EventSource)
			delete(rawMsg, key)
		case "eventType":
			err = unpopulate(val, "EventType", &e.EventType)
			delete(rawMsg, key)
		case "faqs":
			err = unpopulate(val, "Faqs", &e.Faqs)
			delete(rawMsg, key)
		case "header":
			err = unpopulate(val, "Header", &e.Header)
			delete(rawMsg, key)
		case "hirStage":
			err = unpopulate(val, "HirStage", &e.HirStage)
			delete(rawMsg, key)
		case "impact":
			err = unpopulate(val, "Impact", &e.Impact)
			delete(rawMsg, key)
		case "impactMitigationTime":
			err = unpopulateTimeRFC3339(val, "ImpactMitigationTime", &e.ImpactMitigationTime)
			delete(rawMsg, key)
		case "impactStartTime":
			err = unpopulateTimeRFC3339(val, "ImpactStartTime", &e.ImpactStartTime)
			delete(rawMsg, key)
		case "impactType":
			err = unpopulate(val, "ImpactType", &e.ImpactType)
			delete(rawMsg, key)
		case "isHIR":
			err = unpopulate(val, "IsHIR", &e.IsHIR)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateTimeRFC3339(val, "LastUpdateTime", &e.LastUpdateTime)
			delete(rawMsg, key)
		case "level":
			err = unpopulate(val, "Level", &e.Level)
			delete(rawMsg, key)
		case "links":
			err = unpopulate(val, "Links", &e.Links)
			delete(rawMsg, key)
		case "platformInitiated":
			err = unpopulate(val, "PlatformInitiated", &e.PlatformInitiated)
			delete(rawMsg, key)
		case "priority":
			err = unpopulate(val, "Priority", &e.Priority)
			delete(rawMsg, key)
		case "recommendedActions":
			err = unpopulate(val, "RecommendedActions", &e.RecommendedActions)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &e.Status)
			delete(rawMsg, key)
		case "summary":
			err = unpopulate(val, "Summary", &e.Summary)
			delete(rawMsg, key)
		case "title":
			err = unpopulate(val, "Title", &e.Title)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventPropertiesAdditionalInformation.
func (e EventPropertiesAdditionalInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventPropertiesAdditionalInformation.
func (e *EventPropertiesAdditionalInformation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventPropertiesArticle.
func (e EventPropertiesArticle) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "articleContent", e.ArticleContent)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventPropertiesArticle.
func (e *EventPropertiesArticle) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "articleContent":
			err = unpopulate(val, "ArticleContent", &e.ArticleContent)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventPropertiesRecommendedActions.
func (e EventPropertiesRecommendedActions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", e.Actions)
	populate(objectMap, "localeCode", e.LocaleCode)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventPropertiesRecommendedActions.
func (e *EventPropertiesRecommendedActions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			err = unpopulate(val, "Actions", &e.Actions)
			delete(rawMsg, key)
		case "localeCode":
			err = unpopulate(val, "LocaleCode", &e.LocaleCode)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventPropertiesRecommendedActionsItem.
func (e EventPropertiesRecommendedActionsItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionText", e.ActionText)
	populate(objectMap, "groupId", e.GroupID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventPropertiesRecommendedActionsItem.
func (e *EventPropertiesRecommendedActionsItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionText":
			err = unpopulate(val, "ActionText", &e.ActionText)
			delete(rawMsg, key)
		case "groupId":
			err = unpopulate(val, "GroupID", &e.GroupID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Events.
func (e Events) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Events.
func (e *Events) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &e.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Faq.
func (f Faq) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "answer", f.Answer)
	populate(objectMap, "localeCode", f.LocaleCode)
	populate(objectMap, "question", f.Question)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Faq.
func (f *Faq) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "answer":
			err = unpopulate(val, "Answer", &f.Answer)
			delete(rawMsg, key)
		case "localeCode":
			err = unpopulate(val, "LocaleCode", &f.LocaleCode)
			delete(rawMsg, key)
		case "question":
			err = unpopulate(val, "Question", &f.Question)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Impact.
func (i Impact) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "impactedRegions", i.ImpactedRegions)
	populate(objectMap, "impactedService", i.ImpactedService)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Impact.
func (i *Impact) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "impactedRegions":
			err = unpopulate(val, "ImpactedRegions", &i.ImpactedRegions)
			delete(rawMsg, key)
		case "impactedService":
			err = unpopulate(val, "ImpactedService", &i.ImpactedService)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImpactedServiceRegion.
func (i ImpactedServiceRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "impactedRegion", i.ImpactedRegion)
	populate(objectMap, "impactedSubscriptions", i.ImpactedSubscriptions)
	populateTimeRFC3339(objectMap, "lastUpdateTime", i.LastUpdateTime)
	populate(objectMap, "status", i.Status)
	populate(objectMap, "updates", i.Updates)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImpactedServiceRegion.
func (i *ImpactedServiceRegion) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "impactedRegion":
			err = unpopulate(val, "ImpactedRegion", &i.ImpactedRegion)
			delete(rawMsg, key)
		case "impactedSubscriptions":
			err = unpopulate(val, "ImpactedSubscriptions", &i.ImpactedSubscriptions)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateTimeRFC3339(val, "LastUpdateTime", &i.LastUpdateTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &i.Status)
			delete(rawMsg, key)
		case "updates":
			err = unpopulate(val, "Updates", &i.Updates)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Link.
func (l Link) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bladeName", l.BladeName)
	populate(objectMap, "displayText", l.DisplayText)
	populate(objectMap, "extensionName", l.ExtensionName)
	populate(objectMap, "parameters", &l.Parameters)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Link.
func (l *Link) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bladeName":
			err = unpopulate(val, "BladeName", &l.BladeName)
			delete(rawMsg, key)
		case "displayText":
			err = unpopulate(val, "DisplayText", &l.DisplayText)
			delete(rawMsg, key)
		case "extensionName":
			err = unpopulate(val, "ExtensionName", &l.ExtensionName)
			delete(rawMsg, key)
		case "parameters":
			err = unpopulate(val, "Parameters", &l.Parameters)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &l.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LinkDisplayText.
func (l LinkDisplayText) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "localizedValue", l.LocalizedValue)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LinkDisplayText.
func (l *LinkDisplayText) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "localizedValue":
			err = unpopulate(val, "LocalizedValue", &l.LocalizedValue)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
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
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
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

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

// MarshalJSON implements the json.Marshaller interface for type Update.
func (u Update) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "summary", u.Summary)
	populateTimeRFC3339(objectMap, "updateDateTime", u.UpdateDateTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Update.
func (u *Update) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "summary":
			err = unpopulate(val, "Summary", &u.Summary)
			delete(rawMsg, key)
		case "updateDateTime":
			err = unpopulateTimeRFC3339(val, "UpdateDateTime", &u.UpdateDateTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
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
