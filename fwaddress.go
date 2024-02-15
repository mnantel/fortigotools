package main

import (
	"encoding/json"
	"fmt"
)

type FirewallAddress struct {
	Name                string `json:"name"`
	QOriginKey          string `json:"q_origin_key"`
	UUID                string `json:"uuid"`
	Subnet              string `json:"subnet"`
	Type                string `json:"type"`
	RouteTag            int    `json:"route-tag"`
	SubType             string `json:"sub-type"`
	ClearpassSpt        string `json:"clearpass-spt"`
	Macaddr             []any  `json:"macaddr"`
	Country             string `json:"country"`
	CacheTTL            int    `json:"cache-ttl"`
	Sdn                 string `json:"sdn"`
	FssoGroup           []any  `json:"fsso-group"`
	Interface           string `json:"interface"`
	ObjType             string `json:"obj-type"`
	TagDetectionLevel   string `json:"tag-detection-level"`
	TagType             string `json:"tag-type"`
	Dirty               string `json:"dirty"`
	HwVendor            string `json:"hw-vendor"`
	HwModel             string `json:"hw-model"`
	Os                  string `json:"os"`
	SwVersion           string `json:"sw-version"`
	Comment             string `json:"comment"`
	AssociatedInterface string `json:"associated-interface"`
	Color               int    `json:"color"`
	Filter              string `json:"filter"`
	SdnAddrType         string `json:"sdn-addr-type"`
	NodeIPOnly          string `json:"node-ip-only"`
	ObjID               string `json:"obj-id"`
	List                []any  `json:"list"`
	Tagging             []any  `json:"tagging"`
	AllowRouting        string `json:"allow-routing"`
	FabricObject        string `json:"fabric-object"`
}

func (fw *Fos) GetFirewallAddress() ([]FirewallAddress, error) {

	body, err := fw.MakeApiCall("GET", "/api/v2/cmdb/firewall/address", "")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the ApiResponse struct
	type ApiResponse struct {
		Results []FirewallAddress `json:"results"`
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}
	return apiResponse.Results, nil
}
