package google

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProjectServiceServiceValidateFunc(t *testing.T) {
	cases := map[string]struct {
		val                   interface{}
		ExpectValidationError bool
	}{
		"ignoredProjectService": {
			val:                   "dataproc-control.googleapis.com",
			ExpectValidationError: true,
		},
		"bannedProjectService": {
			val:                   "bigquery-json.googleapis.com",
			ExpectValidationError: true,
		},
		"third party API": {
			val:                   "whatever.example.com",
			ExpectValidationError: false,
		},
		"not a domain": {
			val:                   "monitoring",
			ExpectValidationError: true,
		},
		"not a string": {
			val:                   5,
			ExpectValidationError: true,
		},
	}

	for tn, tc := range cases {
		_, errs := validateProjectServiceService(tc.val, "service")
		if tc.ExpectValidationError && len(errs) == 0 {
			t.Errorf("bad: %s, %q passed validation but was expected to fail", tn, tc.val)
		}
		if !tc.ExpectValidationError && len(errs) > 0 {
			t.Errorf("bad: %s, %q failed validation but was expected to pass. errs: %q", tn, tc.val, errs)
		}
	}
}

// Test that services can be enabled and disabled on a project
func TestAccProjectService_basic(t *testing.T) {
	t.Parallel()
	// Multiple fine-grained resources
	skipIfVcr(t)

	org := getTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", randInt(t))
	services := []string{"iam.googleapis.com", "cloudresourcemanager.googleapis.com"}
	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectService_basic(services, pid, pname, org),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectService(t, services, pid, true),
				),
			},
			{
				ResourceName:            "google_project_service.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disable_on_destroy"},
			},
			{
				ResourceName:            "google_project_service.test2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disable_on_destroy", "project"},
			},
			// Use a separate TestStep rather than a CheckDestroy because we need the project to still exist.
			{
				Config: testAccProject_create(pid, pname, org),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectService(t, services, pid, false),
				),
			},
			// Create services with disabling turned off.
			{
				Config: testAccProjectService_noDisable(services, pid, pname, org),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectService(t, services, pid, true),
				),
			},
			// Check that services are still enabled even after the resources are deleted.
			{
				Config: testAccProject_create(pid, pname, org),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectService(t, services, pid, true),
				),
			},
		},
	})
}

func TestAccProjectService_disableDependentServices(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	org := getTestOrgFromEnv(t)
	billingId := getTestBillingAccountFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", randInt(t))
	services := []string{"cloudbuild.googleapis.com", "containerregistry.googleapis.com"}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectService_disableDependentServices(services, pid, pname, org, billingId, "false"),
			},
			{
				ResourceName:            "google_project_service.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disable_on_destroy"},
			},
			{
				Config:      testAccProjectService_dependencyRemoved(services, pid, pname, org, billingId),
				ExpectError: regexp.MustCompile("Please specify disable_dependent_services=true if you want to proceed with disabling all services."),
			},
			{
				Config: testAccProjectService_disableDependentServices(services, pid, pname, org, billingId, "true"),
			},
			{
				ResourceName:            "google_project_service.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disable_on_destroy"},
			},
			{
				Config:      testAccProjectService_dependencyRemoved(services, pid, pname, org, billingId),
				ExpectError: regexp.MustCompile("service .* not in enabled services for project"),
			},
		},
	})
}

func TestAccProjectService_handleNotFound(t *testing.T) {
	t.Parallel()

	org := getTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", randInt(t))
	service := "iam.googleapis.com"
	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectService_handleNotFound(service, pid, pname, org),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectService(t, []string{service}, pid, true),
				),
			},
			// Delete the project, implicitly deletes service, expect the plan to want to create the service again
			{
				Config:             testAccProjectService_handleNotFoundNoProject(service, pid),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccProjectService_renamedService(t *testing.T) {
	t.Parallel()

	if len(renamedServices) == 0 {
		t.Skip()
	}

	var newName string
	for _, new := range renamedServices {
		newName = new
	}

	org := getTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", randInt(t))
	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectService_single(newName, pid, pname, org),
			},
			{
				ResourceName:            "google_project_service.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disable_on_destroy", "disable_dependent_services"},
			},
		},
	})
}

func testAccCheckProjectService(t *testing.T, services []string, pid string, expectEnabled bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)

		currentlyEnabled, err := listCurrentlyEnabledServices(pid, config.userAgent, config, time.Minute*10)
		if err != nil {
			return fmt.Errorf("Error listing services for project %q: %v", pid, err)
		}

		for _, expected := range services {
			exists := false
			for actual := range currentlyEnabled {
				if expected == actual {
					exists = true
				}
			}
			if expectEnabled && !exists {
				return fmt.Errorf("Expected service %s is not enabled server-side", expected)
			}
			if !expectEnabled && exists {
				return fmt.Errorf("Expected disabled service %s is enabled server-side", expected)
			}
		}

		return nil
	}
}

func testAccProjectService_basic(services []string, pid, name, org string) string {
	return fmt.Sprintf(`
provider "google" {
  user_project_override = true
}
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
}

resource "google_project_service" "test" {
  project = google_project.acceptance.project_id
  service = "%s"
}

resource "google_project_service" "test2" {
  project = google_project.acceptance.id
  service = "%s"
}
`, pid, name, org, services[0], services[1])
}

func testAccProjectService_disableDependentServices(services []string, pid, name, org, billing, disableDependentServices string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id      = "%s"
  name            = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "test" {
  project = google_project.acceptance.project_id
  service = "%s"
}

resource "google_project_service" "test2" {
  project                    = google_project.acceptance.project_id
  service                    = "%s"
  disable_dependent_services = %s
}
`, pid, name, org, billing, services[0], services[1], disableDependentServices)
}

func testAccProjectService_dependencyRemoved(services []string, pid, name, org, billing string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id      = "%s"
  name            = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "test" {
  project = google_project.acceptance.project_id
  service = "%s"
}
`, pid, name, org, billing, services[0])
}

func testAccProjectService_noDisable(services []string, pid, name, org string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
}

resource "google_project_service" "test" {
  project            = google_project.acceptance.project_id
  service            = "%s"
  disable_on_destroy = false
}

resource "google_project_service" "test2" {
  project            = google_project.acceptance.project_id
  service            = "%s"
  disable_on_destroy = false
}
`, pid, name, org, services[0], services[1])
}

func testAccProjectService_handleNotFound(service, pid, name, org string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
}

// by passing through locals, we break the dependency chain
// see terraform-provider-google#1292
locals {
  project_id = google_project.acceptance.project_id
}

resource "google_project_service" "test" {
  project = local.project_id
  service = "%s"
}
`, pid, name, org, service)
}

func testAccProjectService_handleNotFoundNoProject(service, pid string) string {
	return fmt.Sprintf(`
resource "google_project_service" "test" {
  project = "%s"
  service = "%s"
}
`, pid, service)
}

func testAccProjectService_single(service string, pid, name, org string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
}

resource "google_project_service" "test" {
  project = google_project.acceptance.project_id
  service = "%s"

  disable_dependent_services = true
}
`, pid, name, org, service)
}
