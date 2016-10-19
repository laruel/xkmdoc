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

var _ = Resource("xkm", func() {                // Resources group related API endpoints
	BasePath("/rest")                       // together. They map to REST resources for REST
	DefaultMedia("application/json")                  // services.

	Action("update", func() {                    // Actions define a single API endpoint together
		Description("更新任务属性")    // with its path, parameters (both path
		Routing(PUT("/task/add/v1"))         // parameters and querystring values) and payload
		Params(func() {                    // (shape of the request body).
			Param("param", String, "修改的参数名字")
			Param("type", Integer, "修改的类型  0 普通类型  1 增加成员  2 删除成员")
		})
		Response(OK)                       // Responses define the shape and status code
	})
	Action("update", func() {                    // Actions define a single API endpoint together
		Description("更新子任务属性")    // with its path, parameters (both path
		Routing(PUT("/task/subtask/add/v1"))         // parameters and querystring values) and payload
		Params(func() {                    // (shape of the request body).
			Param("param", String, "修改的参数名字")
			Param("type", Integer, "修改的类型  0 普通类型  1 增加成员  2 删除成员")
		})
		Response(OK)                       // Responses define the shape and status code
		Response(NotFound)                 // of HTTP responses.
	})

	Action("show", func() {
		Description("获取科室的名称和医院名称  供添加团队时使用\n参数 id  科室完整ID\n返回结果：name 科室名称  hospital 医院名称")
		Routing(GET("/department/name/v1/{id}"))
		Response(OK)
	})

	Action("create", func() {
		Description("加入科室")
		Routing(POST("/department/in/v1"))
		Params(func() {                    // (shape of the request body).
			Param("id", String, "当前用户ID")
			Param("departId", String, "科室ID")
		})
		Response(OK)
	})

})
