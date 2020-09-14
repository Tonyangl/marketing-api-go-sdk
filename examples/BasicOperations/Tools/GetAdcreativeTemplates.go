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

type AdcreativeTemplatesGetExample struct {
	TAds                       *ads.SDKClient
	AccessToken                string
	AdcreativeTemplatesGetOpts *api.AdcreativeTemplatesGetOpts
}

func (e *AdcreativeTemplatesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AdcreativeTemplatesGetOpts = &api.AdcreativeTemplatesGetOpts{
		AccountId: optional.NewInt64(int64(0)),

		Fields: optional.NewInterface([]string{"adcreative_template_id", "adcreative_template_name", "adcreative_template_description", "adcreative_template_size", "adcreative_template_style", "adcreative_template_appellation", "site_set", "promoted_object_type", "adcreative_sample_image_list", "ad_attributes", "adcreative_attributes", "adcreative_elements", "support_page_type", "support_billing_spec_list", "support_dynamic_ability_spec_list"}),
	}
}

func (e *AdcreativeTemplatesGetExample) RunExample() (model.AdcreativeTemplatesGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AdcreativeTemplates().Get(ctx, e.AdcreativeTemplatesGetOpts)
}

func main() {
	e := &AdcreativeTemplatesGetExample{}
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
