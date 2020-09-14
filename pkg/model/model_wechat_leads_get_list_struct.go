/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type WechatLeadsGetListStruct struct {
	AdgroupId    int64              `json:"adgroup_id,omitempty"`
	LeadsInfo    *[]LeadsInfoStruct `json:"leads_info,omitempty"`
	CampaignId   int64              `json:"campaign_id,omitempty"`
	CampaignName string             `json:"campaign_name,omitempty"`
	AdgroupName  string             `json:"adgroup_name,omitempty"`
	AgencyId     string             `json:"agency_id,omitempty"`
	AgencyName   string             `json:"agency_name,omitempty"`
	ClickId      string             `json:"click_id,omitempty"`
}
