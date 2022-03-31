// Copyright 2020-2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/onsi/ginkgo/v2"
)

func Kubectl(input io.Reader, args ...string) (string, error) {
	return cliRunner("kubectl", input, args...)
}

func Tanzu(input io.Reader, args ...string) (string, error) {
	return cliRunner("tanzu", input, args...)
}

func cliRunner(name string, input io.Reader, args ...string) (string, error) {
	ginkgo.GinkgoWriter.Printf("+ %s %s\n", name, strings.Join(args, " "))
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdin = input
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		rc := -1
		if ee, ok := err.(*exec.ExitError); ok {
			rc = ee.ExitCode()
		}

		ginkgo.GinkgoWriter.Println(stderr.String())
		return "", fmt.Errorf("%s\nexit status: %d", stderr.String(), rc)
	}

	ginkgo.GinkgoWriter.Println(stdout.String())
	return stdout.String(), nil
}
