package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/pkg/errors"
//	"google.golang.org/grpc"
//	"learning/err/http/proto"
//	"log"
//	"net"
//	"os"
//)
//
//// FoundFile 表示找到文件时的错误类型
//type FoundFile struct {
//	Filename string
//	Err      error
//}
//
//// 实现 Unwrap 方法
//func (e *FoundFile) Unwrap() error {
//	return e.Err
//}
//
//// 实现 error 接口的 Error 方法
//func (e *FoundFile) Error() string {
//	return fmt.Sprintf("%s: %v", e.Filename, e.Err)
//}
//
//var ErrFileNotFound = errors.New("file not found")
//
//type FileServiceServer struct{}
//
//func (s *FileServiceServer) OpenFile(ctx context.Context, req *proto.FileRequest) (*proto.FileResponse, error) {
//	// 在这里调用 openFile 函数
//	_, err := os.Open(req.FileName)
//	if err != nil {
//		return &proto.FileResponse{Filename: req.Filename, ErrorMessage: err.Error()}, nil
//	}
//
//	return &proto.FileResponse{Filename: req.Filename}, nil
//}
//
////
////func openFile(filename string) (*os.File, error) {
////	file, err := os.Open(filename)
////	if err != nil {
////		return nil, &FoundFile{filename, ErrFileNotFound}
////	}
////	return file, nil
////}
//
//func main() {
//	listen, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	proto.RegisterFileServiceServer(grpcServer, &FileServiceServer{})
//	grpcServer.Serve(listen)
//}
