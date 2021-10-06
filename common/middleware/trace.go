package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Trace 链路追踪
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 创建链路追踪接口
		var sp opentracing.Span
		opName := ctx.Request.URL.Path
		// Attempt to join a trace by getting trace context from the headers.
		// 尝试通过从标头获取跟踪上下文来加入跟踪。

		// Extract() 返回给定'format'和'carrier'的SpanContext实例, SpanContext表示必须传播到子代跨度和跨进程的跨度状态
		wireContext, err := opentracing.GlobalTracer().Extract(
			//TextMap将上下文表示为键：值字符串对。
			//与HTTPHeaders不同，TextMap格式不限制键或以任何方式为字符集赋值。
			//对于Tracer.Inject（）：载体必须是“TextMapWriter”。
			//对于Tracer.Extract（）：载体必须是“TextMapReader”`
			opentracing.TextMap,
			// HTTPHeadersCarrier同时满足TextMapWriter和TextMapReader的要求
			opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			sp = opentracing.StartSpan(opName)
		} else {
			sp = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
		}
		ctx.Set("traceSpan", sp)
		ctx.Next()
		// 设置结束时间戳并完成跨度状态.
		// 除了对Context（）的调用（总是允许的），Finish（）必须是对任何span实例的最后一次调用，并且否则会导致未定义的行为。
		sp.Finish()
	}
}
