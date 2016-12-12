package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"chapter0301/library"
)

var lib *library.MusicManager
var mp *library.MP3Player
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
		case "list":
			fmt.Println("List:")
			for i := 0; i < lib.Len(); i++ {
				e, _ := lib.Get(i)
				fmt.Println(i + 1, ":", e.Name, e.Artist, e.Source, e.Type)
			}
		case "add": {
			if len(tokens) == 6 {
				id++
				m0 := &library.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]}
				lib.Add(m0)
			} else {
				fmt.Println("USAGE: lib add <name><artist><source><type>")
			}
		}
		case "remove":
			if len(tokens) > 2 {
				lib.Remove(0)
			} else {
				fmt.Println("USAGE: lib remove <name>")
			}
		default:
			fmt.Println("Unknowed command", tokens[1])
	}
}

func handlePlayCommand(tokens []string)  {

	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("not exist")
		return
	}
	//mp.Play(e.Source, e.Type, ctrl, signal)
}


func main() {

	fmt.Println("================================")
	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized comman", tokens[0])
		}
	}
}
