terraform {
  backend "gcs" {} # Intentionally empty, to be filled by terragrunt
}

resource "random_string" "suffix" {
  length  = 3
  special = false
  upper   = false
}

module "project-factory" {
  source  = "terraform-google-modules/project-factory/google"
  version = "~> 10.1"

  name                 = "us-dev-gcp-free-tier"
  random_project_id    = true
  org_id               = var.org_id
  billing_account      = var.billing_account

   activate_apis = [
    "compute.googleapis.com",
    "iam.googleapis.com",
    "container.googleapis.com",
  ]
}