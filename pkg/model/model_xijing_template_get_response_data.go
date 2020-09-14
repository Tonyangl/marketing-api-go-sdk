/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type XijingTemplateGetResponseData struct {
	PageTemplateId    string            `json:"page_template_id,omitempty"`
	PageType          TemplatesPageType `json:"page_type,omitempty"`
	PageName          string            `json:"page_name,omitempty"`
	PageTitle         string            `json:"page_title,omitempty"`
	ComponentSpecList *[]string         `json:"component_spec_list,omitempty"`
	MobileAppId       string            `json:"mobile_app_id,omitempty"`
}
