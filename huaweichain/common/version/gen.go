// Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.

// +build ignore

package main

// This program generate constants.go, which has the updated release version and platform version.

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

const (
	filename = "constants.go"
	perm     = 0644
)

func main() {
	releaseVersion := flag.String("releaseVersion", "Not provided", "release version")
	platformVersion := flag.String("platformVersion", "Not provided", "platform version")
	flag.Parse()

	var buf bytes.Buffer
	generate(&buf, *releaseVersion, *platformVersion)
	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, data, perm)
	if err != nil {
		log.Fatal(err)
	}
}

func generate(buf *bytes.Buffer, releaseVersion, platformVersion string) {
	println(buf, "/*")
	println(buf, " * Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.")
	println(buf, " */")
	println(buf)
	println(buf, "// Code generated by go run gen.go -releaseVersion "+
		"$RELEASE_VERSION -platformVersion $PLATFORM_VERSION; DO NOT EDIT.")
	println(buf)
	println(buf, "package version")
	println(buf)
	println(buf, "// These constants are updated by "+
		"'RELEASE_VERSION=xx.xx.xx.xx PLATFORM_VERSION=xx.xx go generate common/version'.")
	println(buf, "// It will execute 'go:generate go run gen.go' defined in release.go, which will")
	println(buf, "// update below constants based on input RELEASE_VERSION and PLATFORM_VERSION.")
	println(buf, "const (")
	println(buf, "// ReleaseVersion is the release version when wienerchain is released.")
	println(buf, "// It consists of 4 segments, i.e. 2.1.0.6")
	println(buf, fmt.Sprintf("ReleaseVersion = \"%s\"\n", releaseVersion))
	println(buf)
	println(buf, "// PlatformVersion is platform version when wienerchain is released.")
	println(buf, "// It consists of 2 segments, i.e. 1.0")
	println(buf, fmt.Sprintf("PlatformVersion = \"%s\"\n", platformVersion))
	println(buf)
	println(buf, "// GitSHA is the last commit id when wienerchain is released.")
	println(buf, "GitSHA = notProvided")
	println(buf, ")")
}

func println(buf *bytes.Buffer, a ...interface{}) {
	_, _ = fmt.Fprintln(buf, a...)
}
