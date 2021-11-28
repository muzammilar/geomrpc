/*
 * The package implements a basic examle of an http server
 */

package geomgenerator

import (
	"math/rand"
	"time"
)

/*
 * Init Function
 */

func init() {
	// initialize global pseudo random generator
	// Random numbers are generated by a Source.
	// Top-level functions, such as Float64 and Int, use a default shared Source that produces a deterministic sequence
	// of values each time a program is run.
	// Use the Seed function to initialize the default Source if different behavior is required for each run.
	// The default Source is safe for concurrent use by multiple goroutines, but Sources created by NewSource are not.
	rand.Seed(time.Now().Unix())

}