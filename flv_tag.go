package cosgoflv

import (
	"io"
	"fmt"
)

type FlvTagType uint8
const (
	AUDIO FlvTagType = 8
	VIDEO FlvTagType = 9
	SCRIPTDATA FlvTagType = 18
	OTHER FlvTagType = 0
)
func IsValidFlvTagType(tagType FlvTagType) error {
	switch tagType {
	case AUDIO, VIDEO, SCRIPTDATA:
		return nil
	default:
		return fmt.Errorf("Invalid tag type.")
	}
}

type FlvTag struct {
	TagType FlvTagType
	DataSize uint32
	TimeStamp uint32
	StreamID uint32
	// Keep as-is, as it's not consumed by the parser
	Data []byte
}

func(tag *FlvTag) Write(w io.Writer) error {
	header := make([]byte, 11)
	header[0] = byte(tag.TagType)

	header[1] = byte(tag.DataSize >> 16)
	header[2] = byte(tag.DataSize >> 8)
	header[3] = byte(tag.DataSize)

	header[4] = byte(tag.TimeStamp >> 16)
	header[5] = byte(tag.TimeStamp >> 8)
	header[6] = byte(tag.TimeStamp)
	header[7] = byte(tag.TimeStamp >> 24)

	header[8] = byte(tag.StreamID >> 16)
	header[9] = byte(tag.StreamID >> 8)
	header[10] = byte(tag.StreamID)

	if _, err := w.Write(header); err != nil {
		return err
	}

	if _, err := w.Write(tag.Data); err != nil {
		return err
	}
	return nil
}

func(tag *FlvTag) Read(r io.Reader) error {
	buf := make([]byte, 11)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	tag.TagType = FlvTagType(buf[0])
	if err := IsValidFlvTagType(tag.TagType); err != nil {
		return err
	}
	tag.DataSize = uint32(buf[1]) << 16 | uint32(buf[2]) << 8 | uint32(buf[3])
	tag.TimeStamp = uint32(buf[4]) << 16 | uint32(buf[5]) << 8 | uint32(buf[6])
	tag.TimeStamp = uint32(buf[7]) << 24 | tag.TimeStamp
	tag.StreamID = uint32(buf[8]) << 16 | uint32(buf[9]) << 8 | uint32(buf[10])
	if tag.StreamID != 0 {
		return fmt.Errorf("Stream ID is always 0, found %d", tag.StreamID)
	}

	tag.Data = make([]byte, tag.DataSize)
	if _, err := io.ReadFull(r, tag.Data); err != nil {
		return err
	}

	return nil
}

func(tag *FlvTag) TryIntoAudio(audioTag *FlvAudioTag) error {
	if audioTag == nil {
		return fmt.Errorf("audioTag is nil.")
	}
	if tag.TagType != AUDIO {
		return fmt.Errorf("Not an audio tag.")
	}
	if err := audioTag.Parse(tag.Data); err != nil {
		return err
	}
	return nil
}

func(tag *FlvTag) TryIntoVideo(videoTag *FlvVideoTag) error {
	if videoTag == nil {
		return fmt.Errorf("videoTag is nil.")
	}
	if tag.TagType != VIDEO {
		return fmt.Errorf("Not a video tag.")
	}
	if err := videoTag.Parse(tag.Data); err != nil {
		return err
	}
	return nil
}
