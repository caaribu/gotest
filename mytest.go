package mytest

import "runtime"

func GoVersion() string {
  return runtime.Version()
}
