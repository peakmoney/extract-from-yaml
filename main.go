package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	yamlpatch "github.com/krishicks/yaml-patch"
	yaml "gopkg.in/yaml.v2"
)

var (
	rfc6901Decoder = strings.NewReplacer("~1", "/", "~0", "~")
)

// This section is extracted from the `yaml-patch` Container module
// because they are private, but extremely useful for what we're doing here.
// Used under license: https://github.com/krishicks/yaml-patch/blob/master/LICENSE
// Changes: Modified the arguments list to reference types within the yamlpatch
//          package instead of assuming we're still in it.
func decodePatchKey(k string) string {
	return rfc6901Decoder.Replace(k)
}

func findContainer(c yamlpatch.Container, path *yamlpatch.OpPath) (yamlpatch.Container, string, error) {
	parts, key, err := path.Decompose()
	if err != nil {
		return nil, "", err
	}

	foundContainer := c

	for _, part := range parts {
		node, err := foundContainer.Get(decodePatchKey(part))
		if err != nil {
			return nil, "", err
		}

		if node == nil {
			return nil, "", fmt.Errorf("path does not exist: %s", path)
		}

		foundContainer = node.Container()
	}

	return foundContainer, decodePatchKey(key), nil
}

// End

func handleError(silent bool, message string) {
	if silent == true {
		fmt.Print("")
		os.Exit(0)
	}
	log.Fatalf("Fatal error: %s", message)
}

func main() {
	silentPtr := flag.Bool("silent", false, "Silence any error and return empty strings instead")
	pathPtr := flag.String("path", "", "RFC6902-ish search path")
	flag.Parse()

	doc, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		handleError(*silentPtr, fmt.Sprintf("Unable to read stdin\n%s", err))
	}

	if len(doc) == 0 {
		handleError(*silentPtr, "no data was provided to stdin")
	}

	if pathPtr == nil {
		handleError(*silentPtr, "no search path specifed")
	}

	if len(*pathPtr) == 0 {
		handleError(*silentPtr, "cannot search for an empty search path")
	}

	var iface interface{}

	err = yaml.Unmarshal(doc, &iface)
	if err != nil {
		handleError(*silentPtr, fmt.Sprintf("Unable to parse yaml\n%s", err))
	}

	var c yamlpatch.Container
	c = yamlpatch.NewNode(&iface).Container()
	pathfinder := yamlpatch.NewPathFinder(c)
	paths := pathfinder.Find(*pathPtr)

	if len(paths) == 0 {
		handleError(*silentPtr, "Search path returned empty results")
	}

	newPath := paths[0]

	var opPath yamlpatch.OpPath
	opPath = yamlpatch.OpPath(newPath)
	con, key, fErr := findContainer(c, &opPath)
	if fErr != nil {
		handleError(*silentPtr, fmt.Sprintf("Unable to locate node with search path\n%s", fErr))
	}

	val, gErr := con.Get(key)

	if gErr != nil {
		handleError(*silentPtr, fmt.Sprintf("Unable to fetch value from key\n%s", err))
	}

	if val == nil {
		fmt.Print("")
		os.Exit(0)
	}

	result := val.Value()

	fmt.Print(result)
}
