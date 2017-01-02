# csvq make plain text acceptable by photoshop data set import
This tool converts strings into format that is acceptable by apply csv quoting and some other ticks.

Photoshop can import strings from a file, but often it shows error message like

> Could not parse the file contents as a data set. Too many values were listed for data set 2 (or any integer).

The problem is photo shop import file as csv format, but plain text needs csv quoting and some other tricks.

## Install
- install golang https://golang.org/doc/install
- `go get github.com/helinwang/csvq`

## Usage
- `cat foo.txt | csvq > foo_new.txt`
- And then import `foo_new.txt` from photoshop data set import.
