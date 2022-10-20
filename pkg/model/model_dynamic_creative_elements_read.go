/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意元素
type DynamicCreativeElementsRead struct {
	Image                            *string                            `json:"image,omitempty"`
	Image2                           *string                            `json:"image2,omitempty"`
	Image3                           *string                            `json:"image3,omitempty"`
	Title                            *string                            `json:"title,omitempty"`
	Description                      *string                            `json:"description,omitempty"`
	Corporate                        *AdcreativeCorporate               `json:"corporate,omitempty"`
	Video                            *string                            `json:"video,omitempty"`
	ImageOptions                     *[]string                          `json:"image_options,omitempty"`
	ImageListOptions                 *[]AdcreativeImageList             `json:"image_list_options,omitempty"`
	TitleOptions                     *[]string                          `json:"title_options,omitempty"`
	DescriptionOptions               *[]string                          `json:"description_options,omitempty"`
	VideoOptions                     *[]string                          `json:"video_options,omitempty"`
	ShortVideoStructOptions          *[]ShortVideoStruct                `json:"short_video_struct_options,omitempty"`
	DeepLinkType                     *string                            `json:"deep_link_type,omitempty"`
	LinkNameType                     LinkNameType                       `json:"link_name_type,omitempty"`
	ImageList                        *[]string                          `json:"image_list,omitempty"`
	ElementStory                     *[]AdcreativeElementStoryArrayItem `json:"element_story,omitempty"`
	Url                              *string                            `json:"url,omitempty"`
	ButtonText                       *string                            `json:"button_text,omitempty"`
	BottomText                       *string                            `json:"bottom_text,omitempty"`
	ActivityType                     *string                            `json:"activity_type,omitempty"`
	CountdownBegin                   *int64                             `json:"countdown_begin,omitempty"`
	CountdownExpiringTimestamp       *int64                             `json:"countdown_expiring_timestamp,omitempty"`
	CountdownPrice                   *string                            `json:"countdown_price,omitempty"`
	CountdownTimeType                AdCreativeCountdownTimeType        `json:"countdown_time_type,omitempty"`
	MiniProgramId                    *string                            `json:"mini_program_id,omitempty"`
	MiniProgramPath                  *string                            `json:"mini_program_path,omitempty"`
	Label                            *[]CreativeLabel                   `json:"label,omitempty"`
	ProductTags                      *[]string                          `json:"product_tags,omitempty"`
	LogoDescription                  *string                            `json:"logo_description,omitempty"`
	Logo                             *string                            `json:"logo,omitempty"`
	LeftButton                       *string                            `json:"left_button,omitempty"`
	RightButton                      *string                            `json:"right_button,omitempty"`
	LeftBottomTxt                    *string                            `json:"left_bottom_txt,omitempty"`
	AnimationEffect                  *string                            `json:"animation_effect,omitempty"`
	Phone                            *string                            `json:"phone,omitempty"`
	Canvas                           *string                            `json:"canvas,omitempty"`
	Caption                          *string                            `json:"caption,omitempty"`
	MyappPageId                      *string                            `json:"myapp_page_id,omitempty"`
	LogoPage                         *AdcreativeLogoPage                `json:"logo_page,omitempty"`
	VideoPopupUrl                    *string                            `json:"video_popup_url,omitempty"`
	VideoPopupButton                 *AdcreativeVideoPopupButton        `json:"video_popup_button,omitempty"`
	VideoPopupButtonText             *string                            `json:"video_popup_button_text,omitempty"`
	VideoPopupButtonUrl              *string                            `json:"video_popup_button_url,omitempty"`
	ButtonUrl                        *string                            `json:"button_url,omitempty"`
	Brand                            *BrandStruct                       `json:"brand,omitempty"`
	MyappLandingPageInfo             *AdcreativeMyappLandingPageInfo    `json:"myapp_landing_page_info,omitempty"`
	LongVideo1                       *string                            `json:"long_video1,omitempty"`
	LongVideo2                       *string                            `json:"long_video2,omitempty"`
	ShortVideoStruct                 *ShortVideoStruct                  `json:"short_video_struct,omitempty"`
	QzoneVideoPageId                 *string                            `json:"qzone_video_page_id,omitempty"`
	Qq                               *string                            `json:"qq,omitempty"`
	LeftCanvas                       *string                            `json:"left_canvas,omitempty"`
	RightCanvas                      *string                            `json:"right_canvas,omitempty"`
	SunText                          *string                            `json:"sun_text,omitempty"`
	CloudText                        *string                            `json:"cloud_text,omitempty"`
	OvercastText                     *string                            `json:"overcast_text,omitempty"`
	RainText                         *string                            `json:"rain_text,omitempty"`
	SnowText                         *string                            `json:"snow_text,omitempty"`
	FogText                          *string                            `json:"fog_text,omitempty"`
	SandText                         *string                            `json:"sand_text,omitempty"`
	HazeText                         *string                            `json:"haze_text,omitempty"`
	LabelledImg                      *AdcreativeLabelledImg             `json:"labelled_img,omitempty"`
	ShareImg                         *string                            `json:"share_img,omitempty"`
	LongVideoStruct                  *LongVideoStruct                   `json:"long_video_struct,omitempty"`
	BannerContent                    *AdcreativeBannerContent           `json:"banner_content,omitempty"`
	CardContent                      *AdcreativeCardContent             `json:"card_content,omitempty"`
	HeadLine                         *string                            `json:"head_line,omitempty"`
	ShopImageStruct                  *AdCreativeShopImageStruct         `json:"shop_image_struct,omitempty"`
	ChosenButton                     *ChosenButton                      `json:"chosen_button,omitempty"`
	LivingDescStruct                 *AdCreativeLivingDescStruct        `json:"living_desc_struct,omitempty"`
	FloatingZoneStruct               *FloatingZone                      `json:"floating_zone_struct,omitempty"`
	CanvasShareImage                 *string                            `json:"canvas_share_image,omitempty"`
	WegameInfoSpec                   *WegameInfoSpec                    `json:"wegame_info_spec,omitempty"`
	FinderObjectVisibility           *bool                              `json:"finder_object_visibility,omitempty"`
	TitleComponentOptions            *[]TextComponentOption             `json:"title_component_options,omitempty"`
	DescriptionComponentOptions      *[]TextComponentOption             `json:"description_component_options,omitempty"`
	ImageComponentOptions            *[]ImageComponentOption            `json:"image_component_options,omitempty"`
	ImageListComponentOptions        *[]ImageListComponentOption        `json:"image_list_component_options,omitempty"`
	ImageTextListComponentOptions    *[]ImageTextListComponentOption    `json:"image_text_list_component_options,omitempty"`
	Image2ComponentOptions           *[]ImageComponentOption            `json:"image2_component_options,omitempty"`
	LandingPageComponentOptions      *[]LandingPageComponentOption      `json:"landing_page_component_options,omitempty"`
	BrandComponentOptions            *[]BrandComponentOption            `json:"brand_component_options,omitempty"`
	LongSublinkComponentOptions      *[]LongSublinkComponentOption      `json:"long_sublink_component_options,omitempty"`
	LongSublinkListComponentOptions  *[]LongSublinkListComponentOption  `json:"long_sublink_list_component_options,omitempty"`
	ShortSublinkComponentOptions     *[]ShortSublinkComponentOption     `json:"short_sublink_component_options,omitempty"`
	ShortSublinkListComponentOptions *[]ShortSublinkListComponentOption `json:"short_sublink_list_component_options,omitempty"`
	VideoComponentOptions            *[]VideoComponentOption            `json:"video_component_options,omitempty"`
	ConsultComponentOptions          *[]ConsultComponentOption          `json:"consult_component_options,omitempty"`
	PhoneComponentOptions            *[]PhoneComponentOption            `json:"phone_component_options,omitempty"`
	FormComponentOptions             *[]FormComponentOption             `json:"form_component_options,omitempty"`
}
