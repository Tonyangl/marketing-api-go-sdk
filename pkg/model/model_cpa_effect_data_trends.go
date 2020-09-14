/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 成本趋势
type CpaEffectDataTrends struct {
	Rank      int64     `json:"rank,omitempty"`
	TargetCpa *[]string `json:"target_cpa,omitempty"`
	RealCpa   *[]string `json:"real_cpa,omitempty"`
}
