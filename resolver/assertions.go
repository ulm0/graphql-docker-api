package resolver

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/swarm"
)

/*	Specific assertion funcs */

// ptrIPNetString does type assertion []*registry.NetIPNet to *[]string
func ptrIPNetString(nets []*registry.NetIPNet) *[]string {
	// Starts an slice with len 0 and then appends elements to it
	// So it will not output empty string
	// Such as:
	// "registryConfig": {
	//      "allowNondistributableArtifactsCidrs": [
	//        "", # This occurs when the slice is started with len 1
	//      ]
	ipNets := make([]string, 0)
	for _, ipNet := range nets {
		ipNets = append(ipNets, ipNet.String())
	}
	return &ipNets
}

/*	Common assertion funcs */

// ptrString gets does type assertion,
// gets reference from []byte and string
func ptrString(input interface{}) *string {
	switch input := input.(type) {
	case []byte:
		ptr := string(input)
		return &ptr
	case string:
		return &input
	default:
		return nil
	}
}

// ptrCommit does type assertion of types.Commit
func ptrCommit(c types.Commit) *systemInfoCommitResolver {
	commit := systemInfoCommitResolver(c)
	return &commit
}

func ptrSwVer(c swarm.Version) *swarmVersionResolver {
	ver := swarmVersionResolver(c)
	return &ver
}

func ptrInt32(input interface{}) *int32 {
	switch input := input.(type) {
	case int:
		ptr := int32(input)
		return &ptr
	case int8:
		ptr := int32(input)
		return &ptr
	case int16:
		ptr := int32(input)
		return &ptr
	case int32:
		return &input
	case int64:
		ptr := int32(input)
		return &ptr
	case uint:
		ptr := int32(input)
		return &ptr
	case uint8:
		ptr := int32(input)
		return &ptr
	case uint16:
		ptr := int32(input)
		return &ptr
	case uint32:
		ptr := int32(input)
		return &ptr
	case uint64:
		ptr := int32(input)
		return &ptr
	case *int:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int8:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int16:
		value := *input
		ptr := int32(value)
		return &ptr
	case *int32:
		return input
	case *int64:
		value := *input
		ptr := int32(value)
		return &ptr
	case uintptr:
		value := input
		ptr := int32(value)
		return &ptr
	case *uint8:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint16:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint32:
		value := *input
		ptr := int32(value)
		return &ptr
	case *uint64:
		value := *input
		ptr := int32(value)
		return &ptr
	default:
		// Wrong type
		return nil
	}
}
