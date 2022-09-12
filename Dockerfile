FROM alpine as runtime

COPY go-url-expander /app/go-url-expander

CMD ["/app/go-url-expander"]