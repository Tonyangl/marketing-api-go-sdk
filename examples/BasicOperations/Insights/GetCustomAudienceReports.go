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

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type CustomAudienceReportsGetExample struct {
	TAds                         *ads.SDKClient
	AccessToken                  string
	AccountId                    int64
	Filtering                    []model.FilteringStruct
	DateRange                    model.DateRange
	CustomAudienceReportsGetOpts *api.CustomAudienceReportsGetOpts
}

func (e *CustomAudienceReportsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Filtering = []model.FilteringStruct{{
		Field:    "audience_id",
		Operator: "IN",
		Values:   &[]string{"YOUR AUDIENCE ID"},
	}, {
		Field:    "audience_platform",
		Operator: "EQUALS",
		Values:   &[]string{"DMP"},
	}}
	e.DateRange = model.DateRange{
		StartDate: "REPORT START DATE",
		EndDate:   "REPORT END DATE",
	}
	e.CustomAudienceReportsGetOpts = &api.CustomAudienceReportsGetOpts{

		Fields: optional.NewInterface([]string{"audience_id", "account_id", "adgroup_id", "campaign_id", "wechat_adgroup_id", "wechat_campaign_id", "model_id", "audience_predict_task_id", "action_type", "cost", "action_count", "user_count"}),
	}
}

func (e *CustomAudienceReportsGetExample) RunExample() (model.CustomAudienceReportsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomAudienceReports().Get(ctx, e.AccountId, e.Filtering, e.DateRange, e.CustomAudienceReportsGetOpts)
}

func main() {
	e := &CustomAudienceReportsGetExample{}
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
