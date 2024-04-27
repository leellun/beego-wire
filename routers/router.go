// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zhiqu/controllers"
	"zhiqu/wire"
)

func init() {
	beanFactory, err := wire.NewApp()
	if err != nil {
		panic(err)
	}
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/category",
			beego.NSInclude(
				beanFactory.CategoryController,
			),
		),

		beego.NSNamespace("/collect_video",
			beego.NSInclude(
				&controllers.CollectVideoController{},
			),
		),

		beego.NSNamespace("/history",
			beego.NSInclude(
				&controllers.HistoryController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				beanFactory.UserController,
			),
		),

		beego.NSNamespace("/media",
			beego.NSInclude(
				&controllers.MediaController{},
			),
		),

		beego.NSNamespace("/media_category",
			beego.NSInclude(
				&controllers.MediaCategoryController{},
			),
		),

		beego.NSNamespace("/media_later",
			beego.NSInclude(
				&controllers.MediaLaterController{},
			),
		),

		beego.NSNamespace("/media_tags",
			beego.NSInclude(
				&controllers.MediaTagsController{},
			),
		),

		beego.NSNamespace("/users_block",
			beego.NSInclude(
				&controllers.UsersBlockController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/media_playlist",
			beego.NSInclude(
				&controllers.MediaPlaylistController{},
			),
		),

		beego.NSNamespace("/media_playlistmedia",
			beego.NSInclude(
				&controllers.MediaPlaylistmediaController{},
			),
		),

		beego.NSNamespace("/rating",
			beego.NSInclude(
				&controllers.RatingController{},
			),
		),

		beego.NSNamespace("/reply", beego.NSInclude(&controllers.ReplyController{})),
		beego.NSNamespace("/login", beego.NSInclude(beanFactory.LoginController)),
	)
	//beego.SetStaticPath("/swagger", "swagger")
	beego.AddNamespace(ns)
	// 注册Swagger路由
}
