// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jaegerlegacyreceiver

import (
	"go.opentelemetry.io/collector/config/configmodels"
	jaegercore "go.opentelemetry.io/collector/receiver/jaegerreceiver"
)

// The config field name to load the protocol map from
const protocolsFieldName = "protocols"

// Config defines configuration for Jaeger receiver.
type Config struct {
	TypeVal   configmodels.Type                    `mapstructure:"-"`
	NameVal   string                               `mapstructure:"-"`
	Protocols map[string]*jaegercore.SecureSetting `mapstructure:"protocols"`
}

// Name gets the receiver name.
func (rs *Config) Name() string {
	return rs.NameVal
}

// SetName sets the receiver name.
func (rs *Config) SetName(name string) {
	rs.NameVal = name
}

// Type sets the receiver type.
func (rs *Config) Type() configmodels.Type {
	return rs.TypeVal
}

// SetType sets the receiver type.
func (rs *Config) SetType(typeStr configmodels.Type) {
	rs.TypeVal = configmodels.Type(typeStr)
}
