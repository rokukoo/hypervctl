# hyperctl

A simple sdk for Microsoft Hyper-V virtual machine management.

## Installation

```bash
go get github.com/rokukoo/hyperctl
```

## Usage

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