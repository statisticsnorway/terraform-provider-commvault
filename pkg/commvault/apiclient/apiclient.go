package apiclient

import (
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient/apiexplorer"
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient/sp36"
)

type CommvaultApiClient struct {
	Sp36ApiClient        *sp36.APIClient
	ApiExplorerApiClient *apiexplorer.APIClient
}
