# t2

T2 is a CLI tool for creating nested folders/files much like the `mkdir -p` command. Written in Golang, it is cross-platform and can be used on Windows, Linux, and Mac.

> Note: This is a work in progress. I am still working on the CLI and the documentation.
**essentiallly, this is a glorified mkdir -p and touch**
## Installation
```bash
go get github.com/karim-w/t2@latest
```

## Usage
### Create Nested Folders
```bash
t2 foo/bar/baz
```
This will create the following folder structure:
```
foo
└── bar
    └── baz
```
### Create a File Within Nested Folders
```bash
t2 foo/bar/baz/file.txt
```
This will create the following folder structure:
```
foo
└── bar
    └── baz
        └── file.txt
```


