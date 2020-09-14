/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 行为上报数据
type UserAction struct {
	ActionTime       int64                  `json:"action_time,omitempty"`
	UserId           *ActionsUserId         `json:"user_id,omitempty"`
	ActionType       ActionType             `json:"action_type,omitempty"`
	ActionParam      map[string]interface{} `json:"action_param,omitempty"`
	CustomAction     string                 `json:"custom_action,omitempty"`
	Trace            *Trace                 `json:"trace,omitempty"`
	Url              string                 `json:"url,omitempty"`
	ExternalActionId string                 `json:"external_action_id,omitempty"`
}
