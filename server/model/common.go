/*
 * Licensed to the LF AI & Data foundation under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 * //
 *     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"github.com/zilliztech/milvus-cdc/core/util"
	"net/url"
	"strconv"
	"strings"
)

//go:generate easytags $GOFILE json,mapstructure
type MilvusConnectParam struct {
	URI             string          `json:"uri" mapstructure:"uri"`
	Token           string          `json:"token" mapstructure:"token"`
	Host            string          `json:"host" mapstructure:"host"`
	Port            int             `json:"port" mapstructure:"port"`
	Username        string          `json:"username,omitempty" mapstructure:"username,omitempty"`
	Password        string          `json:"password,omitempty" mapstructure:"password,omitempty"`
	EnableTLS       bool            `json:"enable_tls" mapstructure:"enable_tls"`
	DialConfig      util.DialConfig `json:"dial_config" mapstructure:"dial_config"`
	IgnorePartition bool            `json:"ignore_partition" mapstructure:"ignore_partition"`
	// ConnectTimeout unit: s
	ConnectTimeout int `json:"connect_timeout" mapstructure:"connect_timeout"`
}

func (p *MilvusConnectParam) ParseURI() error {
	if p.URI != "" {
		parsedURL, err := url.Parse(p.URI)
		if err != nil {
			return err
		}

		// parse scheme
		if parsedURL.Scheme == "https" {
			p.EnableTLS = true
		} else {
			p.EnableTLS = false
		}

		// parse host
		p.Host = parsedURL.Host

		// parse host
		if parsedURL.Port() == "" {
			p.Port = 443
		} else {
			p.Port, err = strconv.Atoi(parsedURL.Port())
			if err != nil {
				return err
			}
		}

		//// parse database
		//if parsedURL.Path == "" || parsedURL.Path == "/" {
		//	p.Database = "default"
		//} else {
		//	p.Database = strings.Substring(1)
		//}
	}

	if p.Token != "" {
		splitStrings := strings.Split(p.Token, ":")
		p.Username = splitStrings[0]
		p.Password = splitStrings[1]
	}

	return nil
}

type CollectionInfo struct {
	Name      string            `json:"name" mapstructure:"name"`
	Positions map[string]string `json:"positions" mapstructure:"positions"`
}

type ChannelInfo struct {
	Name     string `json:"name" mapstructure:"name"`
	Position string `json:"position" mapstructure:"position"`
}

type BufferConfig struct {
	Period int `json:"period" mapstructure:"period"`
	Size   int `json:"size" mapstructure:"size"`
}

const (
	// TmpCollectionID which means it's the user custom collection position
	TmpCollectionID int64 = -1
	// TmpCollectionName TODO if replicate the rbac info, the collection id will be set it
	TmpCollectionName = "-1"

	ReplicateCollectionID   int64 = -10
	ReplicateCollectionName       = "-10"
)
