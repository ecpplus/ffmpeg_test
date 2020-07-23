package main

import (
	"fmt"
	"log"
	"os"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
)

func main() {
	// filename := "sample1.wav"
	filename := "sample1.wav"

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

	// ctx.AvDumpFormat(0, filename, 0)

	log.Println("aaa")
	log.Println(avformat.AVMEDIA_TYPE_AUDIO)

	for i := 0; i < int(ctx.NbStreams()); i++ {
		log.Println(ctx.Streams()[i].CodecParameters().AvCodecGetType())
		if avformat.AVMEDIA_TYPE_AUDIO == ctx.Streams()[i].CodecParameters().AvCodecGetType() {
			pCodecCtxOrig := ctx.Streams()[i].Codec()
			pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtxOrig.GetCodecId()))

			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				os.Exit(1)
			}

			log.Printf("codec: %#v", pCodec)

			log.Printf("codecType: %#v", pCodecCtxOrig.GetCodecType())
			log.Printf("codecId: %#v", pCodecCtxOrig.GetCodecId())
			log.Println(avcodec.AV_CODEC_ID_MP3)

			// MP3, M4A などは、 GetCodecId と比較して得られる
			log.Println(pCodecCtxOrig.GetCodecId() == avformat.CodecId(avcodec.AV_CODEC_ID_MP3))
			log.Println(pCodecCtxOrig.GetCodecId() == avformat.CodecId(avcodec.AV_CODEC_ID_WAVPACK))

			// log.Printf("%#v", pCodecCtxOrig.Type())
			// avformat.co

			// pCodecCtx := pCodec.AvcodecAllocContext3()
		}
	}

	// inputFormat := ctx.Iformat()
	// log.Println(inputFormat)

	// log.Println(ctx.AudioPreload())

	// log.Println(ctx.AvFormatGetProbeScore())
	// log.Println(ctx.Duration())
	// log.Println(ctx.BitRate())
	// // avcodec := ctx.AvFormatGetAudioCodec()
	// log.Println(ctx.Filename())
	// log.Println(ctx.AvFormatGetAudioCodec())
	// log.Println(ctx.AudioCodecId())
	// log.Println(ctx.Oformat())
	// log.Println(ctx.AvFormatGetAudioCodec())
	// log.Println(ctx.AudioCodec())

	// outputFormat := avformat.AvGuessFormat("", filename, "")
	// log.Println(outputFormat

	log.Println("OK")
	//...

}
