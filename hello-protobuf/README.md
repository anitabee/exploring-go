# Hello Protocol Buffers

Protocol buffers, are essentially a data format, much like JSON or XML in the sense that they store structured data which can be serialized or de-serialized by a wide number of different languages.

## This is an example using Protocol Buffers in Go language

Install protobuf compiler (which is a standalone binary named protoc) and plugin `protoc-gen-go` requried by `protoc` compiler to generate Go code.

` $ go get -v -u github.com/golang/protobuf/{proto,protoc-gen-go} `

If you are having issues with binary not found on your system make sure that you add both GOPATH and GOROOT in PATH. Also check following [issue](https://github.com/golang/protobuf/issues/795) for more details.

Once protoc binary is available in terminal we can create *.proto data definition; in this case that is message.proto file:

```go
syntax="proto3";

option go_package = "/greeting";


message Message {
      string greet = 1;

}
```

Now we can compile structure using the protocol buffer compiler.
The `protoc` compiler produces Go output when invoked with the go_out flag. The argument to the go_out flag is the directory where you want the compiler to write your Go output.

` $ protoc --go_out=. message.proto `

The compiler will read input file message.proto, and write output file message.pb.go to the out directory.
