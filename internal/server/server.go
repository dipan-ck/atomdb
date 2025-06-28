package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/dipan-ck/atomdb.git/internal/resp"
)

type client struct {
	conn            net.Conn
	reader          *bufio.Reader
	isAuthenticated bool
	remoteAddr      string
	secretKey       string
}

var authenticatedUser = make(map[string]string)

func StartServer(port string) error {

	conn, err := net.Listen("tcp", port)

	if err != nil {
		return err
	}

	fmt.Println("🚀 Redis Clone Server started on", port)

	for {
		userConn, err := conn.Accept()

		if err != nil {
			fmt.Println("❌ Failed to accept connection:", err)
			continue // 🔁 Do not exit the server loop!
		}

		TTLWatcher()
		go handleConnection(userConn)

	}

}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	user := &client{
		conn:            conn,
		reader:          bufio.NewReader(conn),
		isAuthenticated: false,
		remoteAddr:      conn.RemoteAddr().String(),
	}

	fmt.Println("📥 New client connected:", user.remoteAddr)

	for {
		rawInput, err := resp.ReadRESP(user.reader)

		if err != nil {
			fmt.Println("❌ Error reading RESP:", err)
			return err
		}

		parsedSlice, err := resp.RespParsing(rawInput)

		if err != nil {
			fmt.Println("❌ Failed to parse command:", err)
			return err
		}

		if !user.isAuthenticated {
			cmd := parsedSlice[0]

			if cmd == "AUTH" {

				SECRET := parsedSlice[1]

				exists, ok := authenticatedUser[SECRET]

				if ok && exists != user.remoteAddr {
					user.conn.Write([]byte("-ERR Secret already in use\r\n"))
				} else {
					authenticatedUser[SECRET] = user.remoteAddr
					user.isAuthenticated = true
					user.secretKey = SECRET
					user.conn.Write([]byte(fmt.Sprintf("+OK Registered & Authenticated with Secret: %s and RemoteAddr: %s\r\n", SECRET, user.remoteAddr)))
				}

			} else {
				user.conn.Write([]byte("-NOAUTH Please authenticate with AUTH <secret>\r\n"))
			}
		} else {
			fmt.Println("Parsed Command:", parsedSlice)
			HandleCommand(parsedSlice, user)
		}

	}

}
