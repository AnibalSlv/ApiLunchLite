package manager

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func (m *ApiManager) Core() (*httputil.ReverseProxy, error) {

	var routes = map[string]string{
		"/api":  "http://localhost:8002",
		"/auth": "http://localhost:8001",
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// 1. Identificar la ruta base
			path := req.URL.Path
			var targetURL string

			for prefix, target := range routes {
				if strings.HasPrefix(path, prefix) {
					targetURL = target
					break
				}
			}

			// 2. Si no hay ruta definida, enviamos a un destino por defecto
			if targetURL == "" {
				targetURL = "http://localhost:9000"
			}

			// 3. Parsear y configurar el destino
			target, _ := url.Parse(targetURL)
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host // Importante: Algunos servidores requieren el Host original

			// 4. Inyectar headers de identificación
			req.Header.Set("X-Forwarded-For", req.RemoteAddr)
		},
	}
	return proxy, nil
}
