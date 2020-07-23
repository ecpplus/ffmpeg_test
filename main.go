package main

import (
	"log"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
)

func main() {
	filename := "sample1.mp3"
	// filename := "sample1.flac"
	// filename := "sample1.m4a"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	// ファイルの情報を標準出力
	// ctx.AvDumpFormat(0, filename, 0)

	for i := 0; i < int(ctx.NbStreams()); i++ {
		if avformat.AVMEDIA_TYPE_AUDIO == ctx.Streams()[i].CodecParameters().AvCodecGetType() {
			pCodecCtxOrig := ctx.Streams()[i].Codec()

			// MP3, M4A などは、 GetCodecId と比較して得られる

			switch pCodecCtxOrig.GetCodecId() {
			case avformat.CodecId(avcodec.AV_CODEC_ID_MP3):
				log.Printf("this is mp3")
				break
			case avformat.CodecId(avcodec.AV_CODEC_ID_FLAC):
				log.Printf("this is flac")
				break
			case avformat.CodecId(avcodec.AV_CODEC_ID_AAC):
				log.Printf("this is m4a")
				break
			}
		}
	}
}
