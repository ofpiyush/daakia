package listeners
//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"net"
//	"net/http"
//	"strings"
//	"sync/atomic"
//)
//
//type HttpListener struct {
//	mx *http.ServeMux
//}
//
//func NewHTTP() *HttpListener {
//	return &HttpListener{
//		mx: http.NewServeMux(),
//	}
//}
//
//// Muhaha, this missing p will haunt you for years
//type httPipe struct {
//	dest    string
//	r       io.ReadCloser
//	running int32
//}
//
//func (hp *httPipe) Read(p []byte) (n int, err error) {
//	return hp.r.Read(p)
//}
//
//func (hp *httPipe) Write(p []byte) (n int, err error) {
//	n = len(p)
//	_, err = http.Post(hp.dest, "text/plain", bytes.NewReader(p))
//	return
//}
//
//func (hp *httPipe) Close() error {
//	atomic.SwapInt32(&hp.running, 0)
//	return hp.r.Close()
//}
//
//func (h *HttpListener) Handle(path string, ha func(net.Conn) error) {
//	arr := strings.Split(path, "->")
//	path = arr[0]
//	pr, pw := io.Pipe()
//	hP := &httPipe{
//		dest: arr[1],
//		r:    pr,
//	}
//	atomic.SwapInt32(&hP.running, 1)
//	h.mx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
//		go ha(hP)
//		io.Copy(pw, r.Body)
//		fmt.Fprintf(w, "ok")
//	})
//}
//
//func (h *HttpListener) Listen(addr string) error {
//	return http.ListenAndServe(addr, h.mx)
//}
