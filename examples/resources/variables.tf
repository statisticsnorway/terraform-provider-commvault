variable "commvault_base_url" {
  description = "Base URL for the Commvault API"
  type        = string
}

variable "commvault_username" {
  description = "Username for Commvault authentication"
  type        = string
}

variable "commvault_password" {
  description = "Password for Commvault authentication"
  type        = string
  sensitive   = true
}

variable "client_name" {
  description = "Name of the Commvault client"
  type        = string
}

variable "plan_id" {
  description = "Plan ID to associate with the client"
  type        = number
}

variable "credential_id" {
  description = "Credential ID for the Commvault client"
  type        = number
}

variable "access_node_id" {
  description = "Access Node ID for the Commvault client"
  type        = number
}

variable "project_id" {
  description = "GCP Project ID"
  type        = string
}