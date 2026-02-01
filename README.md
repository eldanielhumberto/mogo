# Mogo

A CLI application for easily managing monorepos.

## Installation

Install the package using Go:

```bash
go install github.com/eldanielhumberto/mogo@latest
```

After installation, make sure to add the Go bin directory to your PATH:

```bash
export PATH="$PATH:$HOME/go/bin"
```

## Usage

Mogo provides a set of commands to help you manage your monorepo projects efficiently.

1. Run the following command in your monorepo project:

```bash
mogo init
```

2. Run the following command to add the packages or workspaces

```bash
mogo add <dir_package>

# Example: mogo add ./apps/frontend
```

3. After adding packages or workspaces, you can configure commands in `mogo.json`. An example is shown below:

```json
{
  "workspaces": {
    "frontend": {
      "context": "./apps/frontend",
      "commands": {
        "dev": "npm run dev",     // Command 'dev' in frontend
        "build": "npm run build"
      }
    },
    "backend": {
      "context": "./apps/backend",
      "commands": {
        "dev": "go run main.go",  // Command 'dev' in backend
        "build": "go build -o bin/backend main.go"
      }
    }
  }
}

```

4. Run the following command to execute commands in the workspaces:

```bash
mogo run <command> --workspace=<workspace>

# Example: mogo run dev --workspace=frontend
```

5. To run two commands in parallel, the two commands configured in `mogo.json` must have the same name, so that the following command can be executed:

```bash
mogo run <command>

# Example: mogo run dev
# This will run the 'dev' command in both frontend and backend workspaces in parallel.
```

Run `mogo --help` to see all available commands and options.
