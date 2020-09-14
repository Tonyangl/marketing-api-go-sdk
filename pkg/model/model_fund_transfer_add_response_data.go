/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type FundTransferAddResponseData struct {
	Amount         int64          `json:"amount,omitempty"`
	ExternalBillNo string         `json:"external_bill_no,omitempty"`
	TransTime      int64          `json:"trans_time,omitempty"`
	FundType       AccountTypeMap `json:"fund_type,omitempty"`
	IsRepeated     bool           `json:"is_repeated,omitempty"`
}
