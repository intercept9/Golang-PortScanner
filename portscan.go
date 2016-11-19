//
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Intercept9.
//
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package main

import (
	//"log"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {

		usage() //he fucked with arguments tell him how to do it
	} else {

		callme()
	}
}

func usage() {

	fmt.Printf(`
==============================Usage============================
			
		go build portscan.go
		  ./portscan ip
	# example > ./portscan 127.0.0.1
 	
	`)
}

func callme() {

	var status string
	getip := os.Args[1]
	getport := 0
	for i := 20; i < 65536; i++ {
		getport = i
		data := getip + ":" + strconv.Itoa(getport)
		conn, err := net.Dial("tcp", data)
		if err != nil {
			//log.Println("Connection error:", err)     //When want a specific reason why it dint connect
			status = strconv.Itoa(getport) + " - " + "Closed"
		} else {

			status = strconv.Itoa(getport) + " - " + "Open"
			conn.Close()

			fmt.Println(status)
			fmt.Printf("ok")

		}

	}

}
