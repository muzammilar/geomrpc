/*
 * The package implements a basic go client for both geomserver and dataserver
 */

package client

/*
 * Client Package
 */
import (
	"sync"

	"google.golang.org/grpc"
)

/*
 * Public Functions
 */

// Create a Geometry Service client and perform the functions required
func GeometryClient(c *ServiceClient) {

	// notify the wait group when client finishes
	defer c.wg.Done()

	// create a connection
	opts := []grpc.DialOption{c.TLSOptions()}
	conn, err := grpc.Dial(c.addr, opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// setup a shared wait group
	var wg sync.WaitGroup

	// setup a service client

	wg.Add(1)
	go func() {

		wg.Done()
	}()

	// wait for internal wait groups
	wg.Wait()
}
