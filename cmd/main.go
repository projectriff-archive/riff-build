/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/projectriff/riff-init/pkg/initializer"
)

var (
	functionInvoker string
	initOptions     initializer.InitOptions
)

func main() {
	invoker, err := initializer.LoadInvoker(functionInvoker)
	if err != nil {
		glog.Fatalf("Unable to get invoker: %v", err)
	}
	err = initializer.Initialize(invoker, &initOptions)
	if err != nil {
		glog.Fatalf("Unable initialize function: %v", err)
	}
}

func init() {
	initOptions = initializer.InitOptions{}

	flag.StringVar(&functionInvoker, "invoker", "", "Path to the riff invoker.yaml, may be a local path or http")
	flag.StringVar(&initOptions.FunctionName, "name", "", "The name of the function")
	flag.StringVar(&initOptions.Artifact, "artifact", "", "Path to the function artifact, source code or jar file (attempts detection if not specified)")
	flag.StringVar(&initOptions.Handler, "handler", "", "Name of method or class to invoke (see specific invoker for detail)")
	flag.StringVar(&initOptions.FilePath, "filepath", "", "Path or directory used for the function (defaults to the current directory)")

	flag.Parse()
}
