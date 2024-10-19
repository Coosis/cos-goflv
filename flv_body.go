package cosgoflv

import (
	"io"
	"fmt"
	"encoding/binary"
)

type TagnSize struct {
	Tag FlvTag
	PrevSize uint32
}

func(ts *TagnSize) Write(w io.Writer) error {
	if err := ts.Tag.Write(w); err != nil {
		return err
	}
	ts.PrevSize = 11 + uint32(len(ts.Tag.Data))
	if err := binary.Write(w, binary.BigEndian, ts.PrevSize); err != nil {
		return err
	}
	return nil
}

func(ts *TagnSize) Read(r io.Reader) error {
	if err := ts.Tag.Read(r); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &ts.PrevSize); err != nil {
		return err
	}
	return nil
}

func EmptyFlvBody() *FlvBody {
	return &FlvBody{
		Tag: []TagnSize{},
	}
}

type FlvBody struct {
	Tag []TagnSize
}

func(b *FlvBody) Write(w io.Writer) error {
	// First tag size is always 0
	if err := binary.Write(w, binary.BigEndian, uint32(0)); err != nil {
		return err
	}
	// Write all tags
	for _, tag := range b.Tag {
		if err := tag.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func(b *FlvBody) Read(r io.Reader) error {
	var tagSize0 uint32
	// Start is always 0 in UI32
	if err := binary.Read(r, binary.BigEndian, &tagSize0); err != nil {
		return err
	}
	if tagSize0 != 0 {
		return fmt.Errorf("First tag size is not 0")
	}
	for {
		tag := TagnSize{}
		if err := tag.Read(r); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		b.Tag = append(b.Tag, tag)
	}
	return nil
}

func(b *FlvBody) AddTag(tags ...FlvTag) {
	for _, tag := range tags {
		b.Tag = append(b.Tag, TagnSize{
			Tag: tag,
		})
	}
}
