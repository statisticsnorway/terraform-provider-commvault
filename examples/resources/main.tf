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

resource "commvault_client" "gcp_client" {
  // Per team
  name    = var.client_name
  plan_id = var.plan_id

  credential_id  = var.credential_id
  access_node_id = var.access_node_id
  project_id     = var.project_id

}

# data "commvault_subclient" "default_gcp_subclient" {
#   client_id = commvault_client.gcp_client.id
# }
#
resource "commvault_subclient" "subclient" {
  //Per miljø // kilde/ikke-kildedata -> plan
  // id - Computed
  name      = var.client_name
  app_name  = ""
  client_id = commvault_client.gcp_client.id

  use_local_content = true


  gcp_contents = [
    {
      bucket_name  = "bucket1"
      project_id = ""
    }
  ]
}
#
#
# resource "commvault_plan" "plan" {
#   // ?
#   subclient = commvault_subclient.subclient.client_id
#   credential_id  = var.credential_id
#   access_node_id = var.access_node_id
#   project_id     = var.project_id
# }

output "created_client_id" {
  value = commvault_client.gcp_client.id
}

output "client_response" {
  value = commvault_client.gcp_client.response
}
