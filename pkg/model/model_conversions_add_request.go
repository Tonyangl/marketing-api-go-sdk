/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ConversionsAddRequest struct {
	AccountId                    *int64                     `json:"account_id,omitempty"`
	ConversionName               *string                    `json:"conversion_name,omitempty"`
	AccessType                   AccessType                 `json:"access_type,omitempty"`
	ConversionScene              ConversionScene            `json:"conversion_scene,omitempty"`
	PromotedObjectId             *string                    `json:"promoted_object_id,omitempty"`
	AppAndroidChannelPackageId   *string                    `json:"app_android_channel_package_id,omitempty"`
	ClaimType                    ClaimType                  `json:"claim_type,omitempty"`
	FeedbackUrl                  *string                    `json:"feedback_url,omitempty"`
	LandingPageUrl               *string                    `json:"landing_page_url,omitempty"`
	MiniProgramId                *string                    `json:"mini_program_id,omitempty"`
	SelfAttributed               *bool                      `json:"self_attributed,omitempty"`
	OptimizationGoal             IntOptimizationGoal        `json:"optimization_goal,omitempty"`
	DeepBehaviorOptimizationGoal IntOptimizationGoal        `json:"deep_behavior_optimization_goal,omitempty"`
	DeepWorthOptimizationGoal    ConversionOptimizationGoal `json:"deep_worth_optimization_goal,omitempty"`
	DeepWorthAdvancedGoal        ConversionOptimizationGoal `json:"deep_worth_advanced_goal,omitempty"`
	UserActionSetId              *int64                     `json:"user_action_set_id,omitempty"`
	ConversionLinkId             *int64                     `json:"conversion_link_id,omitempty"`
	ImpressionFeedbackUrl        *string                    `json:"impression_feedback_url,omitempty"`
	AttributionWindow            *int64                     `json:"attribution_window,omitempty"`
	DeepBehaviorAdvancedGoal     IntOptimizationGoal        `json:"deep_behavior_advanced_goal,omitempty"`
	DeepOptimizationGoalType     DeepOptimizationGoalType   `json:"deep_optimization_goal_type,omitempty"`
}
