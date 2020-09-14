/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// DeepConversionWorthGoal : 优化ROI目标，不可更改，如修改深度优化效果值，需在结构体中传入已设置的优化ROI目标
type DeepConversionWorthGoal string

// List of DeepConversionWorthGoal
const (
	DeepConversionWorthGoal_30DAYPURCHASEROAS     DeepConversionWorthGoal = "GOAL_30DAY_PURCHASE_ROAS"
	DeepConversionWorthGoal_30DAYMONETIZATIONROAS DeepConversionWorthGoal = "GOAL_30DAY_MONETIZATION_ROAS"
	DeepConversionWorthGoal_30DAYORDERROAS        DeepConversionWorthGoal = "GOAL_30DAY_ORDER_ROAS"
	DeepConversionWorthGoal_1DAYPURCHASEROAS      DeepConversionWorthGoal = "GOAL_1DAY_PURCHASE_ROAS"
	DeepConversionWorthGoal_1DAYMONETIZATIONROAS  DeepConversionWorthGoal = "GOAL_1DAY_MONETIZATION_ROAS"
)
