terraform {
  required_providers {
    commvault = {
      source  = "statisticsnorway/commvault"
      version = "0.1.0"
    }
  }
}

provider "commvault" {
  base_url = var.commvault_base_url
  username = var.commvault_username
  password = var.commvault_password
}

resource "commvault_client" "gcp" {
  name           = var.client_name
  plan_id        = var.plan_id
  credential_id  = var.credential_id
  access_node_id = var.access_node_id
  project_id     = var.project_id

  dynamic "bucket_contents" {
    for_each = var.bucket_contents
    content {
      name    = bucket_contents.value.name
      project = coalesce(try(bucket_contents.value.project, null), var.project_id)
    }
  }
}

output "client_id" {
  value = commvault_client.gcp.id
}