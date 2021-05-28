package main

import (
	"net/http"
)

func main() {
	// dirname, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Current directory: %v\n", dirname)

	// dir, err := os.Open(filepath.Join(dirname, "../../staticweb/public/sb-admin"))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Name of ../: %v\n", dir.Name())

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../../staticweb/public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":9000", mux)
}
