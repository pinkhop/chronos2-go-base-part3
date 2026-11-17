# chronos2-go-base-part3

Shards 7 through 9 of the 12 that make up Amazon's [Chronos-2](https://huggingface.co/amazon/chronos-2) checkpoint, for [`chronos2-go-base`](https://github.com/pinkhop/chronos2-go-base).

Do not import this module. Import `github.com/pinkhop/chronos2-go-base`, which requires all four parts and assembles them into a ready-to-use model; `go get` fetches the parts automatically. The checkpoint is split across four modules so that each stays small enough for the Go module proxy to serve.


## Contents

| File | Bytes | SHA-256 |
|---|---|---|
| `model-00007-of-00012.safetensors` | 37,759,488 | `ece4d17ba53d78eaf75208683c35338e7faec77b04390ad249744a948c563529` |
| `model-00008-of-00012.safetensors` | 37,759,488 | `5a3d8c10d7717d2767b2f82fac705b18283ca744948328fab65fd7cf35e23e5b` |
| `model-00009-of-00012.safetensors` | 37,759,488 | `f0f86171eeb00cf5a92e1fbf1a76d3952752b7140fdee7edbbcafbbe9b6ad4f3` |

The shards are fp32 safetensors files cut from the upstream `model.safetensors` at revision [`29ec3766d36d6f73f0696f85560a422f50e8498c`](https://huggingface.co/amazon/chronos-2/tree/29ec3766d36d6f73f0696f85560a422f50e8498c), with every tensor bit-identical to upstream. [`weights/manifest.json`](weights/manifest.json) records the upstream source and each shard's size and SHA-256, and `make verify-weights` checks the files against it.


## Contributing

[`AGENTS.md`](AGENTS.md) describes the repository, the `make` targets, and the rules every change follows, for human and AI contributors alike.


## License

Licensed under the Apache License, Version 2.0; see [`LICENSE`](LICENSE).

Chronos-2 was created by Amazon, and its weights are distributed under Apache-2.0; see [`NOTICE`](NOTICE).
