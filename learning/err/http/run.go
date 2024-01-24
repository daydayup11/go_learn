package main

//
//import (
//	"context"
//	"fmt"
//	"learning/err/http/proto/proto"
//	"log"
//	"net/http"
//
//	"github.com/grpc-ecosystem/grpc-gateway/runtime"
//	"google.golang.org/grpc"
//)
//
//var (
//	bServiceAddr = "localhost:50051" // b服务的地址
//)
//
//func main() {
//	// 创建gRPC连接
//	conn, err := grpc.Dial(bServiceAddr, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("Failed to dial server: %v", err)
//	}
//	defer conn.Close()
//
//	// 创建gRPC客户端
//	client := file.NewFileServiceClient(conn)
//
//	// 注册gRPC网关处理器
//	gwmux := runtime.NewServeMux()
//	err = file.RegisterFileServiceServer(context.Background(), gwmux, bServiceAddr, []grpc.DialOption{grpc.WithInsecure()})
//	if err != nil {
//		log.Fatalf("Failed to register gateway: %v", err)
//	}
//
//	// 启动HTTP服务
//	httpMux := http.NewServeMux()
//	httpMux.Handle("/", gwmux)
//
//	httpAddr := ":8080"
//	fmt.Printf("HTTP server listening on %s\n", httpAddr)
//	err = http.ListenAndServe(httpAddr, httpMux)
//	if err != nil {
//		log.Fatalf("Failed to start HTTP server: %v", err)
//	}
//}
