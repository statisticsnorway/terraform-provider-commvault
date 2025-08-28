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

