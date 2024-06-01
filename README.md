# opmlmerge

Command line tool to merge opml files

## Installation

```bash
go install github.com/aquilax/opmlmerge@latest
```

## Usage

```bash
$ opmlmerge
Usage opmlmerge [FILE]...
```

Order of files matter as all subsequent files will be appended to the first one
excluding duplicate xmlURL-s
