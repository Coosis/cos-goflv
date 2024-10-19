package cosgoflv

import (
	// "fmt"
)

const (
	OBJECTENDMET = "SCRIPTDATAOBJECTEND"
)

type FlvDataTag struct {
	Objects []FlvScriptDataObject
	End uint32
}

// func(tag *FlvDataTag) Parse(data []byte) error {
	// for {
	// 	if len(data) < 3 {
	// 		return fmt.Errorf("Invalid FlvDataTag. Data length is less than 3")
	// 	}
	// 
	// 	if data[0] == 0 && data[1] == 0 && data[2] == SCRIPTDATAOBJECTEND {
	// 		break
	// 	}
	//
	// 	object := &FlvScriptDataObject{}
	// 	cnt, err := object.Parse(data)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	tag.Objects = append(tag.Objects, *object)
	// 	data = data[cnt:]
	// }
	// tag.End = uint32(SCRIPTDATAOBJECTEND)
	// return nil
	//
	// var offset int
 //    var err error
	//
 //    // Parse first AMF0 data type (should be type marker 0x02 for String)
 //    if len(data) <= offset {
 //        return fmt.Errorf("Insufficient data to read type marker")
 //    }
 //    dataType := data[offset]
 //    offset++
	//
 //    if dataType != 0x02 {
 //        return fmt.Errorf("Expected String type marker (0x02), got 0x%02X", dataType)
 //    }
	//
 //    // Parse Variable Name (String)
 //    objName := &FlvScriptDataString{}
 //    n, err := objName.Parse(data[offset:])
 //    if err != nil {
 //        return err
 //    }
 //    offset += n
	//
 //    // Parse second AMF0 data type (e.g., 0x08 for ECMA Array)
 //    if len(data) <= offset {
 //        return fmt.Errorf("Insufficient data to read second type marker")
 //    }
 //    dataType = data[offset]
 //    offset++
	//
 //    // Create a ScriptDataValue to parse the Variable Data
 //    objData := &FlvScriptDataValue{DataType: dataType}
 //    n, err = objData.ParseValue(data[offset:])
 //    if err != nil {
 //        return err
 //    }
 //    offset += n
	//
 //    // Store the parsed object
 //    tag.Objects = append(tag.Objects, FlvScriptDataObject{
 //        ObjectName: *objName,
 //        ObjectData: *objData,
 //    })
	//
 //    // Check for SCRIPTDATAOBJECTEND marker
 //    if len(data) >= offset+3 && data[offset] == 0x00 && data[offset+1] == 0x00 && data[offset+2] == 0x09 {
 //        tag.End = uint32(SCRIPTDATAOBJECTEND)
 //    } else {
 //        return fmt.Errorf("Expected SCRIPTDATAOBJECTEND marker at the end")
 //    }
	//
 //    return nil
// }
