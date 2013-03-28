package mangogzip

import (
	"bytes"
	"compress/gzip"
	. "github.com/paulbellamy/mango"
	"strings"
)

func Zipper(env Env, app App) (status Status, headers Headers, body Body) {
	status, headers, body = app(env)

	ae := env.Request().Header.Get("Accept-Encoding")

	if headers == nil {
		headers = Headers{}
		headers.Set("Content-Type", "text/html")
	}

	contentType := headers.Get("Content-Type")
	if contentType == "" {
		headers.Set("Content-Type", "text/html; charset=utf8")
	}

	if strings.Contains(ae, "gzip") {
		headers.Set("Content-Encoding", "gzip")

		buff := bytes.NewBuffer(nil)
		gw := gzip.NewWriter(buff)
		gw.Write([]byte(body))
		gw.Close()
		body = Body(buff.String())
	}

	return
}
