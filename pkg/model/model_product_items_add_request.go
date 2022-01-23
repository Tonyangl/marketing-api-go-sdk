/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductItemsAddRequest struct {
	AccountId           *int64             `json:"account_id,omitempty"`
	ProductCatalogId    *int64             `json:"product_catalog_id,omitempty"`
	FeedId              *int64             `json:"feed_id,omitempty"`
	RequestSource       *string            `json:"request_source,omitempty"`
	ProductItemSpecList *[]ProductItemSpec `json:"product_item_spec_list,omitempty"`
}
