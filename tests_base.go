package stacksblockchainapi

import (
	"github.com/cbadawi/stacks-go-draft/controllers"
)

var FeesController controllers.FeesController

var MicroblocksController controllers.MicroblocksController

var NamesController controllers.NamesController

var RosettaController controllers.RosettaController

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

	FeesController = *client.FeesController()
	MicroblocksController = *client.MicroblocksController()
	NamesController = *client.NamesController()
	RosettaController = *client.RosettaController()
}
