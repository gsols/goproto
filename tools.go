//go:build tools
// +build tools

/*
 * go-template/go-template
 * tools.go Last modified at 2023/6/29 19:26
 * This file is part of go-template/go-template.
 * Copyright (c) 2023.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (including, but not limited to, loss of data, data being rendered inaccurate, or losses sustained
 * by you or third parties, or a failure of the software to operate with any other programs), even if the author has
 * been advised of the possibility of such damages.
 *
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions set
 * forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute, or otherwise
 * exploit this software.
 *
 * 1. The above copyright notice and this permission notice shall be included in all copies or substantial portions of
 * the Software.
 * 2. If this software is modified, the modified software must prominently display a notice indicating that it has been
 * modified from the original version.
 * 3. If this software is used to create a derived product, a copy of this notice must be included in all derived products.
 *
 */

package main

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"
	_ "github.com/cosmtrek/air"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "golang.org/x/tools/cmd/goimports"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
