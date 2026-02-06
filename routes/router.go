package routes

import (
	"net/http"
)

// Router 路由组结构体
type Router struct {
	prefix      string
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

// NewRouter 创建新的路由组
func NewRouter(prefix string) *Router {
	return &Router{
		prefix:      prefix,
		middlewares: []func(http.HandlerFunc) http.HandlerFunc{},
	}
}

// Use 添加中间件到路由组
func (r *Router) Use(middleware func(http.HandlerFunc) http.HandlerFunc) {
	r.middlewares = append(r.middlewares, middleware)
}

// applyMiddlewares 应用所有中间件到处理函数
func (r *Router) applyMiddlewares(handler http.HandlerFunc) http.HandlerFunc {
	for _, mw := range r.middlewares {
		handler = mw(handler)
	}
	return handler
}

// GET 注册GET请求路由
func (r *Router) GET(path string, handler http.HandlerFunc) {
	fullPath := r.prefix + path
	http.HandleFunc(fullPath, r.applyMiddlewares(handler))
}

// POST 注册POST请求路由
func (r *Router) POST(path string, handler http.HandlerFunc) {
	fullPath := r.prefix + path
	http.HandleFunc(fullPath, r.applyMiddlewares(handler))
}

// PUT 注册PUT请求路由
func (r *Router) PUT(path string, handler http.HandlerFunc) {
	fullPath := r.prefix + path
	http.HandleFunc(fullPath, r.applyMiddlewares(handler))
}

// DELETE 注册DELETE请求路由
func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	fullPath := r.prefix + path
	http.HandleFunc(fullPath, r.applyMiddlewares(handler))
}

// Group 创建新的路由组
func (r *Router) Group(prefix string) *Router {
	return NewRouter(r.prefix + prefix)
}
