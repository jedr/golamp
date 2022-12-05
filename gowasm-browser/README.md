# Go WASM in browser

Go compiler can build WASM modules for the browser, but not for WASI.
This directory contains an example to build a browser WASM module using the official Go compiler.

For a more detailed description, search for "golang wasm" on the web,
here's an example: https://itnext.io/webassemply-with-golang-by-scratch-e05ec5230558.

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
cp $(go env GOROOT)/main/wasm/wasm_exec.js
python3 -m http.server 5000
```

Now open the browser at http://localhost:5000 and press the "Get time" button.

You should see the current time returned by the WASM module displayed in the page.
