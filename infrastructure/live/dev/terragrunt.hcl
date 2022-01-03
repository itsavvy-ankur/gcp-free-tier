terraform {
  # Deploy version v0.0.1 in prod
  source = "../../modules//project"
}

include {
  path = find_in_parent_folders()
}

inputs = {
  org_id = "77598684511"
  billing_account  = "0159E4-888CC5-C3E41C"
}