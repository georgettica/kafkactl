// Copyright © 2018 Dirk Wilden <dirk.wilden@device-insight.com>
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

package main

import (
	"os"

	"github.com/deviceinsight/kafkactl/cmd"
	"github.com/deviceinsight/kafkactl/output"
)

func main() {

	ioStreams := output.DefaultIOStreams()

	if err := cmd.NewKafkactlCommand(ioStreams).Execute(); err != nil {
		output.Warnf("%v", err)
		os.Exit(1)
	}
}
