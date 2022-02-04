// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"strings"
)

// UpdateBuildStageDetails Specifies the build stage.
type UpdateBuildStageDetails struct {

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the build stage.
	Description *string `mandatory:"false" json:"description"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the build stage execution. Specify value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	BuildSourceCollection *BuildSourceCollection `mandatory:"false" json:"buildSourceCollection"`

	// Name of the build source where the build_spec.yml file is located. If not specified, the first entry in the build source collection is chosen as primary build source.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	// Image name for the build environment.
	Image BuildStageImageEnum `mandatory:"false" json:"image,omitempty"`
}

//GetDisplayName returns DisplayName
func (m UpdateBuildStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateBuildStageDetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m UpdateBuildStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateBuildStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateBuildStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateBuildStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBuildStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingBuildStageImageEnum[string(m.Image)]; !ok && m.Image != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Image: %s. Supported values are: %s.", m.Image, strings.Join(GetBuildStageImageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateBuildStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateBuildStageDetails UpdateBuildStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeUpdateBuildStageDetails
	}{
		"BUILD",
		(MarshalTypeUpdateBuildStageDetails)(m),
	}

	return json.Marshal(&s)
}
