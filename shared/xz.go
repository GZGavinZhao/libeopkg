//
// Copyright © 2017-2020 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package shared

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/ulikunitz/xz"
)

// XzFile compress the file given and leave a ".xz" suffixed file in place.
//
// keepOriginal determins whether we'll keep the original file.
func XzFile(inputPath string, keepOriginal bool) error {
	input, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	r := bufio.NewReader(input)

	data, err := ioutil.ReadAll(r)

	output, err := os.Create(inputPath + ".xz")
	if err != nil {
		return err
	}

	w, err := xz.NewWriter(output)
	if err != nil {
		return err
	}

	w.Write(data)

	err = w.Close()
	if err != nil {
		return err
	}

	if !keepOriginal {
		err := os.Remove(inputPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// XzFileCmd is a simple wrapper around the xz utility to compress the input
// file. This will be performed in place and leave a ".xz" suffixed file in
// place.
//
// Keep original determines whether we'll keep the original file.
func XzFileCmd(inputPath string, keepOriginal bool) error {
	cmd := []string{
		"xz",
		"-6",
		"-T", "2",
		inputPath,
	}
	if keepOriginal {
		cmd = append(cmd, "-k")
	}
	c := exec.Command(cmd[0], cmd[1:]...)
	return c.Run()
}

// UnxzFile will decompress the input XZ file and leave a new file in place
// without the .xz suffix
func UnxzFile(inputPath string, keepOriginal bool) error {
	cmd := []string{
		"unxz",
		"-T", "2",
		inputPath,
	}
	if keepOriginal {
		cmd = append(cmd, "-k")
	}
	c := exec.Command(cmd[0], cmd[1:]...)
	return c.Run()
}
