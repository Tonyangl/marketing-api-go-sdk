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
type ConversionsGetListStruct struct {
	ConversionId                 *int64                     `json:"conversion_id,omitempty"`
	ConversionName               *string                    `json:"conversion_name,omitempty"`
	AccessType                   AccessType                 `json:"access_type,omitempty"`
	ClaimType                    ClaimType                  `json:"claim_type,omitempty"`
	FeedbackUrl                  *string                    `json:"feedback_url,omitempty"`
	SelfAttributed               *bool                      `json:"self_attributed,omitempty"`
	OptimizationGoal             IntOptimizationGoal        `json:"optimization_goal,omitempty"`
	DeepBehaviorOptimizationGoal IntOptimizationGoal        `json:"deep_behavior_optimization_goal,omitempty"`
	DeepWorthOptimizationGoal    ConversionOptimizationGoal `json:"deep_worth_optimization_goal,omitempty"`
	UserActionSetId              *int64                     `json:"user_action_set_id,omitempty"`
	UserActionSetKey             *string                    `json:"user_action_set_key,omitempty"`
	SiteSetEnable                *bool                      `json:"site_set_enable,omitempty"`
	IsDeleted                    *bool                      `json:"is_deleted,omitempty"`
	AccessStatus                 AccessStatus               `json:"access_status,omitempty"`
	CreateSourceType             CreateSourceType           `json:"create_source_type,omitempty"`
}
