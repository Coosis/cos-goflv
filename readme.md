# A go lib for handling FLV files
## Warning: This lib only handles flv with the following conditions met:
- Valid flv file
- Flv version is 1
- AAC audio(if parsing SoundData)
- AVC video(if parsing VideoData)
## Warning: The lib does not support parsing data tags. Relevant code is in flv_data_tag.go and flv_scriptdata.go

# Quick Start
```go
package main

import (
    "os"
    . "github.com/Coosis/cos-goflv"
)

func main() {
    // Reading an flv
	// Place the sample.flv file in root dir
	f, err := os.Open("sample.flv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	flv := Flv{}
	if err := flv.Read(f); err != nil {
		panic(err)
	}

	// ...

	// Writing an flv
	// Will create a copy of the sample.flv file in root dir
	f2, err := os.Create("sample_cp.flv")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if err := flv.Write(f2); err != nil {
		panic(err)
	}

}
```

# For better integration with golang, the following types(right) are used for types(left) in the flv spec:
UB\[2\] -> uint32
UB\[4\] -> uint8
UI24 -> uint32
SI24 -> int32

# If not parsed, data will be stored in the Data field of the Tag struct
To have more control over the data, parse them into audio/video tag using `TryIntoAudio` and `TryIntoVideo`
Example:
```go
package main

import (
    "os"
    . "github.com/Coosis/cos-goflv"
)

func main() {
    // Reading an flv
	// Place the sample.flv file in root dir
	f, err := os.Open("sample.flv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	flv := Flv{}
	if err := flv.Read(f); err != nil {
		panic(err)
	}

    audio := &FlvAudioTag{}
    if err := flv.Body.Tags[0].Tag.TryIntoAudio(audio); err != nil {
        panic(err)
    }

    video := &FlvVideoTag{}
    if err := flv.Body.Tags[1].Tag.TryIntoVideo(video); err != nil {
        panic(err)
    }
}
```
