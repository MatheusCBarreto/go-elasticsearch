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
// https://github.com/elastic/elasticsearch-specification/tree/6e0fb6b929f337b62bf0676bdf503e061121fad2

package putindextemplate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putindextemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/6e0fb6b929f337b62bf0676bdf503e061121fad2/specification/indices/put_index_template/IndicesPutIndexTemplateRequest.ts#L36-L95
type Request struct {

	// ComposedOf An ordered list of component template names.
	// Component templates are merged in the order specified, meaning that the last
	// component template specified has the highest precedence.
	ComposedOf []string `json:"composed_of,omitempty"`
	// DataStream If this object is included, the template is used to create data streams and
	// their backing indices.
	// Supports an empty object.
	// Data streams require a matching index template with a `data_stream` object.
	DataStream *types.DataStreamVisibility `json:"data_stream,omitempty"`
	// IndexPatterns Name of the index template to create.
	IndexPatterns []string `json:"index_patterns,omitempty"`
	// Meta_ Optional user metadata about the index template.
	// May have any contents.
	// This map is not automatically generated by Elasticsearch.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// Priority Priority to determine index template precedence when a new data stream or
	// index is created.
	// The index template with the highest priority is chosen.
	// If no priority is specified the template is treated as though it is of
	// priority 0 (lowest priority).
	// This number is not automatically generated by Elasticsearch.
	Priority *int `json:"priority,omitempty"`
	// Template Template to be applied.
	// It may optionally include an `aliases`, `mappings`, or `settings`
	// configuration.
	Template *types.IndexTemplateMapping `json:"template,omitempty"`
	// Version Version number used to manage index templates externally.
	// This number is not automatically generated by Elasticsearch.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putindextemplate request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "composed_of":
			if err := dec.Decode(&s.ComposedOf); err != nil {
				return err
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return err
			}

		case "index_patterns":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.IndexPatterns = append(s.IndexPatterns, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.IndexPatterns); err != nil {
					return err
				}
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "priority":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Priority = &value
			case float64:
				f := int(v)
				s.Priority = &f
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}
