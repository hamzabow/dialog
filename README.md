# Dialog CLI

[![Build Status](https://github.com/hamzabow/dialog/actions/workflows/build.yaml/badge.svg)](https://github.com/hamzabow/dialog/actions/workflows/build.yaml)
[![Release](https://img.shields.io/github/v/release/hamzabow/dialog)](https://github.com/hamzabow/dialog/releases/latest)

A command-line interface that allows users to interact with system dialogs like file pickers, folder pickers, and message boxes.

This CLI is a wrapper around the [sqweek/dialog](https://github.com/sqweek/dialog) library, providing a convenient way to show native GUI dialogs directly from command-line scripts or terminal commands.

## Platform Support

**Currently, only Windows 11 is supported** (this is the only platform where it has been tested).

While the underlying library states that it supports Linux and macOS, this CLI tool hasn't been tested on those platforms yet.

## Installation

### Prerequisites
- Go 1.21 or higher ([Installation Guide](https://go.dev/doc/install))

### Install
```powershell
go install github.com/hamzabow/dialog@latest
```

## Usage

### File Selection Dialog

Open a file selection dialog and get the selected file path:

```powershell
# Open a file dialog
dialog file

# With a custom title
dialog file --title "Select a configuration file"

# Start in a specific directory
dialog file --directory "C:\Users\Public\Documents"

# Show a save dialog instead of open dialog
dialog file --save

# Filter files by type
dialog file --filter "Images (*.jpg *.png)" --filter "Documents (*.pdf *.doc)"
```

### Directory Selection Dialog

Open a directory selection dialog and get the selected directory path:

```powershell
# Open a directory dialog
dialog directory

# Shorthand aliases
dialog dir
dialog folder

# With a custom title
dialog directory --title "Select output folder"

# Start in a specific directory
dialog directory --directory "C:\Users\Public\Documents"
```

### Message Box

Show a message box dialog:

```powershell
# Show an info message
dialog message "Operation completed successfully"

# Show an error message
dialog message "Failed to connect to server" --type error

# With a custom title
dialog message "Hello World" --title "Greeting"
```

### Confirmation Dialog

Show a yes/no confirmation dialog:

```powershell
# Ask for confirmation
dialog confirm "Are you sure you want to proceed?"

# With a custom title
dialog confirm "Delete this file?" --title "Warning"
```

The confirm command returns exit code 0 if the user clicked "Yes" and exit code 1 if the user clicked "No" or cancelled the dialog.

## Examples

### Use in PowerShell scripts

```powershell
# Ask for confirmation
dialog confirm "Do you want to continue?"
if ($LASTEXITCODE -eq 0) {
    Write-Host "User chose YES"

    # Select a file
    $CONFIG_FILE = dialog file --title "Select configuration file"
    Write-Host "Selected file: $CONFIG_FILE"

    # Select an output directory
    $OUTPUT_DIR = dialog directory --title "Select output directory"
    Write-Host "Selected directory: $OUTPUT_DIR"

    # Show success message
    dialog message "Operation completed successfully"
} else {
    Write-Host "User chose NO"
    dialog message "Operation cancelled" --type error
}
```

### Use in WSL (Windows Subsystem for Linux)

Since WSL can access Windows executables that are in the PATH, you can use this tool in bash scripts when running under WSL. You'll need to call it with the `.exe` extension:

```bash
#!/bin/bash

# Ask for confirmation
if dialog.exe confirm "Do you want to continue?"; then
    echo "User chose YES"

    # Select a file
    CONFIG_FILE=$(dialog.exe file --title "Select configuration file")
    echo "Selected file: $CONFIG_FILE"

    # Select an output directory
    OUTPUT_DIR=$(dialog.exe directory --title "Select output directory")
    echo "Selected directory: $OUTPUT_DIR"

    # Show success message
    dialog.exe message "Operation completed successfully"
else
    echo "User chose NO"
    dialog.exe message "Operation cancelled" --type error
fi
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.