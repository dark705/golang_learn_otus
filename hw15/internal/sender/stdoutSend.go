package sender

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

func SendToStdout(msg []byte, i int) error {
	rand.Seed(time.Now().UTC().UnixNano() + int64(i))
	rnd := rand.Intn(2000) + 1000
	time.Sleep(time.Millisecond * time.Duration(rnd)) //emulate delay on sender
	if rnd < 2000 {                                   //emulate error on send
		fmt.Println("Sender:", i, "SendMessage:", string(msg))
		return nil
	}
	return errors.New("Error on sender")
}
