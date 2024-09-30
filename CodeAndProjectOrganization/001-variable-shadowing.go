package codeandprojectorganization

import "fmt"

func VariableShadowing(tracing bool) {
	var client string = "not created"
	if tracing {
		// This client is shadowed because of the declaration.
		// client is within the scope of the if block
		client, err := createClientWithTracing()
		// This is how we would do this without shadowing client
		// client, err = createClientWithTracing()

		if err != nil {
			// some error handling
			return
		}
		fmt.Println(client)
	} else {
		// This client is shadowed because of the declaration.
		// client is within the scope of the if block
		client, err := createClientWithoutTracing()
		// This is how we would do this without shadowing client
		// client, err = createClientWithoutTracing()

		if err != nil {
			// some error handling
			return
		}
		fmt.Println(client)
	}
	// This will always print "not created" because the other assignments were shadowed
	fmt.Println(client)
}

func createClientWithTracing() (string, error) {
	return "Tracing", nil
}
func createClientWithoutTracing() (string, error) {
	return "No trace", nil
}
