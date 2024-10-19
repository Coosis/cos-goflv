package cosgoflv

import (
	"io"
)

type Flv struct {
	Header FlvHeader
	Body FlvBody
}

func EmptyFlv() *Flv {
	return &Flv{
		Header: *EmptyFlvHeader(),
		Body: *EmptyFlvBody(),
	}
}

func(f *Flv) Write(w io.Writer) error {
	if err := f.Header.Write(w); err != nil {
		return err
	}
	if err := f.Body.Write(w); err != nil {
		return err
	}

	return nil
}

func(f *Flv) Read(r io.Reader) error {
	if err := f.Header.Read(r); err != nil {
		return err
	}
	if err := f.Body.Read(r); err != nil {
		return err
	}

	return nil
}
