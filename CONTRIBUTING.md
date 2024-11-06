# Contributing to GoCrab

Thank you for your interest in contributing to GoCrab! Contributions, issues, and feature requests are welcome. This document will guide you through the contribution process.

## How to Contribute

### 1. Fork the Repository

Fork the repository to your GitHub account and clone it locally. This will create a copy of the project where you can make changes.

```sh
git clone https://github.com/skysingh04/GoCrab.git
cd GoCrab
```

### 2. Create a Branch

Create a branch for your feature or fix. Make sure your branch name reflects the purpose of your contribution.

```sh
git checkout -b feature/amazing-feature
```

### 3. Make Changes

Make your changes in the appropriate files. Follow the project’s coding standards, and ensure your code is well-documented.

### 4. Test Your Changes

Run existing tests and add new ones if needed to verify that your changes work as expected. GoCrab uses Go’s testing framework for this purpose.

```sh
go test ./...
```

### 5. Commit Your Changes

Commit your changes with a clear and concise commit message.

```sh
git add .
git commit -m "Add amazing feature"
```

### 6. Push to GitHub

Push your branch to GitHub.

```sh
git push origin feature/amazing-feature
```

### 7. Create a Pull Request

Go to the original repository and create a Pull Request. Include a description of your changes, the reason behind them, and any additional context that may help the maintainers review your contribution.

## Code of Conduct

Please note that GoCrab follows a Code of Conduct to foster an open and welcoming environment. By participating, you are expected to uphold this standard.

## Reporting Issues

If you encounter any bugs or have suggestions for improvements, please open an issue on GitHub. Describe the issue in detail and include any relevant context to help reproduce it.

## Development Guidelines

- **Code Style**: Follow idiomatic Go conventions. Code should be clean, concise, and well-commented.
- **Documentation**: Document public functions, structs, and methods. Ensure that any code added is accompanied by relevant comments.
- **Tests**: Ensure that new features include tests. Aim for meaningful test coverage to maintain stability and prevent regressions.

## Need Help?

If you have any questions about the contribution process, please feel free to reach out by opening a discussion on GitHub or contacting the maintainers.

Thank you for contributing to GoCrab!