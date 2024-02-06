# ixpt-syscheck

A check for commands the caller expects to already exist on the system.

## Usage

Call from another module:

```golang

import "github.com/spqrix/ixpt-syscheck/executables"

// NOTE:
//        'commands' is a list of commands as strings (no paths needed).
//        'verb' is a boolean for verbose-mode.
missing_execs, err := executables.CheckForExecutables(commands, verb)
if err != nil {
  fmt.Printf("Error during system check: %q", err)
  os.Exit(1)
}
for _, c := range missing_execs {
  fmt.Printf("Failed to locate file path for: %s\n", c)
}

```

Run the optional cli:

`./ixpt-syscheck curl wget asdf ping`

Run the optional cli in verbose mode (show full paths for existing commands):

`./ixpt-syscheck curl wget asdf ping -v`
