package main

import (
	"github.com/chrismeyersfsu/strace/pkg/process"
)

func main() {
	process.CombineByTimestamp("/tmp/1.strace", "/tmp/2.strace", "/tmp/combine.strace")
}
