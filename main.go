package main

import (
	"convert_video_to_audio_client/libs"
	"flag"
)

func main() {
	url := flag.String("url", "", "Please input source video url.")

	flag.Parse()
	libs.ConvertVideoToAudio(libs.ConvertVideoToAudioRequest{
		SourceVideoURL: *url,
	})

}
