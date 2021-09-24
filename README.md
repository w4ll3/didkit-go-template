# didkit-go-template
Simple usage demonstration of the `didkit-go` package.

Whenever cloning this repo for the first time make sure to have [Rustup](https://rustup.rs/) and run:
```bash
# update submodules to head
git submodule update --init --recursive
# builds didkit
make -C extern/didkit/lib ../target/test/c.stamp
```