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
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/ghodss/yaml"
	projectriff_v1 "github.com/projectriff/riff-init/pkg/apis/projectriff.io/v1alpha1"
	"github.com/projectriff/riff-init/pkg/osutils"
)

func LoadInvoker(path string) (*projectriff_v1.Invoker, error) {
	invokerURLs, err := resolveInvokerURLs(path)
	if err != nil {
		return nil, err
	}
	if len(invokerURLs) != 1 {
		return nil, fmt.Errorf("found multiple matches for invoker path: %s", path)
	}
	invokerBytes, err := loadInvoker(invokerURLs[0])

	invoker := projectriff_v1.Invoker{}
	err = yaml.Unmarshal(invokerBytes, &invoker)
	if err != nil {
		return nil, err
	}
	return &invoker, nil
}

func loadInvoker(url url.URL) ([]byte, error) {
	if url.Scheme == "file" {
		file, err := ioutil.ReadFile(url.Path)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else if url.Scheme == "http" || url.Scheme == "https" {
		resp, err := http.Get(url.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, fmt.Errorf("Filename must be file, http or https, got %s", url.Scheme)
}

func resolveInvokerURLs(filename string) ([]url.URL, error) {
	u, err := url.Parse(filename)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "file"
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return []url.URL{*u}, nil
	}
	if u.Scheme == "file" {
		if osutils.IsDirectory(u.Path) {
			u.Path = filepath.Join(u.Path, "*-invoker.yaml")
		}
		filenames, err := filepath.Glob(u.Path)
		if err != nil {
			return nil, err
		}
		var urls = []url.URL{}
		for _, f := range filenames {
			urls = append(urls, url.URL{
				Scheme: u.Scheme,
				Path:   f,
			})
		}
		return urls, nil
	}
	return nil, fmt.Errorf("Filename must be file, http or https, got %s", u.Scheme)
}
