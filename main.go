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

func main() {
	pathPtr := flag.String("path", "", "RFC6902-ish path to search for")
	flag.Parse()

	doc, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading from stdin: %s", err)
		return
	}

	if len(doc) == 0 {
		log.Fatalf("error: no data provided to stdin")
		return
	}

	var iface interface{}

	err = yaml.Unmarshal(doc, &iface)
	if err != nil {
		log.Fatalln("can't unmarshall")
		return
	}

	var c yamlpatch.Container
	c = yamlpatch.NewNode(&iface).Container()

	pathfinder := yamlpatch.NewPathFinder(c)
	paths := pathfinder.Find(*pathPtr)

	newPath := paths[0]

	var opPath yamlpatch.OpPath
	opPath = yamlpatch.OpPath(newPath)
	con, key, nerr := findContainer(c, &opPath)
	if nerr != nil {
		log.Fatalln("Whoops... no thing...")
	}

	val, err := con.Get(key)
	result := val.Value()

	fmt.Printf("%s\n", result)
}
