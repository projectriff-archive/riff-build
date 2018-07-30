/*
 * Copyright 2018 the original author or authors.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package initializer

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	projectriff_v1 "github.com/projectriff/riff-build/pkg/apis/projectriff.io/v1alpha1"
)

type resource struct {
	Path    string
	Content string
}

func generateResources(invoker projectriff_v1.Invoker, opts *InitOptions) error {
	var resources []resource

	// Invoker defined files
	for _, file := range invoker.Spec.Files {
		content, err := generateFileContents(file.Template, file.Path, *opts)
		if err != nil {
			return err
		}
		resources = append(resources, resource{
			Path:    file.Path,
			Content: content,
		})
	}

	workdir, err := filepath.Abs(opts.FilePath)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		err = writeFile(
			filepath.Join(workdir, resource.Path),
			resource.Content)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeFile(filename string, text string) error {
	fmt.Printf("Initializing %s\n", filename)
	return ioutil.WriteFile(filename, []byte(strings.TrimLeft(text, "\n")), 0644)
}
