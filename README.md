# GHelp 🚀

A powerful Go CLI tool that streamlines the creation of new Go projects by automatically setting up directory structure, initializing Git repositories, creating GitHub repositories and opens the new project in VsCode with a single command.

## Features ✨

- 📁 Creates organized project directory structure
- 🐙 Creates remote GitHub repository (optional)
- 🔧 Initializes local Git repository (if you choose to create remote repo)
- 🏠 Works relative to user's home directory
- ⚡ Fast and efficient setup process
- 🎯 Customizable project structure

## Installation

```bash
go install github.com/RajivTathireddy/GHelp@latest
```

## Prerequisites

- [Git](https://git-scm.com/) installed and configured
- [GitHub CLI](https://cli.github.com/) installed and authenticated with classic token created for current active account. You can check it with `gh auth token` command (required for `-r` flag)
- Go 1.19+ (for building from source)

## Usage

```bash
GHelp [flags]
```

### Basic Usage

```bash
# Create a new project with default settings
GHelp

# Create project at specific path
GHelp -p myawesome-project -n test_module

# Create project with GitHub repository
GHelp -p web-app -r my-web-app -d "My awesome web application"

# Create project with pkg structure instead of cmd
GHelp -p library-project -cmd=false
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-p` | `string` | `NewProject` | Creates directory at specified path relative to user's home directory. Will throw an error if directory already exists |
| `-cmd` | `bool` | `true` | When `true`, creates `cmd` subdirectory. When `false`, creates `pkg` subdirectory |
| `-r` | `string` | - | Name for the remote repository. Only creates GitHub repo if this flag is provided |
| `-d` | `string` | - | Description for the GitHub repository. Use with `-r` flag |
| `-n` | `string` | test_module | Use it to name the module if remote repo is not created.It will ignored if `-r` flag is provided.|

## Examples

### Create a simple project
```bash
GHelp -p calculator -n calculator
```
Creates: `~/calculator/` with `cmd/` subdirectory

### Create a library project
```bash
GHelp -p math-utils -cmd=false
```
Creates: `~/math-utils/` with `pkg/` subdirectory

### Create project with GitHub repository
```bash
GHelp -p awesome-api -r awesome-api -d "RESTful API for awesome features"
```
Creates: `~/awesome-api/` with local Git repo connected to GitHub

### Full example with all flags
```bash
GHelp -p web-scraper -cmd=true -r web-scraper-tool -d "A powerful web scraping tool built in Go"
```

## Project Structure

When you run GHelp, it creates the following structure:

```
~/your-project-name/
├── README.md
├── go.mod
├── .gitignore
├── .env
├── cmd/           # if -cmd=true (default)
│   └── main.go
└── pkg/           # if -cmd=false
    └── main.go
```

## What GHelp Does

1. **Directory Creation**: Creates a new directory at the specified path relative to your home directory
2. **Go Module Initialization**: Creates Go module 
3. **Git Repository Setup**: Initializes a local Git repository
4. **GitHub Integration**: Creates a remote repository on GitHub (if `-r` flag is provided)
5. **Remote Connection**: Links local repository to GitHub repo if you choose to create remote repo
6. **File Generation**: Creates essential files like `README.md`, `.gitignore`, and `.env`


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

If you encounter any issues or have questions, please:

- 🐛 [Report bugs](https://github.com/RajivTathireddy/GHelp/issues)
- 💬 [Start discussions](https://github.com/RajivTathireddy/GHelp/discussions)
- ⭐ Star this repository if you find it helpful!

---

<div align="center">
Personal Project by #Rajiv
</div>