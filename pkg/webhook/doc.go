/*
Copyright 2018 The Kubernetes Authors.

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

/*
Package webhook provides methods to build and bootstrap a webhook server.

Currently, it only supports admission webhooks. It will support CRD conversion webhooks in the near future.

Build webhooks

	webhook1 := &Webhook{
		Path: "/mutating-pods",
		Handler: mutatingHandler,
	}

	webhook2 := &Webhook{
		Path: "/validating-pods",
		Handler: validatingHandler,
	}

Create a webhook server.

	as, err := NewServer("baz-admission-server", mgr, ServerOptions{
		CertDir: "/tmp/cert",
	})
	if err != nil {
		// handle error
	}

Register the webhooks in the server.

	err = as.Register(webhook1, webhook2)
	if err != nil {
		// handle error
	}

Start the server by starting the manager

	err := mrg.Start(signals.SetupSignalHandler())
	if err != nil {
		// handle error
	}
*/
package webhook

import (
	logf "sigs.k8s.io/controller-runtime/pkg/internal/log"
)

var log = logf.RuntimeLog.WithName("webhook")
