package middleware

import (
	"log"
	"net/http"
)

type MiddlewareHandler map[string]http.Handler

const (
	stsHeader           = "Strict-Transport-Security: max-age=31536000; includeSubDomains"
	frameOptionsHeader  = "X-Frame-Options: Deny"
	contentTypeHeader   = "X-Content-Type-Options: nosniff"
	xssProtectionHeader = "X-XSS-Protection: 1; mode=block"
	cspHeader           = "Content-Security-Policy: default-src 'self'"
)

func (mh MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Basic Logging Middleware
	log.Println(r.RemoteAddr, ":", r.Method, r.RequestURI, r.Proto)

	// Security Headers
	AddSecurityHeaders(w)

	if mh[r.Host] == nil {
		http.Error(w, "Forbidden", 403)
		log.Println("+-->  HostHeaderSecurityMiddleware -- Invalid Host", r.Host)
		return
	}

	mh[r.Host].ServeHTTP(w, r)
}

func AddSecurityHeaders(w http.ResponseWriter) {
	headers := make(map[string]string)
	headers["Strict-Transport-Security"] = "max-age=31536000; includeSubDomains"
	headers["X-Content-Type-Options"] = "nosniff"
	headers["X-Frame-Options"] = "Deny"
	headers["X-XSS-Protection"] = "1; mode=block"
	headers["Content-Security-Policy"] = "default-src 'self'"

	for k, v := range headers {
		w.Header().Add(k, v)
	}
}
