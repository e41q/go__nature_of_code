# The Nature of Code on pure Go

While learning Go, I decided to translate examples and exercises from the [book](https://natureofcode.com) into plain Go.

I took inspiration from [philyawj/nature-of-code-ebiten]((https://github.com/philyawj/nature-of-code-ebiten/)) — for instance, I borrowed p5math from there.

Also referenced [stdiopt/gowasm-experiments](https://github.com/stdiopt/gowasm-experiments/)

The file [wasm_exec.js](./wasm_exec.js) is taken from tinygo/targets/wasm_exec.js.

## Complete exercise and examples

* [Example 0.1: Random walk](./example_0_1_random_walk/)
* [Example 0.2: Random number distribution](./example_0_2_random_number_distribution/)
* [Example 0.3: A Walker That Tends to Move to the Right](./example_0_3_random_walk_that_tends/)
* [Exercise 0.3: Random walk that tends to mouse](./exercise_0_3_random_walk_to_mouse/)
* [Example 0.4: A Gaussian Distribution](./example_0_4_gaussian_distribution/)
* [Exercise 0.4: Gaussian splatter](./exercise_0_4_gaussian_painter/)
* [Exercise 0.5: Gaussian walker](./exercise_0_5_gaussian_walker/)
* [Example 0.5: An Accept-Reject Distribution](./example_0_5_accept_reject_distribution/)
  * [Example 0.0: gaussian distribution](./example_0_0_gaussian_distribution/)
* [Exercise 0.6: Accept-reject walker](./exercise_0_6_accept_reject_walker/)
* [figure 0.4 right: Random Noise](./figure_0_4_right_random_noise/)
* [figure 0.4 left: Perlin Noise](./figure_0_4_left_perlin_noise/)

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