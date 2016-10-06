Another CLI interface for https://github.com/ghodss/yaml

### Installation

    go get github.com/fgrosse/json2yaml

### Usage

    Usage: json2yaml <FILENAME>
    
    Convert the json contents of FILENAME to yaml and print it on stdout. If
    FILENAME is '-' or omitted, stdin is read instead.
    The input is read completely before re-encoding begins.

### Related

See https://github.com/pschultz/yaml2json if you are searching for the inverse transformation.
