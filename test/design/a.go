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
		Routing(PUT("/add/v1"))         // parameters and querystring values) and payload
		Params(func() {                    // (shape of the request body).
			Param("param", String, "修改的参数名字")
			Param("type", Integer, "修改的类型  0 普通类型  1 增加成员  2 删除成员")
		})
		Response(OK)                       // Responses define the shape and status code
		Response(NotFound)                 // of HTTP responses.
	})

})
