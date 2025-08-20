terraform {
  required_providers {
    commvaultx = {
      source  = "statisticsnorway/commvaultx"
      version = "0.1.0"
    }
  }
}

provider "commvaultx" {
  base_url = var.commvault_base_url
  username = var.commvault_username
  password = var.commvault_password
}

resource "commvaultx_client" "gcp_client" {
  name           = var.client_name
  plan_id        = var.plan_id
  credential_id  = var.credential_id
  access_node_id = var.access_node_id
  project_id     = var.project_id
}

output "created_client_id" {
  value = commvaultx_client.gcp_client.id
}

output "client_response" {
  value = commvaultx_client.gcp_client.response
}