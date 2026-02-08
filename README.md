# FocusBrew

A lightweight terminal-based Pomodoro Timer built with Go and Bubble Tea.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed (1.20+).

### Install via Go
```bash
go install github.com/maskedsyntax/focusbrew@latest
```
This will install the `focusbrew` binary to your `$GOPATH/bin` directory. 

**Note:** Make sure `$(go env GOPATH)/bin` is in your `PATH`. You can add it by adding this to your `.zshrc` or `.bashrc`:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Build from source
```bash
git clone https://github.com/maskedsyntax/focusbrew.git
cd focusbrew
go build -o focusbrew
./focusbrew
```

## Usage

Run `focusbrew` to start with default settings (25m work, 5m short break, 15m long break).

### Custom Durations
You can customize the timer using flags:
```bash
focusbrew -work 50 -short 10 -long 30
```

### Controls
- `s`: Start the timer
- `p`: Pause the timer
- `r`: Reset current session
- `n`: Skip to next session
- `q`: Quit

The timer follows the standard Pomodoro technique:
- 25-minute work sessions (default)
- 5-minute short breaks (default)
- 15-minute long break after every 4 work sessions (default)
