package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/brunoshiroma/go-url-expander/web"
)

var (
	client http.Client
)

func init() {
	handleRedirect := func(req *http.Request, via []*http.Request) error {
		if len(via) == Config.MaxRedirects {
			return errors.New("TARGET_URL")
		}
		return nil
	}

	client = http.Client{
		CheckRedirect: handleRedirect, //TODO check why not calling
		Timeout:       time.Second * time.Duration(Config.RequestTimeout),
	}
}
func handle(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("url") {
		w.WriteHeader(400)
		return
	}

	var (
		targetUrl    = req.URL.Query().Get("url")
		cleanUTM     = req.URL.Query().Get("no-utm")
		redirect     = req.URL.Query().Get("redirect")
		resp         *http.Response
		err          error
		resultUrl    string
		resultRawUrl url.URL
	)

	resp, err = client.Get(targetUrl)
	if err != nil && !strings.Contains(err.Error(), "TARGET_URL") {
		w.Write([]byte(err.Error()))
		return
	}

	if cleanUTM == "1" {
		resultRawUrl = *resp.Request.URL
		query := resultRawUrl.Query()
		query.Del("utm_source")
		query.Del("utm_medium")
		query.Del("utm_campaign")
		resultRawUrl.RawQuery = query.Encode()
		resultUrl = resultRawUrl.String()
	} else {
		resultUrl = resp.Request.URL.String()
	}

	if redirect == "1" && req.Header.Get("X-Requested-With") != "XMLHttpRequest" { // only redirect if not ajax
		w.Header().Set("Location", resultUrl)
		w.WriteHeader(302)
		return
	}

	w.Write([]byte(resultUrl))

}

func NewHttpServer(port int, host string) *http.Server {
	http.HandleFunc("/expand", handle)
	http.Handle("/", http.FileServer(http.FS(web.Content)))
	var server = &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}

	return server
}
