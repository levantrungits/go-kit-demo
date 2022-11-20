# Go + gRPC with Go-kit (Go in the modern enterprise)
### Website: https://gokit.io/
### Pkg:     https://pkg.go.dev/github.com/go-kit/kit
### Author Go-kit:  [Peter Bourgon](http://peter.bourgon.org/go-kit/#what-go-needs) <br> <br> 

## I. Start service:
1. Gen Proto:   
```make gen-proto```
2. Run:         
```make go-run``` <br> <br> 

## II. Test gRPC with **grpcurl**
Embed: [Reflection](https://github.com/grpc/grpc-go/tree/master/reflection) <br>
Install: [grpcurl](https://github.com/fullstorydev/grpcurl)
```
grpcurl -plaintext localhost:50051 list

grpcurl --plaintext localhost:50051 list MathService

grpcurl --plaintext localhost:50051 describe MathService.Add

grpcurl --plaintext localhost:50051 describe .MathRequest

grpcurl --plaintext -msg-template localhost:50051 
describe .MathRequest

grpcurl --plaintext -d '{
"numA":5,
"numB":102
}' localhost:50051 MathService/Add
{
  "result": 107
}
```