/*
 * The package implements a basic go client for both geomserver and dataserver
 */

package client

/*
 * Client Package
 */
import (
	"context"
	"io"
	"sync"

	"github.com/muzammilar/geometric-shapes/protos/shape"
	"github.com/muzammilar/geometric-shapes/protos/shapestore"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

/*
 * Public Functions
 */

// Create a Generator Service client and perform the functions required
func GeneratorClient(c *ServiceClient) {

	// notify the wait group when client finishes
	defer c.wg.Done()

	// create a connection
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(c.creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// setup a shared wait group
	var wg sync.WaitGroup

	// setup a service client
	storeClient := shapestore.NewStoreClient(conn)
	generatorClient := shapestore.NewGeneratorClient(conn)

	c.logger.Infof("Store service client: %#v", storeClient)
	c.logger.Infof("Generator service client: %#v", generatorClient)

	go func() {
		wg.Add(1)
		generatorCuboid(generatorClient, c)
		wg.Done()
	}()

	// wait for internal wait groups
	wg.Wait()
}

func generatorCuboid(g shapestore.GeneratorClient, c *ServiceClient) {
	// Based on this, context is generally unique (so we can use it to identify information):
	// https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md
	ctxGen, cancelGen := context.WithCancel(context.Background())
	defer cancelGen()
	serverStream, err := g.Cuboid(ctxGen, &emptypb.Empty{})
	if err != nil {
		c.logger.Errorf("Error setting up server stream for Generator service's Cuboid function: %s", err.Error())
		return
	}

	recvCount := 0
	var cuboid *shape.Cuboid
	for cuboid, err = serverStream.Recv(); err == nil; cuboid, err = serverStream.Recv() {
		// ignore empty values for now
		if cuboid == nil {
			continue
		}
		// log data
		c.logger.Info(cuboid)
		// keep recv count
		recvCount++
		if recvCount >= numMessagesMax {
			break

		}
	}
	// EOF - stream closed
	if err == io.EOF || err == nil {
		return
	}
}
