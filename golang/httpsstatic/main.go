package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

var domain string = "example.com"

func redirect(w http.ResponseWriter, req *http.Request) {

	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	target := "https://" + req.Host
	agent := req.UserAgent()
	log.Printf("Redirecting [%s] - [%s] to: %s", agent, ip, target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)

}

func main() {

	fdomain := flag.String("host", "", "Host domain.")
	domain = *fdomain

	staticDir := "static"
	certCache := "tmp/certs"
	TLSAddr := ":https"

	// Use LE key & cert retrieved from tls.Config GetCertificate
	certFile, keyFile := "", ""

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache(certCache),
	}

	SetHSTSHeader := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)

			l := "[" + r.UserAgent() + "] - [" + ip + "]"
			log.Print(l)

			headers := map[string]string{
				"Strict-Transport-Security": "max-age=31557600; includeSubDomains",
				"X-Content-Type-Options":    "nosniff",
				"X-XSS-Protection":          "1; mode=block",
				"Content-Security-Policy":   "default-src 'self'; script-src 'self'",
				"X-Frame-Options":           "DENY",
				"Referrer-Policy":           "no-referrer",
				"Public-Key-Pins": "pin-sha256=\"uIgDNRW0N1ZqFBvx6qJWFqIlaR2rZH/Yr35ZNB+KdHE=\";" + // Site cert
					"pin-sha256=\"BackupBackupBackupBackupBackupBackupBackups=\";" + // Site backup cert
					"pin-sha256=\"YLh1dUR9y6Kja30RrAn7JKnbQG/uEtLMkBgFF2Fuihg=\";" + // LE X3 root CA
					"pin-sha256=\"Vjs8r4z+80wjNcr1YKepWQboSIRi63WsWXhIMN+eWys=\";" + // LE DST X3 root CA
					"includeSubdomains; max-age=2592000",
			}
			for k, v := range headers {
				//fmt.Println(k, ":", v)
				w.Header().Add(k, v)
			}
			h.ServeHTTP(w, r)
		}
	}

	HTTPSServer := &http.Server{
		Addr: TLSAddr,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate, MinVersion: tls.VersionTLS10, // MaxVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, // Required by Go (and HTTP/2 RFC), even if you only present ECDSA certs
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				/*
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				*/
			},
		},
	}

	http.Handle("/", SetHSTSHeader(http.FileServer(http.Dir(staticDir))))

	// HTTP (redirect) Listener
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	// HTTPS Listener
	err := HTTPSServer.ListenAndServeTLS(certFile, keyFile)

	log.Fatal(err)
}
