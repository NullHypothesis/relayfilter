package main

import (
	"flag"
	"fmt"
	"log"

	"git.torproject.org/user/phw/zoossh.git"
)

func main() {

	nickname := flag.String("nickname", "", "Nickname.")
	data := flag.String("data", "", "File to analyse.")
	version := flag.String("version", "", "Tor version number.")
	flags := flag.String("flags", "", "Flags.")
	orport := flag.Int("orport", 0, "Onion routing port.")
	dirport := flag.Int("dirport", 0, "Directory port.")
	bandwidth := flag.Int("bandwidth", 0, "Advertised bandwidth.")

	flag.Parse()

	consensus, err := zoossh.ParseConsensusFile(*data)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for s := range consensus.Iterate() {

		status := s.(*zoossh.RouterStatus)

		if *nickname != "" && (status.Nickname != *nickname) {
			continue
		}

		if *version != "" && (status.TorVersion != *version) {
			continue
		}

		if *dirport != 0 && (status.DirPort != uint16(*dirport)) {
			continue
		}

		if *orport != 0 && (status.ORPort != uint16(*orport)) {
			continue
		}

		if *bandwidth != 0 && (status.Bandwidth != uint64(*bandwidth)) {
			continue
		}

		if *flags != "" && (status.Flags.String() != *flags) {
			continue
		}

		fmt.Println(status)
		count++
	}

	log.Printf("Filtered %d out of %d relays.\n", count, consensus.Length())
}
