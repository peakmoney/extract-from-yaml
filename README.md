# extract-from-yaml

A very simple command line utility to extract a specific value from a YAML file.

## Usage

Given

```yaml
# my_file.yaml
my_array:
- name: Fred
  age: 25
- name: Sally
  age: 27
```

Command

`cat my_file.yaml | extract-from-yaml -path /my_array/name=Fred/age`

Result

`25`

`extract-from-yaml` uses Stdin to receive your YAML file. If you do not pipe it in on the command line, the program will wait for an EOF condition before proceeding. This can be triggered by Ctrl-D on most systems. Feel free to just type or paste YAML directly into your console to play around.

## Why?

There are other tools out there. We even use some of them.

`yq`: yq is a great general purpose YAML utility. While it supports reading from YAML file and extracting a specific value, it doesn't support the fuzzy-search system that this command does.

`yaml-patch`: We use yaml-patch to easily modify our YAML manifests. However, it doesn't provide functionality to simply read existing YAML files using the power of its pathfinder engine and RFC6902 paths.

`extract-from-yaml` (better name pending) uses the Pathfinder engine from `yaml-patch` to extract values from any location in a YAML file - even in arrays. When working with arrays, you can use a query interface to select a specific instance of structured data instead of relying on a hard-coded array index like other tools.

## Is it any good?

No. It works when it works. If there is an error (like if something doesn't exist), it's not good right now. I'm working on it.

# License

MIT
