/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告包app应用信息
type AppAndroidUnionSpec struct {
	Packname           string                      `json:"packname,omitempty"`
	Version            string                      `json:"version,omitempty"`
	Icon               string                      `json:"icon,omitempty"`
	PackageSize        string                      `json:"package_size,omitempty"`
	PackageMd5         string                      `json:"package_md5,omitempty"`
	PackageDownloadUrl string                      `json:"package_download_url,omitempty"`
	ChannelPackageSpec *[]ChannelPackageSpecStruct `json:"channel_package_spec,omitempty"`
}
