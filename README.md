# opentracing-examples

Simple opentracing example:
 - start jaeger all in one `docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest`
 - start examples `go build && ./opentracing-examples`
 - make some requests `for i in {1..5}; do curl localhost:8080 && echo && sleep 1; done`
 - open jaeger ui `http://localhost:16686`