package main

import "SONG"

func main() {
	//nil for dm will use the internal vidi mplmentation
	s := SONG.InitServer(".", ":8080", "veni", "vidi", "vici", nil)
	s.Serve()
}
