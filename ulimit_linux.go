package main

import "syscall"

func init() {
	var rLimit syscall.Rlimit
	rLimit.Cur = 1024 * 1024
	rLimit.Max = 1024 * 1024
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
