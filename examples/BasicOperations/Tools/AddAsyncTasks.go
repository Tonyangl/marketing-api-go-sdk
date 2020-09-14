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

type AsyncTasksAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.AsyncTasksAddRequest
}

func (e *AsyncTasksAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.AsyncTasksAddRequest{
		TaskName:  "SDK异步任务5ede252e2f1db",
		AccountId: int64(0),
		TaskType:  model.TaskType_TASK_TYPE_AD_HOURLY_REPORT,
		TaskSpec: &model.TaskSpec{
			TaskTypeAdHourlyReportSpec: &model.TaskTypeAdHourlyReportSpec{
				Date: "REPORT DATE",
			},
		},
	}
}

func (e *AsyncTasksAddExample) RunExample() (model.AsyncTasksAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AsyncTasks().Add(ctx, e.Data)
}

func main() {
	e := &AsyncTasksAddExample{}
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
