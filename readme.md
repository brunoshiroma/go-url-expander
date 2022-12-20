# URL Expander
Open shortened url on remote server and retrieve the final full url ( or without utm )

## env var
| name | desc |
|------|------|
| PORT | used port for the http server |
| BIND_HOST | ip host for bind ( default is 127.0.0.1 ) |
| REQUEST_TIMEOUT | timeout for making the remote http request default 5s |

More details on file [config.go](internal/config.go#L10)
