// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"log"
	"github.com/YuriyLisovskiy/lofp/src"
)

func main() {
	if err := src.RunCLI(); err != nil {
		log.Panic(err)
	}
}
