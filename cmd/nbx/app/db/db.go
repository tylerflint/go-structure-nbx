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

package db

import (
  
)

type Persistable interface {
  Set(bucket, key string, val interface{}) error
  Get(bucket, key string, res interface{}) error
  GetAll(bucket, res interface{}) error
  Delete(bucket, key string) error
  Keys(bucket) (keys []string, err error)
  Truncate(bucket) error
}
