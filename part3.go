// Package part3 embeds shards 7 through 9 of the 12 that make up the
// amazon/chronos-2 checkpoint, for package
// [github.com/pinkhop/chronos2-go-base]. Import that package instead: it
// assembles all four parts into a ready-to-use model.
//
// The checkpoint is split across four modules so that each stays small
// enough for the Go module proxy to serve.
package part3

import _ "embed"

// The shards are embedded as byte slices rather than strings because the Go
// linker currently aligns embedded byte-slice data but not string data (not
// a language guarantee; the alias test in the chronos2-go repository's parity
// module verifies it): chronos2.Load can then view the weights in place
// instead of copying them onto the heap. Nothing may write to these slices;
// the model every goroutine shares reads them for the life of the process.
var (
	// shard07 is weights/model-00007-of-00012.safetensors.
	//
	//go:embed weights/model-00007-of-00012.safetensors
	shard07 []byte

	// shard08 is weights/model-00008-of-00012.safetensors.
	//
	//go:embed weights/model-00008-of-00012.safetensors
	shard08 []byte

	// shard09 is weights/model-00009-of-00012.safetensors.
	//
	//go:embed weights/model-00009-of-00012.safetensors
	shard09 []byte
)

// Shards returns this part's safetensors shards in checkpoint order, for
// package base to pass to chronos2.Load. The shards alias the embedded data,
// so callers must never modify them. Each call returns a new outer slice,
// so appending to it never affects another caller.
func Shards() [][]byte {
	return [][]byte{shard07, shard08, shard09}
}
