
## Running Server


```bash
go run main.go
```

## Generating the bindings
Use [abigen](https://geth.ethereum.org/docs/tools/abigen) to generate binding code from ABI

### Install abigen

``` bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

### Create .abi file and paste the ABI

### Generate Binding
```bash
abigen --abi <.abi file path> --pkg <go package name> --type Storage --out <output filepath>
```

^ Import this output file top call contract functions
