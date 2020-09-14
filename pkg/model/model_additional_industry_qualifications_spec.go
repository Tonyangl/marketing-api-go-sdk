/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 附加行业资质信息，当且仅当qualification_type=ADDITIONAL_INDUSTRY_QUALIFICATION时可填且必填
type AdditionalIndustryQualificationsSpec struct {
	SystemIndustryId  int64     `json:"system_industry_id,omitempty"`
	BusinessScopeId   int64     `json:"business_scope_id,omitempty"`
	QualificationCode string    `json:"qualification_code,omitempty"`
	ImageIdList       *[]string `json:"image_id_list,omitempty"`
}
