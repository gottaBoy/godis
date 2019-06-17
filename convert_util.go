package godis

import (
	"encoding/binary"
	"math"
	"strconv"
)

func IntToByteArray(a int) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, int64(a), 10)
}

func Int64ToByteArray(a int64) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, a, 10)
}

func Float64ToByteArray(a float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(a))
	return buf[:]
}

func StringStringArrayToByteArray(str string, arr []string) [][]byte {
	params := make([][]byte, 0)
	params = append(params, []byte(str))
	for _, v := range arr {
		params = append(params, []byte(v))
	}
	return params
}

func StringArrayToByteArray(arr []string) [][]byte {
	newArr := make([][]byte, len(arr))
	for _, a := range newArr {
		newArr = append(newArr, []byte(a))
	}
	return newArr
}

func StringToFloat64Reply(reply string, err error) (float64, error) {
	f, e := strconv.ParseFloat(reply, 64)
	if e != nil {
		return 0, e
	}
	return f, err
}

func StringArrayToMapReply(reply []string, err error) (map[string]string, error) {
	newMap := make(map[string]string, len(reply)/2)
	for i := 0; i < len(reply); i += 2 {
		newMap[reply[i]] = reply[i+1]
	}
	return newMap, err
}

func Int64ToBoolReply(reply int64, err error) (bool, error) {
	return reply == 1, err
}

func ByteToStringReply(reply []byte, err error) (string, error) {
	return string(reply), err
}

func StringArrToTupleReply(reply []string, err error) ([]Tuple, error) {
	if len(reply) == 0 {
		return []Tuple{}, nil
	}
	newArr := make([]Tuple, len(reply)/2)
	for i := 0; i < len(reply); i += 2 {
		f, err := strconv.ParseFloat(reply[i+1], 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, Tuple{element: []byte(reply[i]), score: f})
	}
	return newArr, err
}
