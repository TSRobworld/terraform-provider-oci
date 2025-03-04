// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalExadataInfrastructureDiscoverySummary The discovery result for the Exadata system infrastructure.
type ExternalExadataInfrastructureDiscoverySummary struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Null for new discover case.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the agent could be used for monitoring.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated connector.
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// The version of the entity.
	Version *string `mandatory:"false" json:"version"`

	// The internal identifier.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the entity.
	Status *string `mandatory:"false" json:"status"`

	// The error code of the discovery on the resource
	DiscoverErrorCode *string `mandatory:"false" json:"discoverErrorCode"`

	// The error message of the discovery on the resource
	DiscoverErrorMsg *string `mandatory:"false" json:"discoverErrorMsg"`

	// The size of the Exadata infrastructure.
	RackSize ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The status of the entity discover.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

//GetId returns Id
func (m ExternalExadataInfrastructureDiscoverySummary) GetId() *string {
	return m.Id
}

//GetAgentId returns AgentId
func (m ExternalExadataInfrastructureDiscoverySummary) GetAgentId() *string {
	return m.AgentId
}

//GetConnectorId returns ConnectorId
func (m ExternalExadataInfrastructureDiscoverySummary) GetConnectorId() *string {
	return m.ConnectorId
}

//GetDisplayName returns DisplayName
func (m ExternalExadataInfrastructureDiscoverySummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetVersion returns Version
func (m ExternalExadataInfrastructureDiscoverySummary) GetVersion() *string {
	return m.Version
}

//GetInternalId returns InternalId
func (m ExternalExadataInfrastructureDiscoverySummary) GetInternalId() *string {
	return m.InternalId
}

//GetStatus returns Status
func (m ExternalExadataInfrastructureDiscoverySummary) GetStatus() *string {
	return m.Status
}

//GetDiscoverStatus returns DiscoverStatus
func (m ExternalExadataInfrastructureDiscoverySummary) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

//GetDiscoverErrorCode returns DiscoverErrorCode
func (m ExternalExadataInfrastructureDiscoverySummary) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

//GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m ExternalExadataInfrastructureDiscoverySummary) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m ExternalExadataInfrastructureDiscoverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataInfrastructureDiscoverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExternalExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEntityDiscoveredDiscoverStatusEnum(string(m.DiscoverStatus)); !ok && m.DiscoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoverStatus: %s. Supported values are: %s.", m.DiscoverStatus, strings.Join(GetEntityDiscoveredDiscoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalExadataInfrastructureDiscoverySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataInfrastructureDiscoverySummary ExternalExadataInfrastructureDiscoverySummary
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeExternalExadataInfrastructureDiscoverySummary
	}{
		"INFRASTRUCTURE_DISCOVER_SUMMARY",
		(MarshalTypeExternalExadataInfrastructureDiscoverySummary)(m),
	}

	return json.Marshal(&s)
}

// ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum Enum with underlying type: string
type ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum
const (
	ExternalExadataInfrastructureDiscoverySummaryRackSizeFull    ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum = "FULL"
	ExternalExadataInfrastructureDiscoverySummaryRackSizeHalf    ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum = "HALF"
	ExternalExadataInfrastructureDiscoverySummaryRackSizeQuarter ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum = "QUARTER"
	ExternalExadataInfrastructureDiscoverySummaryRackSizeEighth  ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum = "EIGHTH"
)

var mappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnum = map[string]ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum{
	"FULL":    ExternalExadataInfrastructureDiscoverySummaryRackSizeFull,
	"HALF":    ExternalExadataInfrastructureDiscoverySummaryRackSizeHalf,
	"QUARTER": ExternalExadataInfrastructureDiscoverySummaryRackSizeQuarter,
	"EIGHTH":  ExternalExadataInfrastructureDiscoverySummaryRackSizeEighth,
}

var mappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnumLowerCase = map[string]ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum{
	"full":    ExternalExadataInfrastructureDiscoverySummaryRackSizeFull,
	"half":    ExternalExadataInfrastructureDiscoverySummaryRackSizeHalf,
	"quarter": ExternalExadataInfrastructureDiscoverySummaryRackSizeQuarter,
	"eighth":  ExternalExadataInfrastructureDiscoverySummaryRackSizeEighth,
}

// GetExternalExadataInfrastructureDiscoverySummaryRackSizeEnumValues Enumerates the set of values for ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum
func GetExternalExadataInfrastructureDiscoverySummaryRackSizeEnumValues() []ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum {
	values := make([]ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum
func GetExternalExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnum(val string) (ExternalExadataInfrastructureDiscoverySummaryRackSizeEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureDiscoverySummaryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
