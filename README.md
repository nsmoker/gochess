# gochess

This is the chess rules library I've been developing alongside my chess GUI, CRD. This app is specifically targeted at Chess analysis GUIs, and probably isn't the best choice for an engine or similar. 

The main file and proto files are the API I use with CRD, which is based on Unix streams. You could easily use this as a library for something written in Go, though, or with TCP if you want. 

To run the Unix Stream API, you can just run `go run main.go`.
