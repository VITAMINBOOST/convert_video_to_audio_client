package libs

import (
	"context"
	"convert_video_to_audio_client/proto"
	"log"
	"os"
	"strings"
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

func errHandler(err error) bool {
	result := true
	if err != nil {
		if code := strings.Split(err.Error(), " "); code[4] != "DeadlineExceeded" {
			log.Fatalf("%s", err)
		} else {
			result = false
		}
	}
	return result
}
