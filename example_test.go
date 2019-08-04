package youtube_test

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	. "github.com/kkdai/youtube"
)

//ExampleDownload : Example code for how to use this package for download video.
func ExampleStartDownload() {
	flag.Parse()
	log.Println(flag.Args())
	usr, _ := user.Current()
	currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	log.Println("download to dir=", currentDir)
	y := NewYoutube(true)
	arg := flag.Arg(0)
	if err := y.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}
	if err := y.StartDownload(filepath.Join(currentDir, "dl.mp4")); err != nil {
		fmt.Println("err:", err)
	}
}
