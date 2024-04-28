// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"zhiqu/wire"
)

func init() {
	beanFactory, err := wire.NewApp()
	if err != nil {
		panic(err)
	}
	//这是一种可实现的依赖注入全自动方式
	web.NewNamespace("/v1", func(namespace *web.Namespace) {
		namespace.Get("", func(ctx *context.Context) {

		})
	})
	web.Get("/v1/user/:id", beanFactory.UserController.GetOne)
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/category", beego.NSInclude(&category.Controller{})),
	//
	//	beego.NSNamespace("/collect_video",
	//		beego.NSInclude(
	//			&controllers.CollectVideoController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/history",
	//		beego.NSInclude(
	//			&controllers.HistoryController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&user.Controller{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media",
	//		beego.NSInclude(
	//			&controllers.MediaController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media_category",
	//		beego.NSInclude(
	//			&controllers.MediaCategoryController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media_later",
	//		beego.NSInclude(
	//			&controllers.MediaLaterController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media_tags",
	//		beego.NSInclude(
	//			&controllers.MediaTagsController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/users_block",
	//		beego.NSInclude(
	//			&controllers.UsersBlockController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/comment",
	//		beego.NSInclude(
	//			&comment.CommentController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media_playlist",
	//		beego.NSInclude(
	//			&controllers.MediaPlaylistController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/media_playlistmedia",
	//		beego.NSInclude(
	//			&controllers.MediaPlaylistmediaController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/rating",
	//		beego.NSInclude(
	//			&controllers.RatingController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/reply", beego.NSInclude(&controllers.ReplyController{})),
	//	beego.NSNamespace("/login", beego.NSInclude(&login.Controller{})),
	//)
	//// 注册Swagger路由
	//beego.SetStaticPath("/swagger", "swagger")
	//beego.AddNamespace(ns)
}
