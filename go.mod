module github.com/Phat-FITUS/web-proxy

go 1.20

replace github.com/Phat-FITUS/web-proxy => ../web-proxy

replace github.com/Phat-FITUS/web-proxy/Server => ./Server

replace github.com/Phat-FITUS/web-proxy/HTTP => ./HTTP

require github.com/Phat-FITUS/web-proxy/Server v0.0.0-00010101000000-000000000000

require github.com/Phat-FITUS/web-proxy/HTTP v0.0.0-00010101000000-000000000000 // indirect
