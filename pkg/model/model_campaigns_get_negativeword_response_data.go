/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CampaignsGetNegativewordResponseData struct {
	CampaignErrorList *[]string                     `json:"campaign_error_list,omitempty"`
	CampaignList      *[]NegativeWordCampaignStruct `json:"campaign_list,omitempty"`
}