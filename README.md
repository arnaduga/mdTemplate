# mdTemplate

This is a very basic program to render a template file, with variables from JSON file.

This is based on `text/Template` [Go package](https://pkg.go.dev/text/template): please follow this syntax

## Usage

In the `./dist` folder, the program has been complied into 6 binaries:

|         | Intel-based                    | ARM-based                      |
| ------- | ------------------------------ | ------------------------------ |
| Windows | `mdTemplate_windows_amd64.exe` | `mdTemplate_windows_arm64.exe` |
| Linux   | `mdTemplate_linux_amd64`       | `mdTemplate_linux_arm64`       |
| MacOS   | `mdTemplate_darwin_amd64`      | `mdTemplate_darwin_arm64`      |


Displaying the help:
```bash
$ mdTemplate_darwin_amd64 -help
Usage of mdTemplate_darwin_amd64:
  -json string
    	Path to the JSON file
  -output string
    	Path to the output file (default "output.md")
  -template string
    	Path to the template file
  -version
    	Show version and exit
```

Converting the example:
```bash
$ ./dist/mdTemplate_darwin_amd64 -json ./example/variables.json -template ./example/template.md
Template successfully executed and output written to output.md

$ cat output.md
# ORR for My wonderful Application

Author : [John Doe](mailto:johndoe@example.com)

## Section: Section 1

![img](inProgress.svg): This is a long criterai, section 1 number 1

- [Tests results](docs/test.pdf)

- [Pentests results](docs/test.pdf)

- [Load tests results](docs/test.pdf)

- [Chaos tests results](docs/test.pdf)


```