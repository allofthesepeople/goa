// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package cli

import (
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa"
	dividerc "goa.design/goa/examples/error/gen/grpc/divider/client"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `divider (integer-divide|divide)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` divider integer-divide --message '{
      "a": 1338266005399228665,
      "b": 7245195139064803075
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		dividerFlags = flag.NewFlagSet("divider", flag.ContinueOnError)

		dividerIntegerDivideFlags       = flag.NewFlagSet("integer-divide", flag.ExitOnError)
		dividerIntegerDivideMessageFlag = dividerIntegerDivideFlags.String("message", "", "")

		dividerDivideFlags       = flag.NewFlagSet("divide", flag.ExitOnError)
		dividerDivideMessageFlag = dividerDivideFlags.String("message", "", "")
	)
	dividerFlags.Usage = dividerUsage
	dividerIntegerDivideFlags.Usage = dividerIntegerDivideUsage
	dividerDivideFlags.Usage = dividerDivideUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "divider":
			svcf = dividerFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "divider":
			switch epn {
			case "integer-divide":
				epf = dividerIntegerDivideFlags

			case "divide":
				epf = dividerDivideFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "divider":
			c := dividerc.NewClient(cc, opts...)
			switch epn {
			case "integer-divide":
				endpoint = c.IntegerDivide()
				data, err = dividerc.BuildIntegerDividePayload(*dividerIntegerDivideMessageFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = dividerc.BuildDividePayload(*dividerDivideMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// dividerUsage displays the usage of the divider command and its subcommands.
func dividerUsage() {
	fmt.Fprintf(os.Stderr, `Service is the divider service interface.
Usage:
    %s [globalflags] divider COMMAND [flags]

COMMAND:
    integer-divide: IntegerDivide implements integer_divide.
    divide: Divide implements divide.

Additional help:
    %s divider COMMAND --help
`, os.Args[0], os.Args[0])
}
func dividerIntegerDivideUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] divider integer-divide -message JSON

IntegerDivide implements integer_divide.
    -message JSON: 

Example:
    `+os.Args[0]+` divider integer-divide --message '{
      "a": 1338266005399228665,
      "b": 7245195139064803075
   }'
`, os.Args[0])
}

func dividerDivideUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] divider divide -message JSON

Divide implements divide.
    -message JSON: 

Example:
    `+os.Args[0]+` divider divide --message '{
      "a": 0.7822693555171186,
      "b": 0.5749246657891343
   }'
`, os.Args[0])
}
