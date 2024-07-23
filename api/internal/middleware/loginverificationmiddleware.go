package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 中间件的执行函数
		if r.Header.Get("token") == "123456" {
			next(w, r)
			return
		}
		// Passthrough to next handler if need
		// 如果不满足要求，则直接返回
		w.Write([]byte("权限不足"))
		return
	}
}
