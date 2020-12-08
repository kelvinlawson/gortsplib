// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/gortsplib"
)

// This example shows how to
// 1. connect to a RTSP server
// 2. read all tracks on a path

func main() {
	// connect to the server and start reading all tracks
	conn, err := gortsplib.DialRead("rtsp://localhost:8554/mystream")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// read track frames
	done := conn.ReadFrames(func(id int, typ gortsplib.StreamType, buf []byte) {
		fmt.Printf("frame from track %d, type %v: %v\n", id, typ, buf)
	})

	// catch any error
	err = <-done
	panic(err)
}
