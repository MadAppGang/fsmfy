# FSMFY
This is a [Go generator ](https://go.dev/blog/generate) tool to create a Golang type-safe code for [Xstate JS library](https://xstate.js.org/).

The implementation is using `"github.com/looplab/fsm"` library, which is not 100% compatible with XState. Nevetheless it works fine for simple sate machines.

## How to use



To make it works you need to install fmsfy tool:
```bash
go install github.com/MadAppGang/fsmfy
```

As usual go generator add the line in any go file:

```go
//go:generate fmsfy test.json MyFSM
```

The first argument `test.json` is FSM machine description from XState JS library. You can gran XState machine and run in console: `JSON.stringify(machine)`. 

The second argument is your FSM type name.

and run:
```bash
go generate ./...
```

All done.


## Limitations

Xstate supports events with multi-source and multi-destinations states. For example one event could be initiated from the list of states and lead to the other states.

In go library it is not supported, the event could be sourced in a list of states but lead only to one state. If you have incompatible event - the tool will fail with error.

