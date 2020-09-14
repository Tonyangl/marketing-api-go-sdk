/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type FundStatementsDailyGetListStruct struct {
	TradeType   TradeType      `json:"trade_type,omitempty"`
	TransTime   int64          `json:"trans_time,omitempty"`
	Amount      int64          `json:"amount,omitempty"`
	Description string         `json:"description,omitempty"`
	FundType    AccountTypeMap `json:"fund_type,omitempty"`
}
