//go:build arm64
// +build arm64

package bolt

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 107374182400 // 100G

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0x7FFFFFFF

// Are unaligned load/stores broken on this arch?
var brokenUnaligned = false