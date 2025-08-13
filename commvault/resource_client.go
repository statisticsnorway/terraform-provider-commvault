package commvault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCommvaultClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceCommvaultClientCreate,
		Read:   resourceCommvaultClientRead,
		Delete: resourceCommvaultClientDelete,

		Schema: map[string]*schema.Schema{
			"id":      {Type: schema.TypeString, Computed: true},
			"name":    {Type: schema.TypeString, Required: true, ForceNew: true},
			"plan_id": {Type: schema.TypeInt, Required: true, ForceNew: true},
			"credential_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"access_node_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"response": {Type: schema.TypeString, Computed: true},
		},
	}
}

func resourceCommvaultClientCreate(d *schema.ResourceData, m interface{}) error {
	cfg := m.(*CommvaultConfig)

	token, err := login(cfg)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	planID := d.Get("plan_id").(int)
	credentialID := d.Get("credential_id").(int)
	accessNodeID := d.Get("access_node_id").(int)
	projectID := d.Get("project_id").(string)

	payload := map[string]interface{}{
		"clientInfo": map[string]interface{}{
			"clientType": 15,
			"cloudClonnectorProperties": map[string]interface{}{
				"instance": map[string]interface{}{
					"cloudAppsInstance": map[string]interface{}{
						"credentialType": "GOOGLE_SERVICE_ACCOUNT",
						"generalCloudProperties": map[string]interface{}{
							"credentials": map[string]interface{}{
								"credentialId": credentialID,
							},
							"memberServers": []map[string]interface{}{
								{
									"client": map[string]interface{}{
										"_type_":   3,
										"clientId": accessNodeID,
									},
								},
							},
							"numberOfBackupStreams": 1,
						},
						"googleCloudInstance": map[string]interface{}{
							"GCPProjectId": projectID,
							"serverName":   "storage.googleapis.com",
						},
						"instanceType":            "GOOGLE_CLOUD",
						"instanceTypeDisplayName": "Google Cloud",
					},
					"instance": map[string]interface{}{
						"applicationId": 134,
						"commCellId":    2,
						"instanceName":  name,
					},
				},
				"instanceType": "GOOGLE_CLOUD",
			},
			"plan": map[string]interface{}{"planId": planID},
		},
		"entity": map[string]interface{}{"clientName": name},
	}

	data, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", cfg.BaseURL+"/Client", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authtoken", token)

	client := newHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("create failed: non-2xx response (%d)", resp.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	var cr createResponse
	if err := json.Unmarshal(bodyBytes, &cr); err != nil {
		return fmt.Errorf("create failed: could not parse response")
	}
	if cr.Response.Entity.ClientID == 0 {
		return fmt.Errorf("create failed: missing clientId")
	}

	minimalResp := map[string]interface{}{
		"clientId":   cr.Response.Entity.ClientID,
		"clientName": cr.Response.Entity.ClientName,
	}
	minimalBytes, _ := json.Marshal(minimalResp)
	d.Set("response", string(minimalBytes))
	d.SetId(fmt.Sprintf("%d", cr.Response.Entity.ClientID))
	return nil
}

func resourceCommvaultClientRead(d *schema.ResourceData, m interface{}) error {
	cfg := m.(*CommvaultConfig)

	token, err := login(cfg)
	if err != nil {
		return err
	}

	clientID := d.Id()
	if _, err := strconv.Atoi(clientID); err != nil {
		d.SetId("")
		return nil
	}

	req, _ := http.NewRequest("GET", cfg.BaseURL+"/Client/"+clientID, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authtoken", token)

	client := newHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("read failed: non-2xx response (%d)", resp.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	var parsed map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &parsed); err == nil {
		if props, ok := parsed["clientProperties"].([]interface{}); ok && len(props) > 0 {
			if first, ok := props[0].(map[string]interface{}); ok {
				if clientInfo, ok := first["client"].(map[string]interface{}); ok {
					if entity, ok := clientInfo["clientEntity"].(map[string]interface{}); ok {
						minimalResp := map[string]interface{}{
							"clientId":   entity["clientId"],
							"clientName": entity["clientName"],
						}
						minimalBytes, _ := json.Marshal(minimalResp)
						d.Set("response", string(minimalBytes))
					}
				}
			}
		}
	}
	return nil
}

func resourceCommvaultClientDelete(d *schema.ResourceData, m interface{}) error {
	cfg := m.(*CommvaultConfig)

	token, err := login(cfg)
	if err != nil {
		return fmt.Errorf("delete failed: login error")
	}

	clientID := d.Id()
	if clientID == "" {
		return fmt.Errorf("delete failed: missing client ID")
	}

	deleteURL := fmt.Sprintf("%s/Client/%s?forceDelete=true", cfg.BaseURL, clientID)
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return fmt.Errorf("delete failed: build request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authtoken", token)

	client := newHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("delete failed: request error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("delete failed: non-2xx response (%d)", resp.StatusCode)
	}

	d.SetId("")
	return nil
}
