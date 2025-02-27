---
layout: "google"
page_title: "Terraform Google Provider 4.0.0 Upgrade Guide"
sidebar_current: "docs-google-provider-guides-version-4-upgrade"
description: |-
  Terraform Google Provider 4.0.0 Upgrade Guide
---

<!-- TOC depthFrom:2 depthTo:2 -->

- [Terraform Google Provider 4.0.0 Upgrade Guide](#terraform-google-provider-400-upgrade-guide)
  - [I accidentally upgraded to 4.0.0, how do I downgrade to `3.X`?](#i-accidentally-upgraded-to-400-how-do-i-downgrade-to-3x)
  - [Provider Version Configuration](#provider-version-configuration)
  - [Provider](#provider)
    - [Redundant default scopes are removed](#redundant-default-scopes-are-removed)
    - [Runtime Configurator (`runtimeconfig`) resources have been removed from the GA provider](#runtime-configurator-runtimeconfig-resources-have-been-removed-from-the-ga-provider)
    - [Service account scopes no longer accept `trace-append` or `trace-ro`, use `trace` instead](#service-account-scopes-no-longer-accept-trace-append-or-trace-ro-use-trace-instead)
  - [Datasource: `google_product_resource`](#datasource-google_product_resource)
    - [Datasource-level change example](#datasource-level-change-example)
  - [Resource: `google_bigquery_job`](#resource-google_bigquery_job)
    - [Exactly one of `query`, `load`, `copy` or `extract` is required](#exactly-one-of-query-load-copy-or-extract-is-required)
    - [At least one of `query.0.script_options.0.statement_timeout_ms`, `query.0.script_options.0.statement_byte_budget`, or `query.0.script_options.0.key_result_statement` is required](#at-least-one-of-query0script_options0statement_timeout_ms-query0script_options0statement_byte_budget-or-query0script_options0key_result_statement-is-required)
    - [Exactly one of `extract.0.source_table` or `extract.0.source_model` is required](#exactly-one-of-extract0source_table-or-extract0source_model-is-required)
  - [Resource: `google_cloudbuild_trigger`](#resource-google_cloudbuild_trigger)
    - [Exactly one of `build.0.source.0.repo_source.0.branch_name`, `build.0.source.0.repo_source.0.commit_sha` or `build.0.source.0.repo_source.0.tag_name` is required](#exactly-one-of-build0source0repo_source0branch_name-build0source0repo_source0commit_sha-or-build0source0repo_source0tag_name-is-required)
  - [Resource: `google_compute_autoscaler`](#resource-google_compute_autoscaler)
    - [At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas` or `autoscaling_policy.0.scale_down_control.0.time_window_sec` is required](#at-least-one-of-autoscaling_policy0scale_down_control0max_scaled_down_replicas-or-autoscaling_policy0scale_down_control0time_window_sec-is-required)
    - [At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed` or `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent` is required](#at-least-one-of-autoscaling_policy0scale_down_control0max_scaled_down_replicas0fixed-or-autoscaling_policy0scale_down_control0max_scaled_down_replicas0percent-is-required)
    - [At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas` or `autoscaling_policy.0.scale_in_control.0.time_window_sec` is required](#at-least-one-of-autoscaling_policy0scale_in_control0max_scaled_in_replicas-or-autoscaling_policy0scale_in_control0time_window_sec-is-required)
    - [At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed` or `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent` is required](#at-least-one-of-autoscaling_policy0scale_in_control0max_scaled_in_replicas0fixed-or-autoscaling_policy0scale_in_control0max_scaled_in_replicas0percent-is-required)
  - [Resource: `google_compute_region_autoscaler`](#resource-google_compute_region_autoscaler)
    - [At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas` or `autoscaling_policy.0.scale_down_control.0.time_window_sec` is required](#at-least-one-of-autoscaling_policy0scale_down_control0max_scaled_down_replicas-or-autoscaling_policy0scale_down_control0time_window_sec-is-required)
    - [At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed` or `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent` is required](#at-least-one-of-autoscaling_policy0scale_down_control0max_scaled_down_replicas0fixed-or-autoscaling_policy0scale_down_control0max_scaled_down_replicas0percent-is-required)
    - [At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas` or `autoscaling_policy.0.scale_in_control.0.time_window_sec` is required](#at-least-one-of-autoscaling_policy0scale_in_control0max_scaled_in_replicas-or-autoscaling_policy0scale_in_control0time_window_sec-is-required)
    - [At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed` or `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent` is required](#at-least-one-of-autoscaling_policy0scale_in_control0max_scaled_in_replicas0fixed-or-autoscaling_policy0scale_in_control0max_scaled_in_replicas0percent-is-required)
  - [Resource: `google_compute_firewall`](#resource-google_compute_firewall)
    - [One of `source_tags`, `source_ranges` or `source_service_accounts` are required on INGRESS firewalls](#one-of-source_tags-source_ranges-or-source_service_accounts-are-required-on-ingress-firewalls)
  - [Resource: `google_compute_instance_group_manager`](#resource-google_compute_instance_group_manager)
    - [`update_policy.min_ready_sec` is removed from the GA provider](#update_policymin_ready_sec-is-removed-from-the-ga-provider)
  - [Resource: `google_compute_region_instance_group_manager`](#resource-google_compute_region_instance_group_manager)
    - [`update_policy.min_ready_sec` is removed from the GA provider](#update_policymin_ready_sec-is-removed-from-the-ga-provider-1)
  - [Resource: `google_compute_instance_template`](#resource-google_compute_instance_template)
    - [`enable_display` is removed from the GA provider](#enable_display-is-removed-from-the-ga-provider)
  - [Resource: `google_compute_url_map`](#resource-google_compute_url_map)
    - [At least one of `default_route_action.0.fault_injection_policy.0.delay.0.fixed_delay` or `default_route_action.0.fault_injection_policy.0.delay.0.percentage` is required](#at-least-one-of-default_route_action0fault_injection_policy0delay0fixed_delay-or-default_route_action0fault_injection_policy0delay0percentage-is-required)
  - [Resource: `google_container_cluster`](#resource-google_container_cluster)
    - [`instance_group_urls` is now removed](#instance_group_urls-is-now-removed)
    - [`master_auth` is now removed](#master_auth-is-now-removed)
    - [`node_config.workload_metadata_config.node_metadata` is now removed](#node_configworkload_metadata_confignode_metadata-is-now-removed)
    - [`workload_identity_config.0.identity_namespace` is now removed](#workload_identity_config0identity_namespace-is-now-removed)
    - [`pod_security_policy_config` is removed from the GA provider](#pod_security_policy_config-is-removed-from-the-ga-provider)
  - [Resource: `google_project`](#resource-google_project)
    - [`org_id`, `folder_id` now conflict at plan time](#org_id-folder_id-now-confict-at-plan-time)
    - [`org_id`, `folder_id` are unset when removed from config](#org_id-folder_id-are-unset-when-removed-from-config)
  - [Resource: `google_project_service`](#resource-google_project_service)
    - [`bigquery-json.googleapis.com` is no longer a valid service name](#bigquery-json.googleapis.com-is-no-longer-a-valid-service-name)
  - [Resource: `google_data_loss_prevention_trigger`](#resource-google_data_loss_prevention_trigger)
    - [Exactly one of `inspect_job.0.storage_config.0.cloud_storage_options.0.file_set.0.url` or `inspect_job.0.storage_config.0.cloud_storage_options.0.file_set.0.regex_file_set` is required](#exactly-one-of-inspect_job0storage_config0cloud_storage_options0file_set0url-or-inspect_job0storage_config0cloud_storage_options0file_set0regex_file_set-is-required)
    - [At least one of `inspect_job.0.storage_config.0.timespan_config.0.start_time` or `inspect_job.0.storage_config.0.timespan_config.0.end_time` is required](#at-least-one-of-inspect_job0storage_config0timespan_config0start_time-or-inspect_job0storage_config0timespan_config0end_time-is-required)
  - [Resource: `google_os_config_patch_deployment`](#resource-google_os_config_patch_deployment)
    - [At least one of `patch_config.0.reboot_config`, `patch_config.0.apt`, `patch_config.0.yum`, `patch_config.0.goo` `patch_config.0.zypper`, `patch_config.0.windows_update`, `patch_config.0.pre_step` or `patch_config.0.pre_step` is required](#at-least-one-of-patch_config0reboot_config-patch_config0apt-patch_config0yum-patch_config0goo-patch_config0zypper-patch_config0windows_update-patch_config0pre_step-or-patch_config0pre_step-is-required)
    - [At least one of `patch_config.0.apt.0.type`, `patch_config.0.apt.0.excludes` or `patch_config.0.apt.0.exclusive_packages` is required](#at-least-one-of-patch_config0apt0type-patch_config0apt0excludes-or-patch_config0apt0exclusive_packages-is-required)
    - [At least one of `patch_config.0.yum.0.security`, `patch_config.0.yum.0.minimal`, `patch_config.0.yum.0.excludes` or `patch_config.0.yum.0.exclusive_packages` is required](#at-least-one-of-patch_config0yum0security-patch_config0yum0minimal-patch_config0yum0excludes-or-patch_config0yum0exclusive_packages-is-required)
    - [At least one of `patch_config.0.zypper.0.with_optional`, `patch_config.0.zypper.0.with_update`, `patch_config.0.zypper.0.categories`, `patch_config.0.zypper.0.severities`, `patch_config.0.zypper.0.excludes` or `patch_config.0.zypper.0.exclusive_patches` is required](#at-least-one-of-patch_config0zypper0with_optional-patch_config0zypper0with_update-patch_config0zypper0categories-patch_config0zypper0severities-patch_config0zypper0excludes-or-patch_config0zypper0exclusive_patches-is-required)
    - [Exactly one of `patch_config.0.windows_update.0.classifications`, `patch_config.0.windows_update.0.excludes` or `patch_config.0.windows_update.0.exclusive_patches` is required](#exactly-one-of-patch_config0windows_update0classifications-patch_config0windows_update0excludes-or-patch_config0windows_update0exclusive_patches-is-required)
    - [At least one of `patch_config.0.pre_step.0.linux_exec_step_config` or `patch_config.0.pre_step.0.windows_exec_step_config` is required](#at-least-one-of-patch_config0pre_step0linux_exec_step_config-or-patch_config0pre_step0windows_exec_step_config-is-required)
    - [At least one of `patch_config.0.post_step.0.linux_exec_step_config` or `patch_config.0.post_step.0.windows_exec_step_config` is required](#at-least-one-of-patch_config0post_step0linux_exec_step_config-or-patch_config0post_step0windows_exec_step_config-is-required)
  - [Resource: `google_spanner_instance`](#resource-google_spanner_instance)
    - [Exactly one of `num_nodes` or `processing_units` is required](#exactly-one-of-num_nodes-or-processing_units-is-required)

<!-- /TOC -->

# Terraform Google Provider 4.0.0 Upgrade Guide

The `4.0.0` release of the Google provider for Terraform is a major version and
includes some changes that you will need to consider when upgrading. This guide
is intended to help with that process and focuses only on the changes necessary
to upgrade from the final `3.X` series release to `4.0.0`.

Most of the changes outlined in this guide have been previously marked as
deprecated in the Terraform `plan`/`apply` output throughout previous provider
releases, up to and including the final `3.X` series release. These changes,
such as deprecation notices, can always be found in the CHANGELOG of the
affected providers. [google](https://github.com/hashicorp/terraform-provider-google/blob/master/CHANGELOG.md)
[google-beta](https://github.com/hashicorp/terraform-provider-google-beta/blob/master/CHANGELOG.md)

## I accidentally upgraded to 4.0.0, how do I downgrade to `3.X`?

If you've inadvertently upgraded to `4.0.0`, first see the
[Provider Version Configuration Guide](#provider-version-configuration) to lock
your provider version; if you've constrained the provider to a lower version
such as shown in the previous version example in that guide, Terraform will pull
in a `3.X` series release on `terraform init`.

If you've only ran `terraform init` or `terraform plan`, your state will not
have been modified and downgrading your provider is sufficient.

If you've ran `terraform refresh` or `terraform apply`, Terraform may have made
state changes in the meantime.

* If you're using a local state, or a remote state backend that does not support
versioning, `terraform refresh` with a downgraded provider is likely sufficient
to revert your state. The Google provider generally refreshes most state
information from the API, and the properties necessary to do so have been left
unchanged.

* If you're using a remote state backend that supports versioning such as
[Google Cloud Storage](https://www.terraform.io/docs/backends/types/gcs.html),
you can revert the Terraform state file to a previous version. If you do
so and Terraform had created resources as part of a `terraform apply` in the
meantime, you'll need to either delete them by hand or `terraform import` them
so Terraform knows to manage them.

## Provider Version Configuration

-> Before upgrading to version 4.0.0, it is recommended to upgrade to the most
recent `3.X` series release of the provider, make the changes noted in this guide,
and ensure that your environment successfully runs
[`terraform plan`](https://www.terraform.io/docs/commands/plan.html)
without unexpected changes or deprecation notices.

It is recommended to use [version constraints](https://www.terraform.io/docs/language/providers/requirements.html#requiring-providers)
when configuring Terraform providers. If you are following that recommendation,
update the version constraints in your Terraform configuration and run
[`terraform init`](https://www.terraform.io/docs/commands/init.html) to download
the new version.

If you aren't using version constraints, you can use `terraform init -upgrade`
in order to upgrade your provider to the latest released version.

For example, given this previous configuration:

```hcl
terraform {
  # ... other configuration ...
  required_providers {
    google = {
      version = "~> 3.87.0"
    }
  }
}
```

An updated configuration:

```hcl
terraform {
  # ... other configuration ...
  required_providers {
    google = {
      version = "~> 4.0.0"
    }
  }
}
```

## Provider

### Redundant default scopes are removed

Several default scopes are removed from the provider:

* "https://www.googleapis.com/auth/compute"
* "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
* "https://www.googleapis.com/auth/devstorage.full_control"
* "https://www.googleapis.com/auth/cloud-identity"

They are redundant with the "https://www.googleapis.com/auth/cloud-platform"
scope per [Access scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam).
After this change the following scopes are enabled, in line with `gcloud`'s
[list of scopes](https://cloud.google.com/sdk/gcloud/reference/auth/application-default/login):

* "https://www.googleapis.com/auth/cloud-platform"
* "https://www.googleapis.com/auth/userinfo.email"

This change is believed to have no user impact. If you find that Terraform
behaves incorrectly as a result of this change, please report a [bug](https://github.com/hashicorp/terraform-provider-google/issues/new?assignees=&labels=bug&template=bug.md).

### Runtime Configurator (`runtimeconfig`) resources have been removed from the GA provider

Earlier versions of the provider accidentally included the Runtime Configurator
service at GA. `4.0.0` has corrected that error, and Runtime Configurator is
only available in `google-beta`.

Affected Resources:

    * `google_runtimeconfig_config`
    * `google_runtimeconfig_variable`
    * `google_runtimeconfig_config_iam_policy`
    * `google_runtimeconfig_config_iam_binding`
    * `google_runtimeconfig_config_iam_member`

Affected Datasources:

    * `google_runtimeconfig_config`


If you have a configuration using the `google` provider like the following:

```
resource "google_runtimeconfig_config" "my-runtime-config" {
  name        = "my-service-runtime-config"
  description = "Runtime configuration values for my service"
}
```

Add the `google-beta` provider to your configuration:

```
resource "google_runtimeconfig_config" "my-runtime-config" {
  provider = google-beta

  name        = "my-service-runtime-config"
  description = "Runtime configuration values for my service"
}
```

### Service account scopes no longer accept `trace-append` or `trace-ro`, use `trace` instead

Previously users could specify `trace-append` or `trace-ro` as scopes for a given service account.
However, to better align with [Google documentation](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes), `trace` will now be the only valid scope, as it's an alias for `trace.append` and
`trace-ro` is no longer a documented option.

## Datasource: `google_product_resource`

### Datasource-level change example

Description of the change and how users should adjust their configuration (if needed).

## Resource: `google_bigquery_job`

### Exactly one of `query`, `load`, `copy` or `extract` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `query.0.script_options.0.statement_timeout_ms`, `query.0.script_options.0.statement_byte_budget`, or `query.0.script_options.0.key_result_statement` is required
The provider will now enforce at plan time that one of these fields be set.

### Exactly one of `extract.0.source_table` or `extract.0.source_model` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_cloudbuild_trigger`

### Exactly one of `build.0.source.0.repo_source.0.branch_name`, `build.0.source.0.repo_source.0.commit_sha` or `build.0.source.0.repo_source.0.tag_name` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_compute_autoscaler`

### At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas` or `autoscaling_policy.0.scale_down_control.0.time_window_sec` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed` or `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas` or `autoscaling_policy.0.scale_in_control.0.time_window_sec` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed` or `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_compute_region_autoscaler`

### At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas` or `autoscaling_policy.0.scale_down_control.0.time_window_sec` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.fixed` or `autoscaling_policy.0.scale_down_control.0.max_scaled_down_replicas.0.percent` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas` or `autoscaling_policy.0.scale_in_control.0.time_window_sec` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.fixed` or `autoscaling_policy.0.scale_in_control.0.max_scaled_in_replicas.0.percent` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_compute_firewall`

### One of `source_tags`, `source_ranges` or `source_service_accounts` are required on INGRESS firewalls

Previously, if all of these fields were left empty, the firewall defaulted to allowing traffic from 0.0.0.0/0, which is a suboptimal default.

## Resource: `google_compute_instance_group_manager`

### `update_policy.min_ready_sec` is removed from the GA provider
This field was incorrectly included in the GA `google` provider in past releases.
In order to continue to use the feature, add `provider = google-beta` to your
resource definition.

## Resource: `google_compute_region_instance_group_manager`

### `update_policy.min_ready_sec` is removed from the GA provider

This field was incorrectly included in the GA `google` provider in past releases.
In order to continue to use the feature, add `provider = google-beta` to your
resource definition.

## Resource: `google_compute_instance_template`

### `enable_display` is removed from the GA provider

This field was incorrectly included in the GA `google` provider in past releases.
In order to continue to use the feature, add `provider = google-beta` to your
resource definition.

## Resource: `google_compute_url_map`

### At least one of `default_route_action.0.fault_injection_policy.0.delay.0.fixed_delay` or `default_route_action.0.fault_injection_policy.0.delay.0.percentage` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_container_cluster`

### `instance_group_urls` is now removed

`instance_group_urls` has been removed in favor of `node_pool.instance_group_urls`

### `master_auth` is now removed

`master_auth` and its subfields have been removed. 
Basic authentication was removed for GKE cluster versions >= 1.19. The cluster cannot be created with basic authentication enabled. Instructions for choosing an alternative authentication method can be found at: cloud.google.com/kubernetes-engine/docs/how-to/api-server-authentication.

### `node_config.workload_metadata_config.node_metadata` is now removed

Removed in favor of `node_config.workload_metadata_config.mode`.

### `workload_identity_config.0.identity_namespace` is now removed

Removed in favor of `workload_identity_config.0.workload_pool`. Switching your
configuration from one value to the other will trigger a diff at plan time, and
a spurious update.

```diff
resource "google_container_cluster" "cluster" {
  name               = "your-cluster"
  location           = "us-central1-a"
  initial_node_count = 1

  workload_identity_config {
-    identity_namespace = "your-project.svc.id.goog"
+   workload_pool = "your-project.svc.id.goog"
  }
```

### `pod_security_policy_config` is removed from the GA provider

This field was incorrectly included in the GA `google` provider in past releases.
In order to continue to use the feature, add `provider = google-beta` to your
resource definition.

## Resource: `google_data_loss_prevention_trigger`

### Exactly one of `inspect_job.0.storage_config.0.cloud_storage_options.0.file_set.0.url` or `inspect_job.0.storage_config.0.cloud_storage_options.0.file_set.0.regex_file_set` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `inspect_job.0.storage_config.0.timespan_config.0.start_time` or `inspect_job.0.storage_config.0.timespan_config.0.end_time` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_os_config_patch_deployment`

### At least one of `patch_config.0.reboot_config`, `patch_config.0.apt`, `patch_config.0.yum`, `patch_config.0.goo` `patch_config.0.zypper`, `patch_config.0.windows_update`, `patch_config.0.pre_step` or `patch_config.0.pre_step` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `patch_config.0.apt.0.type`, `patch_config.0.apt.0.excludes` or `patch_config.0.apt.0.exclusive_packages` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `patch_config.0.yum.0.security`, `patch_config.0.yum.0.minimal`, `patch_config.0.yum.0.excludes` or `patch_config.0.yum.0.exclusive_packages` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `patch_config.0.zypper.0.with_optional`, `patch_config.0.zypper.0.with_update`, `patch_config.0.zypper.0.categories`, `patch_config.0.zypper.0.severities`, `patch_config.0.zypper.0.excludes` or `patch_config.0.zypper.0.exclusive_patches` is required
The provider will now enforce at plan time that one of these fields be set.

### Exactly one of `patch_config.0.windows_update.0.classifications`, `patch_config.0.windows_update.0.excludes` or `patch_config.0.windows_update.0.exclusive_patches` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `patch_config.0.pre_step.0.linux_exec_step_config` or `patch_config.0.pre_step.0.windows_exec_step_config` is required
The provider will now enforce at plan time that one of these fields be set.

### At least one of `patch_config.0.post_step.0.linux_exec_step_config` or `patch_config.0.post_step.0.windows_exec_step_config` is required
The provider will now enforce at plan time that one of these fields be set.

## Resource: `google_project`

### `org_id`, `folder_id` now conflict at plan time

Previously, they were only checked for conflicts at apply time. Terraform will
now report an error at plan time.

### `org_id`, `folder_id` are unset when removed from config

Previously, these fields kept their old value in state when they were removed
from config, changing the value on next refresh. Going forward, removing one of
the values or switching values will generate a correct plan that removes the
value.

## Resource: `google_project_service`

### `bigquery-json.googleapis.com` is no longer a valid service name

`bigquery-json.googleapis.com` was deprecated in the `3.0.0` release, however, at that point the provider
converted it while the upstream API migration was in progress. Now that the API migration has finished,
the provider will no longer convert the service name. Use `bigquery.googleapis.com` instead.

## Resource: `google_spanner_instance`

### Exactly one of `num_nodes` or `processing_units` is required
The provider will now enforce at plan time that one of these fields be set.
