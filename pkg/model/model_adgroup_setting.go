/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告组信息所组成的对象
type AdgroupSetting struct {
	AdgroupId            int64              `json:"adgroup_id,omitempty"`
	AdgroupName          string             `json:"adgroup_name,omitempty"`
	AutomaticSiteEnabled bool               `json:"automatic_site_enabled,omitempty"`
	SiteSet              *[]string          `json:"site_set,omitempty"`
	BidType              CostType           `json:"bid_type,omitempty"`
	BidAmount            int64              `json:"bid_amount,omitempty"`
	PromotedObjectType   PromotedObjectType `json:"promoted_object_type,omitempty"`
	BillingEvent         BillingEvent       `json:"billing_event,omitempty"`
	OptimizationGoal     OptimizationGoal   `json:"optimization_goal,omitempty"`
	PromotedObjectId     string             `json:"promoted_object_id,omitempty"`
	TimeSeries           string             `json:"time_series,omitempty"`
}
