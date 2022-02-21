package pprofutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	pprofProfile "github.com/00security/profefe/internal/pprof/profile"
)

type ProfileParserErrorCode string

const (
	ProfileParserErrorUndefined    ProfileParserErrorCode = "undefined"
	ProfileParserErrorProfileEmpty ProfileParserErrorCode = "profile_empty"
)

type ProfileParserError struct {
	err  error
	code ProfileParserErrorCode
}

func (e *ProfileParserError) Unwrap() error {
	return e.err
}

func (e *ProfileParserError) Error() string {
	return e.err.Error()
}

func (e *ProfileParserError) Code() ProfileParserErrorCode {
	return e.code
}

type ProfileParser struct {
	r    *bytes.Reader
	prof *pprofProfile.Profile
}

func NewProfileParser(data []byte) *ProfileParser {
	return &ProfileParser{r: bytes.NewReader(data)}
}

func (pr *ProfileParser) Read(p []byte) (n int, err error) {
	return pr.r.Read(p)
}

func (pr *ProfileParser) WriteTo(w io.Writer) (n int64, err error) {
	return pr.r.WriteTo(w)
}

func (pr *ProfileParser) Seek(offset int64, whence int) (int64, error) {
	return pr.r.Seek(offset, whence)
}

func (pr *ProfileParser) ParseProfile() (prof *pprofProfile.Profile, err error) {
	if pr.prof == nil {
		pr.prof, err = pprofProfile.Parse(pr.r)
	}
	if err != nil {
		return nil, &ProfileParserError{err, ProfileParserErrorUndefined}
	}
	if len(pr.prof.Sample) == 0 {
		return nil, &ProfileParserError{fmt.Errorf("profile is empty: no samples"), ProfileParserErrorProfileEmpty}
	}
	return pr.prof, nil
}

func ParseProfileFrom(r io.Reader) (*pprofProfile.Profile, error) {
	if pr, _ := r.(*ProfileParser); pr != nil {
		return pr.ParseProfile()
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewProfileParser(data).ParseProfile()
}
