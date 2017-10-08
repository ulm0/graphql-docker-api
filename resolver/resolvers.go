package resolver

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var (
	ctx = context.Background()
	cli = clientInit()
)

/* Docker Client */

func clientInit() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

/* System type related objects */

func (r *Resolver) System() *systemResolver {
	return &systemResolver{system: *cli}
}

/* Swarm type related objects */

func (r *Resolver) Swarm() *swarmResolver {
	sw, _ := cli.SwarmInspect(ctx)
	return &swarmResolver{Swarm: sw}
}
