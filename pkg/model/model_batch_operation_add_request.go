/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type BatchOperationAddRequest struct {
	ResourceType   ResourceType       `json:"resource_type,omitempty"`
	OperationType  BatchOperationType `json:"operation_type,omitempty"`
	ResourceIdList *[]int64           `json:"resource_id_list,omitempty"`
	OperationSpec  *OperationSpec     `json:"operation_spec,omitempty"`
	AccountId      int64              `json:"account_id,omitempty"`
}
