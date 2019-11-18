package factories

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

// HTTPServerFactory represents a factory of HTTP servers for testing.
type HTTPServerFactory struct {
}

// NewDummy creates a new TCP net.Listener at a random port, a new http.ServerMux responding
// "PAYLOAD" at GET /dummy and a new http.Server with the new mux. This function also creates
// a go routine that calls server.Serve(listener).
func (f *HTTPServerFactory) NewDummy() (net.Listener, *http.Server, error) {
	// create listener at a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, nil, err
	}

	// create server and mux
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
	}
	mux.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		const payload = "PAYLOAD"
		const bytesToBeWritten = len(payload)
		w.Header().Add("Content-Length", fmt.Sprintf("%d", bytesToBeWritten))
		bytesWritten, err := w.Write([]byte(payload))
		if err != nil {
			panic(fmt.Errorf("error writing dummy HTTP server response: %w", err))
		}
		if bytesWritten != bytesToBeWritten {
			panic(fmt.Errorf(
				"wrong number of bytes written by dummy HTTP server, want %d, got %d",
				bytesToBeWritten,
				bytesWritten,
			))
		}
	})

	go f.serveAndPanicOnError(server, listener, "dummy")

	return listener, server, nil
}

// NewDummyUploader creates a new TCP net.Listener at a random port, a new http.ServerMux responding
// upload routes and a new http.Server with the new mux. This function also creates
// a go routine that calls server.Serve(listener).
func (f *HTTPServerFactory) NewDummyUploader() (net.Listener, *http.Server, error) {
	// create listener at a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, nil, err
	}

	// create server and mux
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
	}
	mux.HandleFunc("/upload-info", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		port := listener.Addr().(*net.TCPAddr).Port
		uploadInfo := &struct {
			Method, URL string
			Headers     http.Header
		}{
			Method: http.MethodPut,
			URL:    fmt.Sprintf("http://localhost:%d/upload", port),
			Headers: http.Header{
				"Content-Length": []string{r.Header.Get("X-Content-Length")},
			},
		}
		payload, err := json.Marshal(uploadInfo)
		if err != nil {
			panic(fmt.Errorf("error marshaling dummy upload HTTP server response: %w", err))
		}
		bytesToBeWritten := len(payload)
		w.Header().Add("Content-Length", fmt.Sprintf("%d", bytesToBeWritten))
		bytesWritten, err := w.Write(payload)
		if err != nil {
			panic(fmt.Errorf("error writing dummy upload HTTP server response: %w", err))
		}
		if bytesWritten != bytesToBeWritten {
			panic(fmt.Errorf(
				"wrong number of bytes written by dummy upload HTTP server, want %d, got %d",
				bytesToBeWritten,
				bytesWritten,
			))
		}
	})
	var uploadedData []byte
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			var err error
			uploadedData, err = ioutil.ReadAll(r.Body)
			if err != nil {
				panic(fmt.Errorf("error reading upload: %w", err))
			}
			w.WriteHeader(http.StatusOK)
		} else if r.Method == http.MethodGet {
			bytesToBeWritten := len(uploadedData)
			w.Header().Add("Content-Length", fmt.Sprintf("%d", bytesToBeWritten))
			bytesWritten, err := w.Write(uploadedData)
			if err != nil {
				panic(fmt.Errorf("error writing upload: %w", err))
			}
			if bytesWritten != bytesToBeWritten {
				panic(fmt.Errorf(
					"wrong number of bytes written by dummy upload HTTP server, want %d, got %d",
					bytesToBeWritten,
					bytesWritten,
				))
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	go f.serveAndPanicOnError(server, listener, "dummy uploader")

	return listener, server, nil
}

// NewFileServer returns a server that serves files.
func (f *HTTPServerFactory) NewFileServer(relativePath string) (net.Listener, *http.Server, error) {
	// create listener at a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, nil, err
	}

	// create server and mux
	mux := &http.ServeMux{}
	mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		path := r.URL.Query()["path"][0]
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Errorf("error opening requested file: %w", err))
		}
		defer file.Close()
		w.WriteHeader(http.StatusOK)
		if _, err := io.Copy(w, file); err != nil {
			panic(fmt.Errorf("error copying requested file: %w", err))
		}
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			port := listener.Addr().(*net.TCPAddr).Port
			uploadInfo := &struct {
				Method, URL string
				Headers     http.Header
			}{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("http://localhost:%d/upload?path=%s", port, relativePath),
			}
			payload, err := json.Marshal(uploadInfo)
			if err != nil {
				panic(fmt.Errorf("error marshaling file server upload info response: %w", err))
			}
			bytesToBeWritten := len(payload)
			w.Header().Add("Content-Length", fmt.Sprintf("%d", bytesToBeWritten))
			bytesWritten, err := w.Write(payload)
			if err != nil {
				panic(fmt.Errorf("error writing file server upload info response: %w", err))
			}
			if bytesWritten != bytesToBeWritten {
				panic(fmt.Errorf(
					"wrong number of bytes written by file server, want %d, got %d",
					bytesToBeWritten,
					bytesWritten,
				))
			}
		case http.MethodPut:
			file, err := os.Create(r.URL.Query()["path"][0])
			if err != nil {
				panic(fmt.Errorf("error creating file for upload in file server: %w", err))
			}
			defer file.Close()
			if _, err := io.Copy(file, r.Body); err != nil {
				panic(fmt.Errorf("error copying uploaded file for file server: %w", err))
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	server := &http.Server{
		Handler: mux,
	}

	go f.serveAndPanicOnError(server, listener, "dummy")

	return listener, server, nil
}

func (f *HTTPServerFactory) serveAndPanicOnError(
	server *http.Server,
	listener net.Listener,
	serverName string,
) {
	if err := server.Serve(listener); err != http.ErrServerClosed {
		panic(fmt.Errorf("error on %s http.Server.Serve(): %w", serverName, err))
	}
}