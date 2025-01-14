/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AdcreativeTemplateListApiService service

/*
AdcreativeTemplateListApiService 获取创意规格列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param optional nil or *AdcreativeTemplateListGetOpts - Optional Parameters:
     * @param "SiteSet" (optional.String) -
     * @param "CampaignType" (optional.String) -
     * @param "PromotedObjectType" (optional.String) -
     * @param "DynamicAbilityType" (optional.Interface of []string) -
     * @param "IsDynamicCreative" (optional.Bool) -
     * @param "WechatSceneSpecPosition" (optional.Interface of []int64) -
     * @param "AdcreativeTemplateId" (optional.Int64) -
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return AdcreativeTemplateListGetResponse
*/

type AdcreativeTemplateListGetOpts struct {
	SiteSet                 optional.String
	CampaignType            optional.String
	PromotedObjectType      optional.String
	DynamicAbilityType      optional.Interface
	IsDynamicCreative       optional.Bool
	WechatSceneSpecPosition optional.Interface
	AdcreativeTemplateId    optional.Int64
	Page                    optional.Int64
	PageSize                optional.Int64
	Fields                  optional.Interface
}

func (a *AdcreativeTemplateListApiService) Get(ctx context.Context, accountId int64, localVarOptionals *AdcreativeTemplateListGetOpts) (AdcreativeTemplateListGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue AdcreativeTemplateListGetResponseData
		localVarResponse    AdcreativeTemplateListGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/adcreative_template_list/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.SiteSet.IsSet() {
		localVarQueryParams.Add("site_set", parameterToString(localVarOptionals.SiteSet.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CampaignType.IsSet() {
		localVarQueryParams.Add("campaign_type", parameterToString(localVarOptionals.CampaignType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PromotedObjectType.IsSet() {
		localVarQueryParams.Add("promoted_object_type", parameterToString(localVarOptionals.PromotedObjectType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DynamicAbilityType.IsSet() {
		localVarQueryParams.Add("dynamic_ability_type", parameterToString(localVarOptionals.DynamicAbilityType.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.IsDynamicCreative.IsSet() {
		localVarQueryParams.Add("is_dynamic_creative", parameterToString(localVarOptionals.IsDynamicCreative.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WechatSceneSpecPosition.IsSet() {
		localVarQueryParams.Add("wechat_scene_spec_position", parameterToString(localVarOptionals.WechatSceneSpecPosition.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.AdcreativeTemplateId.IsSet() {
		localVarQueryParams.Add("adcreative_template_id", parameterToString(localVarOptionals.AdcreativeTemplateId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			if localVarResponse.Data == nil {
				return localVarReturnValue, localVarHttpResponse.Header, err
			} else {
				return *localVarResponse.Data, localVarHttpResponse.Header, err
			}
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v AdcreativeTemplateListGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
