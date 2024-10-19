package cosgoflv

import (
	"fmt"
)

const (
	SCRIPTDATAOBJECTEND uint8 = 0x09
)

type FlvScriptDataString struct {
	StringLength uint16
	StringData string
}
func(dataString *FlvScriptDataString) Parse(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, fmt.Errorf("Invalid FlvScriptDataString. Data length is less than 2")
	}
	dataString.StringLength = uint16(data[0]) << 8 | uint16(data[1])
	if len(data) < int(dataString.StringLength) + 2 {
		fmt.Println(dataString.StringLength)
		fmt.Println(len(data))
		return 0, fmt.Errorf("Invalid FlvScriptDataString. Data length is less than StringLength")
	}
	dataString.StringData = string(data[2:2+dataString.StringLength])
	return int(dataString.StringLength + 2), nil
}

type FlvScriptDataLongString struct {
	StringLength uint32
	StringData string
}
func(dataString *FlvScriptDataLongString) Parse(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, fmt.Errorf("Invalid FlvScriptDataString. Data length is less than 2")
	}
	dataString.StringLength = uint32(data[0]) << 24 | uint32(data[1]) << 16 | uint32(data[2]) << 8 | uint32(data[3])
	if len(data) < int(dataString.StringLength) + 4 {
		// fmt.Println(dataString.StringLength)
		// fmt.Println(len(data))
		return 0, fmt.Errorf("Invalid FlvScriptDataString. Data length is less than StringLength")
	}
	dataString.StringData = string(data[4:4+dataString.StringLength])
	return int(dataString.StringLength + 4), nil
}


type FlvScriptDataVariable struct {
	VariableName FlvScriptDataString
	VariableData FlvScriptDataValue
}

type FlvScriptDataVariableEnd struct {
	// Always 9
	VariableEndMarker1 uint32
}

type FlvScriptDataDate struct {
	// Number of milliseconds since midnight, January 1, 1970 UTC
	DateTime float64
	// Local time offset in minutes from UTC
	LocalDateTimeOffset int16
}

type FlvScriptDataValue struct {
	DataType uint8
	ECMAArrayLength uint32
	ScriptDataValue []byte
}
func(val *FlvScriptDataValue) Parse(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 1")
	}
	val.DataType = data[0]
	// Empty ECMA array
	if len(data) == 1 {
		val.ECMAArrayLength = 0
		val.ScriptDataValue = nil
		return 1, nil
	}

	if val.DataType == 0 {
		if len(data) < 9 {
			return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 9")
		}
		val.ScriptDataValue = data[1:9]
		return 9, nil
	}

	if val.DataType == 1 {
		if len(data) < 2 {
			return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 2")
		}
		val.ScriptDataValue = data[1:2]
	}

	if val.DataType == 2 || val.DataType == 4 {
		dataString := &FlvScriptDataString{}
		cnt, err := dataString.Parse(data[1:])
		if err != nil {
			return 0, err
		}
		val.ScriptDataValue = data[1:cnt]
		return cnt + 1, nil
	}

	if val.DataType == 5 || val.DataType == 6 {
		val.ScriptDataValue = nil
		return 1, nil
	}

	if val.DataType == 7 {
		if len(data) < 3 {
			return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 3")
		}
		val.ScriptDataValue = data[1:3]
		return 3, nil
	}

	if val.DataType == 11 {
		if len(data) < 11 {
			return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 11")
		}
		val.ScriptDataValue = data[1:11]
		return 11, nil
	}

	if val.DataType == 12 {
		dataString := &FlvScriptDataLongString{}
		cnt, err := dataString.Parse(data[1:])
		if err != nil {
			return 0, err
		}
		val.ScriptDataValue = data[1:cnt]
		return cnt + 1, nil
	}

	if val.DataType == 10 {
		cnt := uint32(data[1]) << 24 | uint32(data[2]) << 16 | uint32(data[3]) << 8 | uint32(data[4])
		val.ScriptDataValue = data[1:(cnt + 5)]
		return int(cnt + 5), nil
	}

	cnt := 0
	for {
		if len(data) < 3 {
			return 0, fmt.Errorf("Invalid FlvScriptDataValue. Data length is less than 1")
		}
		if data[0] == 0 && data[1] == 0 && data[2] == SCRIPTDATAOBJECTEND {
			cnt += 3
			break
		}
		data = data[1:]
		cnt++
	}

	return 1+cnt, nil
}

type FlvScriptDataObject struct {
	ObjectName FlvScriptDataString
	ObjectData FlvScriptDataValue
}
func(obj *FlvScriptDataObject) Parse(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, fmt.Errorf("Invalid FlvScriptDataObject. Data length is less than 1")
	}

	cnt, err := obj.ObjectName.Parse(data)
	if err != nil {
		return 0, err
	}
	if len(data) <= int(cnt) {
		return 0, fmt.Errorf("Invalid FlvScriptDataObject. Data length is less than ObjectName length")
	}
	cnt2, err := obj.ObjectData.Parse(data[cnt:])
	if err != nil {
		return 0, err
	}
	return cnt + cnt2, nil
}
