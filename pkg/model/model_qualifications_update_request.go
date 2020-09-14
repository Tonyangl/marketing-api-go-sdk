/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type QualificationsUpdateRequest struct {
	QualificationType QualificationType `json:"qualification_type,omitempty"`
	QualificationId   int64             `json:"qualification_id,omitempty"`
	ImageIdList       *[]string         `json:"image_id_list,omitempty"`
	AccountId         int64             `json:"account_id,omitempty"`
}
