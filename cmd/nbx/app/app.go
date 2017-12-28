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

package app

// The app package is a global package for setting state and components that the
// entire application will use. This could be accomplished with a struct that
// is passed throughout the entire app, however in careful consideration, it
// would require that nearly every function within the app accept the global
// 'app' struct, which seems overkill.
//
// This package is designed such that all 'components' are public variables
// that implement an interface. This enables the entire app to be testable.
//
// When writing tests for packages that depend on this package, you can
// set the public variables to a mocked implementation of an interface.

import (
	"io"

	"github.com/nanobox-io/nbx/cmd/nbx/app/config"
	"github.com/nanobox-io/nbx/cmd/nbx/app/db"
	"github.com/nanobox-io/nbx/cmd/nbx/app/display"
	"github.com/nanobox-io/nbx/cmd/nbx/app/env"
	"github.com/nanobox-io/nbx/cmd/nbx/app/log"
)

// Stdin
var In io.Reader

// Stdout
var Out io.Writer

// Stderr
var Err io.Writer

// Env
var Env env.Envable

// Logger
var Log log.Logger

// DB
var DB db.Persistable

// Config
var Config config.Configable

// Display
var Display display.Displayable

// Global Directory
var GlobalDir string
