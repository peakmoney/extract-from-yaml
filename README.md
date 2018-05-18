# extract-from-yaml

A very simple command line utility to extract a specific value from a yaml file.

## Usage

Soon.

## Why?

There are other tools out there. We even use some of them.

`yq`: yq is a great general purpose Yaml utility. While it supports reading from Yaml file and extracting a specific value, it doesn't support the fuzzy-search system that this command does.
`yaml-patch`: We use yaml patch to easily modify our Yaml manifests. However, it doesn't provide functionality to simply read existing Yaml files using the power of its pathfinder engine and RFC6902 paths.

extract-from-yaml (better name pending) uses the Pathfinder engine from `yaml-patch` to extract values from any location in a yaml file - even in arrays. When working with arrays, you can use a query interface to select a specific instance of structured data instead of relying on a hardcoded array index like other tools.

## Is it any good?

No. It works when it works. If there is an error (like if something doens't exist), it's not good right now. I'm working on it.

# License

MIT
