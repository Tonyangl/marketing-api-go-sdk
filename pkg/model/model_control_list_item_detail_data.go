/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 表单项
type ControlListItemDetailData struct {
	ItemName       string   `json:"item_name,omitempty"`
	ItemType       ItemType `json:"item_type,omitempty"`
	Placeholder    string   `json:"placeholder,omitempty"`
	SelectElements string   `json:"select_elements,omitempty"`
}
