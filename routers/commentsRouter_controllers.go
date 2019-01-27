package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gobook/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gobook/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticle",
            Router: `/addArticle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gobook/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "QueryAllArticles",
            Router: `/queryAllArticles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gobook/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "QueryArticle",
            Router: `/queryOneArticle/:articleId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ArticleController"] = append(beego.GlobalControllerRouter["gobook/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "QueryUserArticles",
            Router: `/queryUserArticles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gobook/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gobook/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gobook/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gobook/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gobook/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
        beego.ControllerComments{
            Method: "IsLogined",
            Router: `/isLogined`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gobook/controllers:UserController"] = append(beego.GlobalControllerRouter["gobook/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
