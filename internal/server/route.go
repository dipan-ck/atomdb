package server

import (
	"fmt"
)

func HandleCommand(cmdString []string, user *client) {

	if len(cmdString) == 0 {
		user.conn.Write([]byte("-ERR empty command\r\n"))
		return
	}

	cmd := cmdString[0]

	switch cmd {
	case "SET":
		if len(cmdString) < 3 {
			user.conn.Write([]byte("-ERR wrong number of arguments for 'SET'\r\n"))
			return
		} else {
			setSuccess := SetKey(user, cmdString[1], cmdString[2])

			if setSuccess {
				user.conn.Write([]byte("+OK\r\n"))
				return
			} else {
				user.conn.Write([]byte("unable to set key"))
			}
		}

	case "GET":
		if len(cmdString) < 2 {
			user.conn.Write([]byte("-ERR wrong number of arguments for 'SET'\r\n"))
			return
		} else {
			val, ok := GetKey(user, cmdString[1])

			if ok {
				user.conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)))
				return
			} else {
				user.conn.Write([]byte("$-1\r\n"))
				return
			}
		}

	default:
		user.conn.Write([]byte(fmt.Sprintf("-ERR unknown command '%s'\r\n", cmdString[0])))

	}

}
