/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package threshold

import (
	"fmt"
	"strings"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	// Name of plugin
	Name       = "threshold"
	// Version of plugin
	Version    = 1
)

// ThresholdProcessor structure
type ThresholdProcessor struct {
}

// NewThresholdProcessor constructor
func NewThresholdProcessor() *ThresholdProcessor {
	return &ThresholdProcessor{}
}

// Process method to filter data
func (p ThresholdProcessor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	for _, m := range mts {
		key := fmt.Sprintf("/%s", strings.Join(m.Namespace.Strings(), "/"))
		if val, ok := cfg[key]; ok {
			a := convertIntervace(m.Data)
			b := convertIntervace(val)
			if a >= b {
				metrics = append(metrics, m)
			}
		}
	}
	return metrics, nil
}

// GetConfigPolicy returns plugin config
func (p ThresholdProcessor) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}

func convertIntervace(data interface{}) float64 {
	switch data.(type) {
	case int:
		return float64(data.(int))
	case int8:
		return float64(data.(int8))
	case int16:
		return float64(data.(int16))
	case int32:
		return float64(data.(int32))
	case int64:
		return float64(data.(int64))
	case float32:
		return float64(data.(float32))
	case float64:
		return float64(data.(float64))
	default:
		return float64(0)
	}
}
