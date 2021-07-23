package configuration

import (
	"net/http"
	// used for side effects
	_ "net/http/pprof"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// Certificate file path declaration
var (
	FilePathPem = "" // FilePathPem is a global const storing https certificate crt file
	FilePathKey = "" // FilePathKey is a global const storing https certificate key file
)

// Listen is a func to start http server
func Listen(conf AppConfig, handler http.Handler) {
	addr := conf.App.Host + ":" + conf.App.Port
	log.Info("run on " + addr)
	var err error
	if conf.App.Protocol == "http" {
		err = http.ListenAndServe(addr, handler)
	} else {
		_, filename, _, _ := runtime.Caller(1)
		FilePathPem = path.Join(path.Dir(filename), FilePathPem)
		FilePathKey = path.Join(path.Dir(filename), FilePathKey)
		err = http.ListenAndServeTLS(addr, FilePathPem, FilePathKey, handler)
	}
	if err != nil {
		log.Fatal(err)
	}
}
