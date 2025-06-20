/*
   Copyright 2023 Docker Compose CLI authors

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

package tracing

import (
	"go.opentelemetry.io/otel"
)

// skipErrors is a no-op otel.ErrorHandler.
type skipErrors struct{}

// Handle does nothing, ignoring any errors passed to it.
func (skipErrors) Handle(_ error) {}

var _ otel.ErrorHandler = skipErrors{}

