package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	filepath.Walk("../../../../../../../../", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(" [ ", info.IsDir(), " ] = > ", path)
		return nil
	})
	// arg := os.Args[1]
	// if len(arg) < 1 {
	// 	fmt.Fprintf(os.Stderr, "No Enough Arguments")
	// 	os.Exit(-1)
	// }
	// if file, err := os.Open(arg); err != nil {
	// 	fmt.Errorf("%s is bad file ", arg)
	// 	os.Exit(-1)
	// } else {
	// 	defer file.Close()
	// 	h := md5.New()
	// 	io.Copy(h, file)
	// 	fmt.Printf("%x  In hash\n", h.Sum(nil))
	// }

}
