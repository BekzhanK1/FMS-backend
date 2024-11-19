package utils

import "net/http"

func CopyHeaders(r *http.Request, req *http.Request) {
	
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	// Copy other headers as needed
	for key, values := range r.Header {
		// Skip the Host header as it's usually managed by the client
		if key == "Authorization" {
			continue
		}
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}
