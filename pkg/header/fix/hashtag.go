// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package fix

import (
	"io/ioutil"
	"license-checker/pkg/header"
	"os"
	"reflect"
	"strings"
)

// Hashtag adds the configured license header to the files whose comment starts with #.
func Hashtag(file string, config *header.Config, result *header.Result) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if len(content) >= 3 && !reflect.DeepEqual(content[0:3], []byte("#! ")) || len(content) < 3 { // doesn't contains shebang
		lines := "# " + strings.Join(strings.Split(config.License, "\n"), "\n# ") + "\n"

		if err := ioutil.WriteFile(file, append([]byte(lines), content...), stat.Mode()); err != nil {
			return err
		}

		result.Fix(file)
	}

	return nil
}
