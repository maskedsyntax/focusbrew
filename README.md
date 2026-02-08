# FocusBrew

A lightweight terminal-based Pomodoro Timer built with Go and Bubble Tea.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed (1.20+).

### Install via Go
```bash
go install github.com/MaskedSyntax/FocusBrew@latest
```
This will install the `FocusBrew` binary to your `$GOPATH/bin` directory.

### Build from source
```bash
git clone https://github.com/MaskedSyntax/FocusBrew.git
cd focusbrew
go build -o focusbrew
./focusbrew
```

## Usage

- `s`: Start the timer
- `p`: Pause the timer
- `r`: Reset current session
- `n`: Skip to next session
- `q`: Quit

The timer follows the standard Pomodoro technique:
- 25-minute work sessions
- 5-minute short breaks
- 15-minute long break after every 4 work sessions
