package main

import (
	"fmt"
	. "github.com/Coosis/cos-goflv"
)

func printflv(flv *Flv) {
	fmt.Println("Header:")
	fmt.Printf("\tSignature: FLV\n")
	fmt.Printf("\tVersion: %d\n", flv.Header.Version)
	fmt.Printf("\tFlags: %d\n", flv.Header.TypeFlags)
	fmt.Printf("\tDataOffset: %d\n", flv.Header.DataOffset)

	fmt.Println("Tags:")
	for i, tag := range flv.Body.Tag {
		fmt.Printf("\tTag %d:\n", i)
		fmt.Printf("\t\tTagType: %d\n", tag.Tag.TagType)
		if tag.Tag.TagType == SCRIPTDATA {
			fmt.Printf("\t\tA script data tag!\n")
		}
		fmt.Printf("\t\tDataSize: %d\n", tag.Tag.DataSize)
		fmt.Printf("\t\tTimestamp: %d\n", tag.Tag.TimeStamp)
		fmt.Printf("\t\tStreamID: %d\n", tag.Tag.StreamID)
		fmt.Println()
	}

	fmt.Println("End of FLV")
}
