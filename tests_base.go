package stacksblockchainapi

var feesController FeesController

var microblocksController MicroblocksController

var namesController NamesController

var rosettaController RosettaController

// init is an initialization function that sets up the controllers.
// It creates a configuration from the environment with a specified HTTP configuration and initializes the client.
// Then, it assigns the different controllers from the client to the corresponding variables for further use.
func init() {
	config := CreateConfigurationFromEnvironment(
		WithHttpConfiguration(
			CreateHttpConfiguration(
				WithTimeout(30),
			),
		),
	)

	client := NewClient(config)

	feesController = *client.FeesController()
	microblocksController = *client.MicroblocksController()
	namesController = *client.NamesController()
	rosettaController = *client.RosettaController()
}
