package stacks

import (
	"github.com/cbadawi/go-stacks/stacks/controllers"
)

var feesController controllers.FeesController

var microblocksController controllers.MicroblocksController

var namesController controllers.NamesController

var rosettaController controllers.RosettaController

// init is an initialization function that sets up the controllers.
// It creates a configuration from the environment with a specified HTTP configuration and initializes the client.
// Then, it assigns the different controllers from the client to the corresponding variables for further use.
func init() {
	configuration := CreateConfigurationFromEnvironment(
		WithHttpConfiguration(
			CreateHttpConfiguration(
				WithTimeout(30),
			),
		),
	)

	client := NewClient(configuration)

	feesController = *client.FeesController()
	microblocksController = *client.MicroblocksController()
	namesController = *client.NamesController()
	rosettaController = *client.RosettaController()
}
