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

data "commvault_client" client {
  client_name = "gcp-commvault-poc-36"
  # client_id = "3431"
}


output "client_name" {
  value = data.commvault_client.client.client_name
}

output "client_id" {
  value = data.commvault_client.client.client_id
}
