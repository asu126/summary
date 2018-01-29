# error info
```
../../../golang.org/x/net/idna/idna.go:23:2: cannot find package "golang.org/x/text/secure/bidirule" in any of:
	/usr/local/go/src/golang.org/x/text/secure/bidirule (from $GOROOT)
	/home/su/develop/gopath/src/golang.org/x/text/secure/bidirule (from $GOPATH)
../../../golang.org/x/net/idna/idna.go:24:2: cannot find package "golang.org/x/text/unicode/norm" in any of:
	/usr/local/go/src/golang.org/x/text/unicode/norm (from $GOROOT)
	/home/su/develop/gopath/src/golang.org/x/text/unicode/norm (from $GOPATH)
../../../google.golang.org/grpc/status/status.go:36:2: cannot find package "google.golang.org/genproto/googleapis/rpc/status" in any of:
	/usr/local/go/src/google.golang.org/genproto/googleapis/rpc/status (from $GOROOT)
	/home/su/develop/gopath/src/google.golang.org/genproto/googleapis/rpc/status (from $GOPATH)
```

# 解决方法
```
go get -u github.com/golang/protobuf/protoc-gen-go
cd $GOPATH/src/golang.org/x          && git clone https://github.com/grpc/grpc-go.git grpc
cd $GOPATH/src/golang.org/x          && git clone https://github.com/golang/text.git
cd $GOPATH/src/google.golang.org git && clone https://github.com/google/go-genproto.git genproto
```
 
