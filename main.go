package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"net"
	"os"
	"strings"

	"github.com/kevinburke/hostsfile/lib"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setAll(h *hostsfile.Hostsfile, url string) {
	local, err := net.ResolveIPAddr("ip", "127.0.0.1")
	checkError(err)

	ipv6local, err := net.ResolveIPAddr("ip", "fe80::1%lo0")
	checkError(err)

	h.Set(*local, url)
	h.Set(*ipv6local, url)
}

func main() {
	file := flag.String("file", "hostsfile", "File with newline-separated domain names")
	flag.Parse()
	f, err := os.Open(*file)
	checkError(err)
	scanner := bufio.NewScanner(f)
	h, err := hostsfile.Decode(bytes.NewBuffer([]byte{}))
	checkError(err)
	for scanner.Scan() {
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		setAll(&h, line)
	}
	err = hostsfile.Encode(os.Stdout, h)
	checkError(err)
}
