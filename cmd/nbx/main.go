/*
Copyright 2017 Nanobox, Inc.

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

package main

import (
	"fmt"
	"os"

	"github.com/nanobox-io/nbx/cmd/nbx/app"
	"github.com/nanobox-io/nbx/cmd/nbx/app/env"
	"github.com/nanobox-io/nbx/cmd/nbx/app/log"
)

func main() {
	var err error

	// bootstrap the global app env
	app.Env = env.NewSystemEnv()
	if err := app.Env.Bootstrap(); err != nil {
		fmt.Fprintf(os.Stderr, "error: Failed to bootstrap app env: %v\n", err)
		os.Exit(1)
	}

	// set the global dir
	app.GlobalDir, err = app.Env.GlobalDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: Unable to determine global dir for app: %v\n", err)
		os.Exit(1)
	}

	// init logger
	app.Log = log.NewFileLogger(fmt.Sprintf("%s/nbx.log", app.GlobalDir))
	app.Log.SetLevel(log.Trace)
	if err := app.Log.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "error: Failed to open log: %v\n", err)
		os.Exit(1)
	}

	// init db

	// init config

	// init display

	// setup commands
	nbx := newNbxCommand()

	// run
	if err := nbx.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
