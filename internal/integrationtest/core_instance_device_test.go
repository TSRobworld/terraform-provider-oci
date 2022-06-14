// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"
)

var (
	instanceDeviceDataSourceRepresentation = map[string]interface{}{
		"instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"is_available": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"name":         acctest.Representation{RepType: acctest.Optional, Create: `/dev/oracleoci/oraclevdb`},
	}

	InstanceDeviceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceDeviceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceDeviceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_instance_devices.test_instance_devices"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_devices", "test_instance_devices", acctest.Optional, acctest.Create, instanceDeviceDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceDeviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_available", "true"),
				resource.TestCheckResourceAttr(datasourceName, "name", "/dev/oracleoci/oraclevdb"),

				resource.TestCheckResourceAttrSet(datasourceName, "devices.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "devices.0.is_available"),
				resource.TestCheckResourceAttrSet(datasourceName, "devices.0.name"),
			),
		},
	})
}
