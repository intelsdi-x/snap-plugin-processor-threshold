//
// +build small

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestThresholdProcessor(t *testing.T) {
	proc := NewThresholdProcessor()
	Convey("Create threshold processor", t, func() {
		Convey("So proc should not be nil", func() {
			So(proc, ShouldNotBeNil)
		})
		Convey("So proc should be of type ThresholdProcessor", func() {
			So(proc, ShouldHaveSameTypeAs, &ThresholdProcessor{})
		})
		Convey("proc.GetConfigPolicy should return a config policy", func() {
			configPolicy, _ := proc.GetConfigPolicy()
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})

	Convey("Test Threshold Processor", t, func() {
		Convey("Process metrics with empty config and empty metrics", func() {
			config := plugin.Config{}
			metrics := []plugin.Metric{}
			mts, err := proc.Process(metrics, config)
			So(mts, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(mts, ShouldBeEmpty)
		})

		Convey("Process metrics with config containing dynamic namespace", func() {
			config := plugin.Config{
				"/test/*/test1/test11": 8,
			}
			metrics := []plugin.Metric{
				{
					Namespace: plugin.NewNamespace("test", "1", "test1", "test11"),
					Data:      0,
				},
				{
					Namespace: plugin.NewNamespace("test", "2", "test1", "test11"),
					Data:      8,
				},
				{
					Namespace: plugin.NewNamespace("test", "3", "test1", "test11"),
					Data:      10,
				},
			}

			mts, err := proc.Process(metrics, config)
			So(mts, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(mts, ShouldHaveLength, 2)
		})

		Convey("Process metrics with config containing static namespace", func() {
			config := plugin.Config{
				"/test/1/test1/test11": 8,
			}
			metrics := []plugin.Metric{
				{
					Namespace: plugin.NewNamespace("test", "1", "test1", "test11"),
					Data:      10,
				},
			}

			mts, err := proc.Process(metrics, config)
			So(mts, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(mts, ShouldHaveLength, 1)
		})

		Convey("Process metrics with config containing static namespace and filtered out value", func() {
			config := plugin.Config{
				"/test/1/test1/test11": 8,
			}
			metrics := []plugin.Metric{
				{
					Namespace: plugin.NewNamespace("test", "1", "test1", "test11"),
					Data:      6,
				},
			}

			mts, err := proc.Process(metrics, config)
			So(mts, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(mts, ShouldBeEmpty)
		})
	})

	Convey("Test convertInterface", t, func() {
		value := 42
		Convey("Input is type of int", func() {
			valInt := int(value)
			output := convertInterface(valInt)
			So(output, ShouldEqual, float64(valInt))
		})
		Convey("Input is type of int8", func() {
			valInt8 := int8(value)
			output := convertInterface(valInt8)
			So(output, ShouldEqual, float64(valInt8))
		})
		Convey("Input is type of int16", func() {
			valInt16 := int16(value)
			output := convertInterface(valInt16)
			So(output, ShouldEqual, float64(valInt16))
		})
		Convey("Input is type of int32", func() {
			valInt32 := int32(value)
			output := convertInterface(valInt32)
			So(output, ShouldEqual, float64(valInt32))
		})
		Convey("Input is type of int64", func() {
			valInt64 := int64(value)
			output := convertInterface(valInt64)
			So(output, ShouldEqual, float64(valInt64))
		})
		Convey("Input is type of float32", func() {
			valFloat32 := float32(value)
			output := convertInterface(valFloat32)
			So(output, ShouldEqual, float64(valFloat32))
		})
		Convey("Input is type of float64", func() {
			valFloat64 := float64(value)
			output := convertInterface(valFloat64)
			So(output, ShouldEqual, float64(valFloat64))
		})
		Convey("Input is not type of int and float", func() {
			valString := "input"
			output := convertInterface(valString)
			So(output, ShouldEqual, float64(0))
		})

	})

}
