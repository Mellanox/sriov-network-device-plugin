// Copyright 2026 NVIDIA CORPORATION & AFFILIATES
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
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

const (
	sriovDPBinary   = "/usr/bin/sriovdp"
	defaultLogLevel = 10
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"This is an entrypoint for SR-IOV Network Device Plugin\n\n"+
			"./entrypoint\n"+
			"\t-h --help\n"+
			"\t--log-dir=\n"+
			"\t--log-level=%d\n"+
			"\t--resource-prefix=\n"+
			"\t--config-file=\n"+
			"\t--use-cdi\n",
		defaultLogLevel)
}

func run() int {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	logDir         := fs.String("log-dir", "", "Log directory under /var/log/")
	logLevel       := fs.Int("log-level", defaultLogLevel, "Log verbosity level")
	resourcePrefix := fs.String("resource-prefix", "", "Resource prefix for devices")
	configFile     := fs.String("config-file", "", "Path to config file")
	useCDI         := fs.Bool("use-cdi", false, "Enable CDI")
	fs.Usage = usage

	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to parse flags: %v\n", err)
		return 1
	}

	args := []string{sriovDPBinary, fmt.Sprintf("-v=%d", *logLevel)}

	if *logDir != "" {
		logPath := filepath.Join("/var/log", *logDir)
		if err := os.MkdirAll(logPath, 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "failed to create log dir %q: %v\n", logPath, err)
			return 1
		}
		args = append(args, "--log_dir", logPath, "--alsologtostderr")
	} else {
		args = append(args, "--logtostderr")
	}

	if *resourcePrefix != "" {
		args = append(args, "--resource-prefix", *resourcePrefix)
	}

	if *configFile != "" {
		args = append(args, "--config-file", *configFile)
	}

	if *useCDI {
		args = append(args, "--use-cdi")
	}

	if err := syscall.Exec(sriovDPBinary, args, os.Environ()); err != nil {
		fmt.Fprintf(os.Stderr, "failed to exec %q: %v\n", sriovDPBinary, err)
		return 1
	}

	// unreachable after successful exec
	return 0
}

func main() {
	os.Exit(run())
}
