---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_systems"
sidebar_current: "docs-oci-datasource-database_management-external_db_systems"
description: |-
  Provides the list of External Db Systems in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_systems
This data source provides the list of External Db Systems in Oracle Cloud Infrastructure Database Management service.

Lists the external DB systems in the specified compartment.

## Example Usage

```hcl
data "oci_database_management_external_db_systems" "test_external_db_systems" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.external_db_system_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `external_db_system_collection` - The list of external_db_system_collection.

### ExternalDbSystem Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_management_config` - The details required to enable Database Management for an external DB system.
	* `license_model` - The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `discovery_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `home_directory` - The Oracle Grid home directory in case of cluster-based DB system and Oracle home directory in case of single instance-based DB system. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `is_cluster` - Indicates whether the DB system is a cluster DB system or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external DB system resource.
* `time_created` - The date and time the external DB system was created.
* `time_updated` - The date and time the external DB system was last updated.

