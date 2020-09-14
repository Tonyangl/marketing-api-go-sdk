/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 场景定向
type SceneTargeting struct {
	MobileUnion                 *[]string `json:"mobile_union,omitempty"`
	ExcludeMobileUnion          *[]string `json:"exclude_mobile_union,omitempty"`
	MobileUnionIndustry         *[]string `json:"mobile_union_industry,omitempty"`
	UnionPositionPackage        *[]int64  `json:"union_position_package,omitempty"`
	ExcludeUnionPositionPackage *[]int64  `json:"exclude_union_position_package,omitempty"`
	TencentNews                 *[]string `json:"tencent_news,omitempty"`
	DisplayScene                *[]string `json:"display_scene,omitempty"`
}
