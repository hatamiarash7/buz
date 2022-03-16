package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
	"github.com/silverton-io/honeypot/pkg/event"
)

const (
	JSON_CONTENT_TYPE string = "application/json"
)

func PostEvent(url url.URL, payload event.SelfDescribingEvent) {
	data, _ := json.Marshal(payload)
	buff := bytes.NewBuffer(data)
	_, err := http.Post(url.String(), JSON_CONTENT_TYPE, buff)
	if err != nil {
		log.Trace().Err(err).Msg("could not send payload to " + url.String())
	}
}

func Get(url url.URL) (body []byte, err error) {
	resp, err := http.Get(url.String())
	if err != nil {
		log.Trace().Err(err).Msg("could not get url " + url.String())
		return nil, err
	}
	defer resp.Body.Close()
	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		log.Trace().Err(ioerr).Msg("could not read response body")
		return nil, ioerr
	}
	return body, nil
}
