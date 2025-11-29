# The Nature of Code on pure Go

While learning Go, I decided to translate examples and exercises from the [book](https://natureofcode.com) into plain Go.

I took inspiration from [philyawj/nature-of-code-ebiten]((https://github.com/philyawj/nature-of-code-ebiten/)) — for instance, I borrowed p5math from there.

Also referenced [stdiopt/gowasm-experiments](https://github.com/stdiopt/gowasm-experiments/)

The file [wasm_exec.js](./wasm_exec.js) is taken from tinygo/targets/wasm_exec.js.

## How to Run
In the repository root, I start a web server, for example:
```sh
npx http-server
```

Copy the `template` to create a new example or exercise:
```sh
cp -r template exercise...
```

Navigate into the created directory and rebuild the wasm file after changes:
```sh
cd exercise...
make build
```

Refresh the page in your browser: `127.0.0.1:8080/exercise...`

I intentionally kept the `.wasm` files in the repository so that anyone who wants can simply clone and run the project without additional builds.