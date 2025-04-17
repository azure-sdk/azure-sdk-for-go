// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import "encoding/json"

// UnmarshalJSON implements the json.Unmarshaller interface for type ComputeClientListKeysResponse.
func (c *ComputeClientListKeysResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalComputeSecretsClassification(data)
	if err != nil {
		return err
	}
	c.ComputeSecretsClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectionRaiBlocklistItemClientAddBulkResponse.
func (c *ConnectionRaiBlocklistItemClientAddBulkResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.RaiBlocklistItemPropertiesBasicResourceArray)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatastoresClientListSecretsResponse.
func (d *DatastoresClientListSecretsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDatastoreSecretsClassification(data)
	if err != nil {
		return err
	}
	d.DatastoreSecretsClassification = res
	return nil
}
