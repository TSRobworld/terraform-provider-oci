// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
	"strings"
)

// ListInstanceAgentPluginsRequest wrapper for the ListInstanceAgentPlugins operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentPlugins.go.html to see an example of how to use ListInstanceAgentPluginsRequest.
type ListInstanceAgentPluginsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the instance.
	InstanceagentId *string `mandatory:"true" contributesTo:"path" name:"instanceagentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The plugin status
	Status ListInstanceAgentPluginsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// `TIMECREATED` is descending.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInstanceAgentPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The `DISPLAYNAME` sort order
	// is case sensitive.
	SortOrder ListInstanceAgentPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The plugin name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstanceAgentPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstanceAgentPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstanceAgentPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstanceAgentPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstanceAgentPluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListInstanceAgentPluginsStatusEnum[string(request.Status)]; !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListInstanceAgentPluginsStatusEnumStringValues(), ",")))
	}
	if _, ok := mappingListInstanceAgentPluginsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstanceAgentPluginsSortByEnumStringValues(), ",")))
	}
	if _, ok := mappingListInstanceAgentPluginsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstanceAgentPluginsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstanceAgentPluginsResponse wrapper for the ListInstanceAgentPlugins operation
type ListInstanceAgentPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceAgentPluginSummary instances
	Items []InstanceAgentPluginSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstanceAgentPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstanceAgentPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstanceAgentPluginsStatusEnum Enum with underlying type: string
type ListInstanceAgentPluginsStatusEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsStatusEnum
const (
	ListInstanceAgentPluginsStatusRunning      ListInstanceAgentPluginsStatusEnum = "RUNNING"
	ListInstanceAgentPluginsStatusStopped      ListInstanceAgentPluginsStatusEnum = "STOPPED"
	ListInstanceAgentPluginsStatusNotSupported ListInstanceAgentPluginsStatusEnum = "NOT_SUPPORTED"
	ListInstanceAgentPluginsStatusInvalid      ListInstanceAgentPluginsStatusEnum = "INVALID"
)

var mappingListInstanceAgentPluginsStatusEnum = map[string]ListInstanceAgentPluginsStatusEnum{
	"RUNNING":       ListInstanceAgentPluginsStatusRunning,
	"STOPPED":       ListInstanceAgentPluginsStatusStopped,
	"NOT_SUPPORTED": ListInstanceAgentPluginsStatusNotSupported,
	"INVALID":       ListInstanceAgentPluginsStatusInvalid,
}

// GetListInstanceAgentPluginsStatusEnumValues Enumerates the set of values for ListInstanceAgentPluginsStatusEnum
func GetListInstanceAgentPluginsStatusEnumValues() []ListInstanceAgentPluginsStatusEnum {
	values := make([]ListInstanceAgentPluginsStatusEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentPluginsStatusEnumStringValues Enumerates the set of values in String for ListInstanceAgentPluginsStatusEnum
func GetListInstanceAgentPluginsStatusEnumStringValues() []string {
	return []string{
		"RUNNING",
		"STOPPED",
		"NOT_SUPPORTED",
		"INVALID",
	}
}

// ListInstanceAgentPluginsSortByEnum Enum with underlying type: string
type ListInstanceAgentPluginsSortByEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsSortByEnum
const (
	ListInstanceAgentPluginsSortByTimecreated ListInstanceAgentPluginsSortByEnum = "TIMECREATED"
	ListInstanceAgentPluginsSortByDisplayname ListInstanceAgentPluginsSortByEnum = "DISPLAYNAME"
)

var mappingListInstanceAgentPluginsSortByEnum = map[string]ListInstanceAgentPluginsSortByEnum{
	"TIMECREATED": ListInstanceAgentPluginsSortByTimecreated,
	"DISPLAYNAME": ListInstanceAgentPluginsSortByDisplayname,
}

// GetListInstanceAgentPluginsSortByEnumValues Enumerates the set of values for ListInstanceAgentPluginsSortByEnum
func GetListInstanceAgentPluginsSortByEnumValues() []ListInstanceAgentPluginsSortByEnum {
	values := make([]ListInstanceAgentPluginsSortByEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentPluginsSortByEnumStringValues Enumerates the set of values in String for ListInstanceAgentPluginsSortByEnum
func GetListInstanceAgentPluginsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// ListInstanceAgentPluginsSortOrderEnum Enum with underlying type: string
type ListInstanceAgentPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsSortOrderEnum
const (
	ListInstanceAgentPluginsSortOrderAsc  ListInstanceAgentPluginsSortOrderEnum = "ASC"
	ListInstanceAgentPluginsSortOrderDesc ListInstanceAgentPluginsSortOrderEnum = "DESC"
)

var mappingListInstanceAgentPluginsSortOrderEnum = map[string]ListInstanceAgentPluginsSortOrderEnum{
	"ASC":  ListInstanceAgentPluginsSortOrderAsc,
	"DESC": ListInstanceAgentPluginsSortOrderDesc,
}

// GetListInstanceAgentPluginsSortOrderEnumValues Enumerates the set of values for ListInstanceAgentPluginsSortOrderEnum
func GetListInstanceAgentPluginsSortOrderEnumValues() []ListInstanceAgentPluginsSortOrderEnum {
	values := make([]ListInstanceAgentPluginsSortOrderEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentPluginsSortOrderEnumStringValues Enumerates the set of values in String for ListInstanceAgentPluginsSortOrderEnum
func GetListInstanceAgentPluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}
