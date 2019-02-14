// Copyright © 2018 Intel Corporation
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
	"fmt"
	"github.com/clearlinux/mixer-tools/swupd"
)

func main() {
	files := []string{"./Manifest.mixer", "./Manifest.kernel-native-dkms.I.27330"}
	var manifests []*swupd.Manifest

	for _, manifest := range files {
		fmt.Printf("\tblovin: Parsing manifest %s\n", manifest)
		mf, err := swupd.ParseManifestFile(manifest)
		if err != nil {
			fmt.Printf("Parsing Error: ")
			fmt.Println(err)
			break
		}
		manifests = append(manifests, mf)
	}
}
