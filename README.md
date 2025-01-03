# hyperctl

A simple sdk for Microsoft Hyper-V virtual machine management which is implemented in Go and using wmi.

## Installation

```bash
go get github.com/rokukoo/hyperctl
```

## Usage

### Start a virtual machine by name

```go
package main

import (
	"fmt"
	"github.com/rokukoo/hypervctl"
	"log"
)

func main() {
	vmName := "vm-name"
	vm, err := hypervctl.GetVirtualMachineByName(vmName)
	if err != nil {
		log.Panicln(err)
	}

	// Start the virtual machine
	if _, err = vm.Start(); err != nil {
        log.Panicln(err)
    }
}
```

### Stop a virtual machine by name

```go
package main

import (
    "fmt"
    "github.com/rokukoo/hypervctl"
    "log"
)

func main() {
    vmName := "vm-name"
    vm, err := hypervctl.GetVirtualMachineByName(vmName)
    if err != nil {
        log.Panicln(err)
    }

    // Stop the virtual machine 
    force := true
    if _, err = vm.Stop(force); err != nil {
        log.Panicln(err)
    }
}
```
