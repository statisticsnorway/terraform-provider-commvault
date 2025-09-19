variable "commvault_base_url" {
  type        = string
  description = "Base API URL, e.g. https://<host>/commandcenter/api"
}

variable "commvault_username" {
  type        = string
  description = "Commvault username"
}

variable "commvault_password" {
  type        = string
  sensitive   = true
  description = "Commvault password (sensitive)"
}

variable "client_name" {
  type        = string
  description = "Commvault client name to create"
}

variable "plan_id" {
  type        = number
  description = "Plan ID to attach to the client"
}

variable "credential_id" {
  type        = number
  description = "Credential ID for the client"
}

variable "access_node_id" {
  type        = number
  description = "Access node (client) ID"
}

variable "project_id" {
  type        = string
  description = "Default GCP Project ID (used when a bucket item has no project)"
}

variable "bucket_contents" {
  description = "List of buckets to include; project_id defaults to project_id when omitted."
  type = list(object({
    name       = string
    project_id = optional(string)
  }))
}