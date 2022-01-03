remote_state {
  backend = "gcs"
  config = {
    bucket = "data-catalog-demo-323514-tfstate"
    prefix = "gcp-free-tier/${path_relative_to_include()}"
    
    #key = "${path_relative_to_include()}/terraform.tfstate"
    #region         = "us-east-1"
  }
}