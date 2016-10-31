package design

import (
. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("V1.6.2", func() {                     // API defines the microservice endpoint and
	Title("服务端接口")           // other global properties. There should be one
	Description("V1.6.1服务端接口调整")        // and exactly one API definition appearing in
	Scheme("https")                             // the design.
	Host("120.76.168.214:8443")
})

var _ = Resource("xkm", func() {
	// Resources group related API endpoints
	BasePath("/rest")                       // together. They map to REST resources for REST
	DefaultMedia("application/json")                  // services.

	Action("删除个人的通知", func() {
		// Actions define a single API endpoint together
		Description("老接口,通知接受者删除通知，只会删除自己的通知")    // with its path, param eters (both path
		Routing(DELETE("/notice/person/v1/{id}/{noticeId}"))         // parameters and querystring values) and payload
		Params(func() {
			// (shape of the request body).
			Param("id", String, "登录者ID，通知接受者的ID")
			Param("noticeId", String, "接受到的通知ID")
		})
		Response(OK)                       // Responses define the shape and status code
	})
	Action("删除发出的通知", func() {
		Description("老接口，通知发起人使用的删除接口，使用该接口，删除系统中的通知，包括接收人收到的通知")
		Routing(DELETE("/notice/v1/{noticeId}"))         // parameters and querystring values) and payload
		Params(func() {
			// (shape of the request body).
			Param("noticeId", String, "接受到的通知ID")
		})
		Response(OK)                       // Responses define the shape and status code
		Response(NotFound)                 // of HTTP responses.
	})

	Action("创建通知",func(){
		Description("创建公告，老接口，使用通用的创建通知接口\n参数描述中，只描述增加的参数，公告分类：category")
		Routing(POST("/notice/v1"))
		Params(func() {
			Param("category",String,"公告的类别")
		})
	})

	Action("提醒通知", func() {
		Description("老接口，增加5分钟的时间判断，如果不达到5分钟再次提醒，错误信息设置在resonseMsg中")
		Routing(POST("/notice/department/v1"))
	})

	Action("更新日程属性",func(){
		Description("更新日程的一个属性,提交结构同于任务的方式，结构与日程结构一致，增加2个属性\n" +
			"当paramType 为1  增加同事或患者   为2  减少同事或患者   其他值 更新param指定的属性")
		Routing(PUT("/schedule/param/v1"))
		Params(func() {
			Param("param",String,"更改参数的属性值")
			Param("paramType",String,"更改类型  具体说明见描述")
		})
	})
})
