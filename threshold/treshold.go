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
	dynamic := map[string][]string{}
	for k, _ := range cfg {
		if strings.Contains(k, "*") {
			sl := strings.Split(k, "/")
			if sl[0] == "" {
				sl = append(sl[:0], sl[1:]...)
			}
			dynamic[k] = sl
		}
	}
	for _, m := range mts {
		if dynamic != nil {
			for k, v := range dynamic {
				counter := 0
				ignored := 0
				for i, ns := range v {
					if ns == "*" {
						ignored++
					}
					if m.Namespace.Strings()[i] == ns {
						counter++
					}
				}
				if counter == len(v) - ignored {
					a := convertInterface(m.Data)
					b := convertInterface(cfg[k])
					if a >= b {
						metrics = append(metrics, m)
					}
				}
			}
		}
		key := fmt.Sprintf("/%s", strings.Join(m.Namespace.Strings(), "/"))
		if val, ok := cfg[key]; ok {
			a := convertInterface(m.Data)
			b := convertInterface(val)
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

func convertInterface(data interface{}) float64 {
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
