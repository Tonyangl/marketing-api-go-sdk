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

type TargetingsAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.TargetingsAddRequest
}

func (e *TargetingsAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.TargetingsAddRequest{
		Targeting: &model.WriteTargetingSetting{
			Age: &[]model.AgeStruct{{
				Max: int64(0),
				Min: int64(0),
			}},
			Gender: &[]string{"YOUR TARGETING GENDER"},
		},
		AccountId:     int64(0),
		TargetingName: "SDK定向5ede252947141",
		Description:   "SDK 测试定向",
	}
}

func (e *TargetingsAddExample) RunExample() (model.TargetingsAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Targetings().Add(ctx, e.Data)
}

func main() {
	e := &TargetingsAddExample{}
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
