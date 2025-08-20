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