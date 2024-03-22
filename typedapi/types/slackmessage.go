// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SlackMessage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339/specification/watcher/_types/Actions.ts#L130-L137
type SlackMessage struct {
	Attachments        []SlackAttachment       `json:"attachments"`
	DynamicAttachments *SlackDynamicAttachment `json:"dynamic_attachments,omitempty"`
	From               string                  `json:"from"`
	Icon               *string                 `json:"icon,omitempty"`
	Text               string                  `json:"text"`
	To                 []string                `json:"to"`
}

func (s *SlackMessage) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "attachments":
			if err := dec.Decode(&s.Attachments); err != nil {
				return fmt.Errorf("%s | %w", "Attachments", err)
			}

		case "dynamic_attachments":
			if err := dec.Decode(&s.DynamicAttachments); err != nil {
				return fmt.Errorf("%s | %w", "DynamicAttachments", err)
			}

		case "from":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "From", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.From = o

		case "icon":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Icon", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Icon = &o

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Text", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = o

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return fmt.Errorf("%s | %w", "To", err)
			}

		}
	}
	return nil
}

// NewSlackMessage returns a SlackMessage.
func NewSlackMessage() *SlackMessage {
	r := &SlackMessage{}

	return r
}
