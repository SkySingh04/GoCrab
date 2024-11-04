# CargoToGo

**CargoToGo** is a Rust-to-Go transpiler written in Go. Why? Its my attempt to learn about compilers/transpilers, who knows I may love it?

## Expected Features

- Translates core Rust syntax to Go equivalents
- Handles basic data structures, functions, and modules
- Converts Rust's memory safety features to idiomatic Go patterns
- Provides CLI for easy file and directory transpilation

## Roadmap

- [ ] Build the transpiler (to be elaborated)
- [ ] Write Proper tests
- [ ] Have CI/CD for automated releases
- [ ] Build a frontend to showcase the compiler
- [ ] Write Blog

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Installation

To install CargoToGo, run the following command:

```sh
go install github.com/skysingh04/cargotogo@latest
```

### Usage

After installing, you can start using **CargoToGo** to transpile Rust code files to Go:

```sh
cargotogo path/to/rust/file.rs
```

This command will generate a `.go` file with equivalent Go code in the same directory.



## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for more details on how to get involved.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
