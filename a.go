package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("V1.6.1", func() {                     // API defines the microservice endpoint and
	Title("服务端接口")           // other global properties. There should be one
	Description("V1.6.1服务端接口调整")        // and exactly one API definition appearing in
	Scheme("https")                             // the design.
	Host("120.76.168.214:8443")
})

var res = MediaType("application/com.yibei.xkm.res", func() {
	Description("返回结果")
	TypeName("BottleMedia")         // Override default generated name
	ContentType("application/json") // Override default Content-Type header value
	Attributes(func() {
		Attribute("responseMsg", String, "返回结果,如果为1,正常返回.其他,业务异常")
	})
	View("default", func() {
		Attribute("responseMsg")
	})
})

var _ = Resource("xkm", func() {                // Resources group related API endpoints
	BasePath("/rest")                       // together. They map to REST resources for REST
	DefaultMedia("application/json")                  // services.

	Action("更新任务", func() {                    // Actions define a single API endpoint together
		Description("更新任务属性")    // with its path, parameters (both path
		Routing(PUT("/task/add/v1"))         // parameters and querystring values) and payload
		Params(func() {                    // (shape of the request body).
			Param("param", String, "修改的参数名字")
			Param("type", Integer, "修改的类型  0 普通类型  1 增加成员  2 删除成员")
		})
		Response(OK,"application/com.yibei.xkm.res")                       // Responses define the shape and status code
	})
	Action("更新子任务", func() {                    // Actions define a single API endpoint together
		Description("更新子任务属性")    // with its path, parameters (both path
		Routing(PUT("/task/subtask/add/v1"))         // parameters and querystring values) and payload
		Params(func() {                    // (shape of the request body).
			Param("param", String, "修改的参数名字")
			Param("type", Integer, "修改的类型  0 普通类型  1 增加成员  2 删除成员")
		})
		Response(OK)                       // Responses define the shape and status code
		Response(NotFound)                 // of HTTP responses.
	})

	Action("获取医院科室名称", func() {
		Description("获取科室的名称和医院名称  供添加团队时使用\n参数 id  科室完整ID\n返回结果：name 科室名称  hospital 医院名称")
		Routing(GET("/department/name/v1/{id}"))
		Params(func() {                    // (shape of the request body).
			Param("id", String, "科室ID")
		})
		Response(OK)
	})

	Action("获取人员信息", func() {
		Description("根据IMID获取人员信息")
		Routing(GET("/users/depart/v1/{departId}/{ids}"))
		Params(func() {                    // (shape of the request body).
			Param("ids", String, "IMID集合，逗号分隔")
			Param("departId", String, "科室ID")
		})
		Response(OK)
	})

	Action("加入科室", func() {
		Description("加入科室")
		Routing(POST("/department/in/v1"))
		Params(func() {                    // (shape of the request body).
			Param("id", String, "当前用户ID")
			Param("departId", String, "科室ID")
		})
		Response(OK)
	})

	Action("根据权限类型获取权限", func() {
		Description("使用权限值进行区分不同的权限")
		Routing(GET("/access/pc/type/v2/{departId}"))
		Params(func() {                    // (shape of the request body).
			Param("departId", String, "科室ID")
		})
		Response(OK)
	})
	Action("注册医生", func() {
		Description("微信添加医生时使用  如果responseMsg=1 正常添加  =2 已经是团队成员")
		Routing(POST("/doctor/reg/v1"))
		Params(func() {                    // (shape of the request body).
			Param("departId", String, "科室ID")
			Param("doctorId", String, "发送邀请人的ID")
			Param("name", String, "姓名")
			Param("phone", String, "手机号")
		})
		Response(OK)
	})


	Action("注册患者", func() {
		Description("微信添加患者时使用")
		Routing(POST("/patient/reg/v1"))
		Params(func() {                    // (shape of the request body).
			Param("departId", String, "科室ID")
			Param("doctorId", String, "发送邀请人的ID")
			Param("name", String, "姓名")
			Param("phone", String, "手机号")
		})
		Response(OK)
	})})
