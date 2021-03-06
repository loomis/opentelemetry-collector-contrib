// Copyright 2020, OpenTelemetry Authors
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

package kubeletstatsreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configcheck"
	"go.opentelemetry.io/collector/config/configerror"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/testbed/testbed"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/kubelet"
)

func TestValidConfig(t *testing.T) {
	factory := &Factory{}
	err := configcheck.ValidateConfig(factory.CreateDefaultConfig())
	require.NoError(t, err)
}

func TestCreateTraceReceiver(t *testing.T) {
	factory := &Factory{}
	traceReceiver, err := factory.CreateTraceReceiver(
		context.Background(),
		zap.NewNop(),
		factory.CreateDefaultConfig(),
		nil,
	)
	require.Equal(t, err, configerror.ErrDataTypeIsNotSupported)
	require.Nil(t, traceReceiver)
}

func TestCreateMetricsReceiver(t *testing.T) {
	factory := &Factory{
		restClient: func(*zap.Logger, configmodels.Receiver) (kubelet.RestClient, error) {
			return &fakeRestClient{}, nil
		},
	}
	metricsReceiver, err := factory.CreateMetricsReceiver(
		zap.NewNop(),
		factory.CreateDefaultConfig(),
		&testbed.MockMetricConsumer{},
	)
	require.NoError(t, err)
	require.NotNil(t, metricsReceiver)
}
