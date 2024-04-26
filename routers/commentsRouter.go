package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailAddressController"],
		beego.ControllerComments{
			Method:           "AllBlock",
			Router:           `/all/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AccountEmailConfirmationController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:ActionsMediaActionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthGroupPermissionsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthPermissionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:AuthTokenController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoAdminLogController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoContentTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoMigrationsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSessionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:DjangoSiteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesCommentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodeprofileController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesEncodingController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLanguageController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesLicenseController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaCategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesMediaTagsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesPlaylistmediaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingCategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesRatingController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesSubtitleController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:FilesTagController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialaccountController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialappSitesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:SocialaccountSocialtokenController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersChannelSubscribersController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersNotificationController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserController"],
		beego.ControllerComments{
			Method:           "ViewUser",
			Router:           `/:username`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserGroupsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"] = append(beego.GlobalControllerRouter["zhiqu/controllers:UsersUserUserPermissionsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
