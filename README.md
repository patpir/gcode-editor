# gcode-editor

Basic gcode analysis and editing CLI. 


## Usage

Install to gobin with:

```shell
go get github.com/patpir/gcode-editor
```

Cannot find where the binary was installed?
Try `go list -f '{{.Target}}' github.com/patpir/gcode-editor`, it will show the
full path to the binary.


## Development

Development requires Go (it should work with any version that supports modules).

To compile project and run a command (e.g. the visualize command) go to the root
directory of the project and execute:

```shell
go run main.go visualize example.gcode
```

The project currently has no tests, which hopefully will change soon!


## Contributing

**Contributions are welcome!**

Feel free to create pull requests, for example:
- to add tests,
- to improve the visualization (multiple graphs with different scaling, without
  any axis labels - yikes!),
- to add support for different gcode firmwares,
- or to add whatever functionality you see need for!

