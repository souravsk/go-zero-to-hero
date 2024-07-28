package acceptancetest

import (
	"testing"
	"time"

	"github.com/souravsk/go_with_tests/acceptance_test"
	"github.com/souravsk/go_with_tests/assert"
)

const (
	port = "8080"
	url  = "<http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptance_test.LunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	//just check the server works before we shut things down
	assert.CanGet(t, url)

	//time.AfterFunc is a function from the time package that waits for the specified duration (in this case, 50 milliseconds) and then executes the provided function.
	//fire off a request, and before it has a chance to respond send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	//without graceful shutdown, this would fail
	assert.CanGet(t, url)

	//after interrput, the server should be shutdown, and no more requests will work
	assert.CantGet(t, url)
}
