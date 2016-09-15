package daak

import (
	"bufio"
	"io"
	"strings"
)

const (
	con  = "HAI"
	pub  = "SAY"
	ping = "UTHR"
)


func Parse(rc io.ReadCloser, r Routes) (err error) {
	nr := bufio.NewReader(rc)

	for {
		// Start with the stupidest way to do this.
		data, err := nr.ReadString('\n')
		if err != nil {
			return err
		}
		arr := strings.SplitN(strings.Trim(data,"\n"), " ", 2)

		switch arr[0] {
		case con:
			err =r.Connect([]byte(arr[1]))
		case ping:
			err = r.Ping()
		case pub:
			lol := strings.SplitN(arr[1], " ", 2)
			err = r.Pub([]byte(lol[0]), []byte(lol[1]))
		default:
			err = r.Unknown()
		}
		if err != nil {
			return err
		}
	}

	return nil
}
