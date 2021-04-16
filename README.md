# GMX

`gmx` is an extension for golang module management which can enable you to maintain a local collection of go modules in common use. It provides a set of interactive commands for an easy management of the local collection of golang modules. All the info of golang modules you collect will be stored in a file named `.gmx.yaml` in your home directory.

## Installation
Download and install the package with the following command:

```
$ go get github.com/utilbox/gmx
```

This will create the `gmx` executable under your `$GOPATH/bin` directory. You can move it to a permanent directory under `$PATH` so that you can use it without caring what the `$GOPATH` would be changed to.

You can also download the gost executable directly from the release page:

[https://github.com/utilbox/gmx/releases](https://github.com/utilbox/gmx/releases)

## Get started
All the commands provided by `gmx` are interactive. You can easily get to know about what to do next following the prompts. Basically, when you want to use the functions, you need to compose the commands in a format as below:

```
$ gmx [command]
```

All the available commands are listed as below:
- **add**: Add a module or a version to local collection.
- **help**: Help about any command.
- **list**: List all the modules in local collection.
- **rm**: Remove a module or a version from local collection.
- **search**: Search a module from the local collection.
- **update**: Update the info of a module in the local collection.
- **use**: Import a module to current project.

Use `gmx [command] --help` for more information about a command.