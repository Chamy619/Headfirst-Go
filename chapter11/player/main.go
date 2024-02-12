package main

import "headfirst-go/chapter11/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TyrOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var player Player
	player = gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixTape)
	TyrOut(player)
	player = gadget.TapeRecorder{}
	playList(player, mixTape)
	TyrOut(player)
}