package resp

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
)

func RespParsing(input []byte) ([]string, error) {

	reader := bufio.NewReader(bytes.NewReader(input))

	prefix, err := reader.ReadByte()

	if err != nil {
		return nil, err
	}

	if prefix != '*' {
		return nil, errors.New("expected * at beginning of RESP array")
	}

	arrayCount, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	arrayCount = trimLineEnd(arrayCount)

	count, err := strconv.Atoi(arrayCount)

	if err != nil {
		return nil, err
	}

	result := make([]string, 0, count)

	for i := 0; i < count; i++ {
		prefix, err := reader.ReadByte()

		if err != nil {
			return nil, err
		}

		if prefix != '$' {
			return nil, errors.New("expected '$' for bulk string")
		}

		lengthLine, err := reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		lengthLine = trimLineEnd(lengthLine)
		strCount, err := strconv.Atoi(lengthLine)

		if err != nil {
			return nil, err
		}

		strBytes := make([]byte, strCount+2)

		_, err = reader.Read(strBytes)

		if err != nil {
			return nil, err
		}

		result = append(result, string(strBytes[0:strCount]))

	}

	return result, nil

}

func trimLineEnd(s string) string {
	// First remove \r\n if present
	bs := []byte(s)
	bs = bytes.TrimSuffix(bs, []byte("\r\n"))
	// Then remove just \n if present (covers all cases)
	bs = bytes.TrimSuffix(bs, []byte("\n"))
	return string(bs)
}
