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
type LeadsGetListStruct struct {
	CampaignId         int64              `json:"campaign_id,omitempty"`
	CampaignName       string             `json:"campaign_name,omitempty"`
	AdgroupId          int64              `json:"adgroup_id,omitempty"`
	AdgroupName        string             `json:"adgroup_name,omitempty"`
	WechatAdgroupId    int64              `json:"wechat_adgroup_id,omitempty"`
	LeadSpecList       *[]LeadsInfoStruct `json:"lead_spec_list,omitempty"`
	WechatCampaignId   int64              `json:"wechat_campaign_id,omitempty"`
	WechatCampaignName string             `json:"wechat_campaign_name,omitempty"`
	WechatAdgroupName  string             `json:"wechat_adgroup_name,omitempty"`
	WechatAgencyId     string             `json:"wechat_agency_id,omitempty"`
	WechatAgencyName   string             `json:"wechat_agency_name,omitempty"`
	ClickId            string             `json:"click_id,omitempty"`
}
