---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
#
# ----------------------------------------------------------------------------
#
#     This file is managed by Magic Modules (https:#github.com/GoogleCloudPlatform/magic-modules)
#     and is based on the DCL (https:#github.com/GoogleCloudPlatform/declarative-resource-client-library).
#     Changes will need to be made to the DCL or Magic Modules instead of here.
#
#     We are not currently able to accept contributions to this file. If changes
#     are required, please file an issue at https:#github.com/hashicorp/terraform-provider-google/issues/new/choose
#
# ----------------------------------------------------------------------------
subcategory: "Cloud (Stackdriver) Monitoring"
layout: "google"
page_title: "Google: google_monitoring_monitored_project"
sidebar_current: "docs-google-monitoring-monitored-project"
description: |-
Beta only
---

# google\_monitoring\_monitored\_project

Beta only
## Example Usage - basic_monitored_project
A basic example of a monitoring monitored project
```hcl
resource "google_monitoring_monitored_project" "primary" {
  metrics_scope = "my-project-name"
  name          = google_project.basic.name
  provider      = google-beta
}
resource "google_project" "basic" {
  project_id = "id"
  name       = "id"
  org_id     = "123456789"
  provider   = google-beta
}

```

## Argument Reference

The following arguments are supported:

* `metrics_scope` -
  (Required)
  Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
  
* `name` -
  (Required)
  Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
  


- - -



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}`

* `create_time` -
  Output only. The time when this `MonitoredProject` was created.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 10 minutes.
- `delete` - Default is 10 minutes.

## Import

MonitoredProject can be imported using any of these accepted formats:

```
$ terraform import google_monitoring_monitored_project.default locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}
$ terraform import google_monitoring_monitored_project.default {{metrics_scope}}/{{name}}
```



