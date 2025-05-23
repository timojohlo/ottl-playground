/*
 * Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
 * or more contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch B.V. licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	transformprocessorConfig = "transformprocessor.yaml"
)

func Test_TransformProcessorExecutor_ParseConfig(t *testing.T) {
	yamlConfig := readTestData(t, transformprocessorConfig)
	executor := NewTransformProcessorExecutor().(*transformProcessorExecutor)
	assert.NotNil(t, executor)
	parsedConfig, err := executor.parseConfig(yamlConfig)
	require.NoError(t, err)

	require.NotNil(t, parsedConfig)
	require.NotEmpty(t, parsedConfig.ErrorMode)
	require.NotEmpty(t, parsedConfig.TraceStatements)
	require.NotEmpty(t, parsedConfig.MetricStatements)
	require.NotEmpty(t, parsedConfig.MetricStatements)
}
