package libs

import (
	"context"
	"convert_video_to_audio_client/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

const (
	serverAddr = ""
	devAddr    = "localhost:30051"
	deadlineMs = 2 * 1000
)

type ConvertVideoToAudioRequest struct {
	SourceVideoURL string `json:sourceVideoURL`
	SheetName      string `json:sheetName`
}

func rpcConnect() *grpc.ClientConn {
	var conn *grpc.ClientConn
	var err error

	if os.Getenv("env") == "DEV" {
		conn, err = grpc.Dial(devAddr, grpc.WithInsecure())
	} else {
		conn, err = grpc.Dial(serverAddr, grpc.WithInsecure())
	}

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return conn
}

func ConvertVideoToAudio(convertVideoToAudioRequest ConvertVideoToAudioRequest) {
	conn := rpcConnect()
	defer conn.Close()

	c := proto.NewAudioDownloadServiceClient(conn)

	clientDeadline := time.Now().Add(time.Duration(deadlineMs) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	rpcRequest := &proto.ConvertVideoToAudioRequest{
		SourceVideoURL: convertVideoToAudioRequest.SourceVideoURL,
	}
	c.ConvertVideoToAudio(ctx, rpcRequest)
}
