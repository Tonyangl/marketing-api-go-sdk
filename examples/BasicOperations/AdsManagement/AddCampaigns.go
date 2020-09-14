/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type CampaignsAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.CampaignsAddRequest
}

func (e *CampaignsAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.CampaignsAddRequest{
		CampaignName:       "SDK计划5ede2529308ad",
		AccountId:          int64(0),
		CampaignType:       model.CampaignType_NORMAL,
		PromotedObjectType: model.PromotedObjectType_LINK,
		DailyBudget:        int64(0),
	}
}

func (e *CampaignsAddExample) RunExample() (model.CampaignsAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Campaigns().Add(ctx, e.Data)
}

func main() {
	e := &CampaignsAddExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
