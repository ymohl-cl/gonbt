package gonbt_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ymohl-cl/gonbt"
)

const (
	exampleFilesPath = "./example"
)

func Example() {
	var err error
	var files []os.FileInfo

	if files, err = ioutil.ReadDir(exampleFilesPath); err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var dataIn []byte
		var tag gonbt.Tag

		if dataIn, err = ioutil.ReadFile(exampleFilesPath + "/" + file.Name()); err != nil {
			panic(err)
		}

		if tag, err = gonbt.Unmarshal(dataIn); err != nil {
			panic(err)
		}
		if _, err = gonbt.Marshal(tag, gonbt.CompressGZIP); err != nil {
			panic(err)
		}
		if _, err = gonbt.Marshal(tag, gonbt.CompressNone); err != nil {
			panic(err)
		}

		fmt.Printf("%s ok\n", file.Name())
	}
	// Unordered output:
	// servers.dat ok
	// level.dat ok
}
