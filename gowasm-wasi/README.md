# Go WASM on WASI

Let's build a WASI module using TinyGo.

Here's one of the articles on the web about this: https://www.fermyon.com/wasm-languages/go-lang.

This article has a bit more information: https://elewis.dev/are-we-wasm-yet-part-2.

Prerequisites:

- Install TinyGo https://tinygo.org/getting-started/install/
- Install a WASM runtime, e.g. [wasmer](https://docs.wasmer.io/ecosystem/wasmer/getting-started) or [wasmtime](https://github.com/bytecodealliance/wasmtime#installation)

Compile the code to a WASM module:

```sh
tinygo build -wasm-abi=generic -target=wasi -gc=leaking -o main.wasi ./main.go
```

Run the module with a WASM runtime:

```sh
$ wasmer run ./main.wasi
Current time: '2022-12-05 11:42:12.561299287 +0000 UTC'
$ wasmtime run ./main.wasi
Current time: '2022-12-05 11:42:29.266650976 +0000 UTC'
```

## Run with Docker

Read the docs here: https://docs.docker.com/desktop/wasm/.

Prerequisites:

- Docker Desktop v4.15 or higher

After building the WASM module `main.wasi` with TinyGo following the steps above,
build and run the Docker image:

```sh
docker buildx build --load --platform=wasi/wasm32 -t jedr/gowasm-wasi .
docker run --rm --runtime=io.containerd.wasmedge.v1 --platform=wasi/wasm32 jedr/gowasm-wasi
```
