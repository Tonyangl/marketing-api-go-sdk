/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 定向详细设置
type EstimationReadTargetingSetting struct {
	Age                              *[]AgeStruct                       `json:"age,omitempty"`
	Gender                           *[]string                          `json:"gender,omitempty"`
	Education                        *[]string                          `json:"education,omitempty"`
	MaritalStatus                    *[]string                          `json:"marital_status,omitempty"`
	WorkingStatus                    *[]string                          `json:"working_status,omitempty"`
	GeoLocation                      *GeoLocations                      `json:"geo_location,omitempty"`
	UserOs                           *[]string                          `json:"user_os,omitempty"`
	NewDevice                        *[]string                          `json:"new_device,omitempty"`
	DevicePrice                      *[]string                          `json:"device_price,omitempty"`
	NetworkType                      *[]string                          `json:"network_type,omitempty"`
	NetworkOperator                  *[]string                          `json:"network_operator,omitempty"`
	NetworkScene                     *[]string                          `json:"network_scene,omitempty"`
	DressingIndex                    *[]string                          `json:"dressing_index,omitempty"`
	UvIndex                          *[]string                          `json:"uv_index,omitempty"`
	MakeupIndex                      *[]string                          `json:"makeup_index,omitempty"`
	Climate                          *[]string                          `json:"climate,omitempty"`
	Temperature                      *[]TemperatureStruct               `json:"temperature,omitempty"`
	AppInstallStatus                 *[]string                          `json:"app_install_status,omitempty"`
	MiniGameQqStatus                 *[]string                          `json:"mini_game_qq_status,omitempty"`
	ConsumptionStatus                *[]string                          `json:"consumption_status,omitempty"`
	GamerConsumptionAbility          *[]PlayerConsuptStruct             `json:"gamer_consumption_ability,omitempty"`
	GameConsumptionLevel             *[]string                          `json:"game_consumption_level,omitempty"`
	PaidUser                         *[]string                          `json:"paid_user,omitempty"`
	ResidentialCommunityPrice        *[]ResidentialCommunityPriceStruct `json:"residential_community_price,omitempty"`
	WechatAdBehavior                 *LimitWechatAdBehavior             `json:"wechat_ad_behavior,omitempty"`
	CustomAudience                   *[]int64                           `json:"custom_audience,omitempty"`
	ExcludedCustomAudience           *[]int64                           `json:"excluded_custom_audience,omitempty"`
	DeprecatedCustomAudience         *[]int64                           `json:"deprecated_custom_audience,omitempty"`
	DeprecatedExcludedCustomAudience *[]int64                           `json:"deprecated_excluded_custom_audience,omitempty"`
	BehaviorOrInterest               *BehaviorOrInterest                `json:"behavior_or_interest,omitempty"`
	AirQualityIndex                  *[]string                          `json:"air_quality_index,omitempty"`
	WechatOfficialAccountCategory    *[]int64                           `json:"wechat_official_account_category,omitempty"`
	MobileUnionCategory              *[]int64                           `json:"mobile_union_category,omitempty"`
	FinancialSituation               *[]string                          `json:"financial_situation,omitempty"`
	ConsumptionType                  *[]string                          `json:"consumption_type,omitempty"`
}
