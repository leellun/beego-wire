// Package routers
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	controller "zhiqu/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/auth_permission",
			beego.NSInclude(
				&controller.AuthPermissionController{},
			),
		),

		beego.NSNamespace("/files_category",
			beego.NSInclude(
				&controller.FilesCategoryController{},
			),
		),

		beego.NSNamespace("/django_migrations",
			beego.NSInclude(
				&controller.DjangoMigrationsController{},
			),
		),

		beego.NSNamespace("/django_session",
			beego.NSInclude(
				&controller.DjangoSessionController{},
			),
		),

		beego.NSNamespace("/files_comment",
			beego.NSInclude(
				&controller.FilesCommentController{},
			),
		),

		beego.NSNamespace("/files_encodeprofile",
			beego.NSInclude(
				&controller.FilesEncodeprofileController{},
			),
		),

		beego.NSNamespace("/django_site",
			beego.NSInclude(
				&controller.DjangoSiteController{},
			),
		),

		beego.NSNamespace("/actions_media_action",
			beego.NSInclude(
				&controller.ActionsMediaActionController{},
			),
		),

		beego.NSNamespace("/auth_group",
			beego.NSInclude(
				&controller.AuthGroupController{},
			),
		),

		beego.NSNamespace("/django_content_type",
			beego.NSInclude(
				&controller.DjangoContentTypeController{},
			),
		),

		beego.NSNamespace("/auth_token",
			beego.NSInclude(
				&controller.AuthTokenController{},
			),
		),

		beego.NSNamespace("/django_admin_log",
			beego.NSInclude(
				&controller.DjangoAdminLogController{},
			),
		),

		beego.NSNamespace("/account_email_confirmation",
			beego.NSInclude(
				&controller.AccountEmailConfirmationController{},
			),
		),

		beego.NSNamespace("/files_media_category",
			beego.NSInclude(
				&controller.FilesMediaCategoryController{},
			),
		),

		beego.NSNamespace("/files_media_rating_category",
			beego.NSInclude(
				&controller.FilesMediaRatingCategoryController{},
			),
		),

		beego.NSNamespace("/files_encoding",
			beego.NSInclude(
				&controller.FilesEncodingController{},
			),
		),

		beego.NSNamespace("/files_license",
			beego.NSInclude(
				&controller.FilesLicenseController{},
			),
		),

		beego.NSNamespace("/files_language",
			beego.NSInclude(
				&controller.FilesLanguageController{},
			),
		),

		beego.NSNamespace("/files_media_tags",
			beego.NSInclude(
				&controller.FilesMediaTagsController{},
			),
		),

		beego.NSNamespace("/files_tag",
			beego.NSInclude(
				&controller.FilesTagController{},
			),
		),

		beego.NSNamespace("/socialaccount_socialaccount",
			beego.NSInclude(
				&controller.SocialaccountSocialaccountController{},
			),
		),

		beego.NSNamespace("/files_playlist",
			beego.NSInclude(
				&controller.FilesPlaylistController{},
			),
		),

		beego.NSNamespace("/socialaccount_socialapp",
			beego.NSInclude(
				&controller.SocialaccountSocialappController{},
			),
		),

		beego.NSNamespace("/socialaccount_socialapp_sites",
			beego.NSInclude(
				&controller.SocialaccountSocialappSitesController{},
			),
		),

		beego.NSNamespace("/files_rating_category",
			beego.NSInclude(
				&controller.FilesRatingCategoryController{},
			),
		),

		beego.NSNamespace("/files_playlistmedia",
			beego.NSInclude(
				&controller.FilesPlaylistmediaController{},
			),
		),

		beego.NSNamespace("/files_rating",
			beego.NSInclude(
				&controller.FilesRatingController{},
			),
		),

		beego.NSNamespace("/files_subtitle",
			beego.NSInclude(
				&controller.FilesSubtitleController{},
			),
		),

		beego.NSNamespace("/users_channel",
			beego.NSInclude(
				&controller.UsersChannelController{},
			),
		),

		beego.NSNamespace("/users_channel_subscribers",
			beego.NSInclude(
				&controller.UsersChannelSubscribersController{},
			),
		),

		beego.NSNamespace("/users_notification",
			beego.NSInclude(
				&controller.UsersNotificationController{},
			),
		),

		beego.NSNamespace("/users_user_groups",
			beego.NSInclude(
				&controller.UsersUserGroupsController{},
			),
		),

		beego.NSNamespace("/users_user_user_permissions",
			beego.NSInclude(
				&controller.UsersUserUserPermissionsController{},
			),
		),

		beego.NSNamespace("/users_user",
			beego.NSInclude(
				&controller.UsersUserController{},
			),
		),

		beego.NSNamespace("/files_media",
			beego.NSInclude(
				&controller.FilesMediaController{},
			),
		),

		beego.NSNamespace("/auth_group_permissions",
			beego.NSInclude(
				&controller.AuthGroupPermissionsController{},
			),
		),

		beego.NSNamespace("/socialaccount_socialtoken",
			beego.NSInclude(
				&controller.SocialaccountSocialtokenController{},
			),
		),

		beego.NSNamespace("/account_email_address",
			beego.NSInclude(
				&controller.AccountEmailAddressController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
