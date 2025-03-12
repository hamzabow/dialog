# Dialog CLI

A command-line interface that allows users to interact with system dialogs like file pickers, folder pickers, and message boxes.

This CLI is a wrapper around the [sqweek/dialog](https://github.com/sqweek/dialog) library, providing a convenient way to show native GUI dialogs directly from command-line scripts or terminal commands.

## Installation

```bash
go install github.com/hamzabow/dialog@latest
```

### System Dependencies

The dialog library uses native bindings for each platform:

- **Windows**: No additional dependencies required.
- **macOS**: No additional dependencies required.
- **Linux**: Requires GTK+3.0 (if not already installed on your system, you can install the required packages depending on your distribution)

  ```bash
  # Ubuntu/Debian
  sudo apt-get install libgtk-3-dev

  # Fedora
  sudo dnf install gtk3-devel

  # Arch Linux
  sudo pacman -S gtk3
  ```

## Usage

### File Selection Dialog

Open a file selection dialog and get the selected file path:

```bash
# Open a file dialog
dialog file

# With a custom title
dialog file --title "Select a configuration file"

# Start in a specific directory
dialog file --directory "/path/to/documents"

# Show a save dialog instead of open dialog
dialog file --save

# Filter files by type
dialog file --filter "Images (*.jpg *.png)" --filter "Documents (*.pdf *.doc)"
```

### Directory Selection Dialog

Open a directory selection dialog and get the selected directory path:

```bash
# Open a directory dialog
dialog directory

# Shorthand aliases
dialog dir
dialog folder

# With a custom title
dialog directory --title "Select output folder"

# Start in a specific directory
dialog directory --directory "/path/to/start"
```

### Message Box

Show a message box dialog:

```bash
# Show an info message
dialog message "Operation completed successfully"

# Show an error message
dialog message "Failed to connect to server" --type error

# With a custom title
dialog message "Hello World" --title "Greeting"
```

### Confirmation Dialog

Show a yes/no confirmation dialog:

```bash
# Ask for confirmation
dialog confirm "Are you sure you want to proceed?"

# With a custom title
dialog confirm "Delete this file?" --title "Warning"
```

The confirm command returns exit code 0 if the user clicked "Yes" and exit code 1 if the user clicked "No" or cancelled the dialog.

## Examples

### Use in shell scripts

```bash
#!/bin/bash

# Ask for confirmation
if dialog confirm "Do you want to continue?"; then
    echo "User chose YES"

    # Select a file
    CONFIG_FILE=$(dialog file --title "Select configuration file")
    echo "Selected file: $CONFIG_FILE"

    # Select an output directory
    OUTPUT_DIR=$(dialog directory --title "Select output directory")
    echo "Selected directory: $OUTPUT_DIR"

    # Show success message
    dialog message "Operation completed successfully"
else
    echo "User chose NO"
    dialog message "Operation cancelled" --type error
fi
```

## License

MIT