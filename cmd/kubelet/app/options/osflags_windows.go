// +build windows

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

package options

import (
	"github.com/spf13/pflag"
)

func (f *KubeletFlags) addOSFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&f.WindowsService, "windows-service", f.WindowsService, "Enable Windows Service Control Manager API integration")
	// The default priority class associated with any process in Windows is NORMAL_PRIORITY_CLASS. Keeping it as is
	// to maintain backwards compatibility.
	// Source: https://docs.microsoft.com/en-us/windows/win32/procthread/scheduling-priorities
	fs.StringVar(&f.WindowsPriorityClass, "windows-priorityclass", "NORMAL_PRIORITY_CLASS",
		"Set the PriorityClass associated with kubelet process, the default ones are available at "+
			"https://docs.microsoft.com/en-us/windows/win32/procthread/scheduling-priorities")
}
