module github.com/Phat-FITUS/web-proxy/Server

go 1.20

replace github.com/Phat-FITUS/web-proxy/Server => ../Server

replace github.com/Phat-FITUS/web-proxy/HTTP => ../HTTP

replace github.com/Phat-FITUS/web-proxy/Proxy => ../Proxy

require (
	github.com/Phat-FITUS/web-proxy/HTTP v0.0.0-00010101000000-000000000000
	github.com/Phat-FITUS/web-proxy/Proxy v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20230801115018-d63ba01acd4b // indirect
