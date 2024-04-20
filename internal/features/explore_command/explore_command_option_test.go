package explore_command

import (
	qec "query_fetch/query_cache/cache_eviction_config"
	"testing"

	f "test_tools/fixtures"
)

func TestWithAPIEndpoint(t *testing.T) {
	runWithAPIEndpointTests := func() *ExploreCommand {
		command := &ExploreCommand{}
		WithAPIEndpoint(f.APIEndpoint)(command)
		return command
	}

	t.Run("should set the API endpoint for the ExploreCommand instance.", func(t *testing.T) {
		if command := runWithAPIEndpointTests(); command.apiEndpoint != f.APIEndpoint {
			t.Errorf("Expected apiEndpoint to be %s, but got %s", f.APIEndpoint, command.apiEndpoint)
		}
	})
}

func TestWithEvictionConfig(t *testing.T) {
	runWithEvictionConfigTests := func() (actual *qec.QueryEvictionConfig, expected *qec.QueryEvictionConfig) {
		command := &ExploreCommand{}
		ec := qec.NewQueryEvictionConfig()
		WithEvictionConfig(ec)(command)
		return command.ec, ec
	}

	t.Run("should set the eviction configuration for the ExploreCommand instance.", func(t *testing.T) {
		if actual, expected := runWithEvictionConfigTests(); actual != expected {
			t.Errorf("Expected ec to be %v, but got %v", expected, actual)
		}
	})
}

func TestWithDescription(t *testing.T) {
	runWithDescriptionTests := func() (command *ExploreCommand, description string) {
		command = &ExploreCommand{}
		description = "Command Description"
		WithDescription(description)(command)
		return command, description
	}

	t.Run("should set the description for the ExploreCommand instance.", func(t *testing.T) {
		if command, description := runWithDescriptionTests(); command.description != description {
			t.Errorf("Expected description to be %s, but got %s", description, command.description)
		}
	})
}

func TestWithListMarker(t *testing.T) {
	runWithListMarkerTests := func() (command *ExploreCommand, listMarker string) {
		command = &ExploreCommand{}
		listMarker = ">"
		WithListMarker(listMarker)(command)
		return command, listMarker
	}

	t.Run("should set the list marker for the ExploreCommand instance.", func(t *testing.T) {
		if command, listMarker := runWithListMarkerTests(); command.listMarker != listMarker {
			t.Errorf("Expected listMarker to be %s, but got %s", listMarker, command.listMarker)
		}
	})
}

func TestWithListTitle(t *testing.T) {
	runWithListTitleTests := func() (command *ExploreCommand, listTitle string) {
		command = &ExploreCommand{}
		listTitle = "List Title"
		WithListTitle(listTitle)(command)
		return command, listTitle
	}

	t.Run("should set the list title for the ExploreCommand instance.", func(t *testing.T) {
		if command, listTitle := runWithListTitleTests(); command.listTitle != listTitle {
			t.Errorf("Expected listTitle to be %s, but got %s", listTitle, command.listTitle)
		}
	})
}

func TestWithName(t *testing.T) {
	runWithNameTests := func() (command *ExploreCommand, name string) {
		command = &ExploreCommand{}
		name = "Command Name"
		WithName(name)(command)
		return command, name
	}

	t.Run("should set the name for the ExploreCommand instance.", func(t *testing.T) {
		if command, name := runWithNameTests(); command.name != name {
			t.Errorf("Expected name to be %s, but got %s", name, command.name)
		}
	})
}

func TestWithNoEncountersFoundErrorMessage(t *testing.T) {
	runWithNoEncountersFoundErrorMessageTests := func() (command *ExploreCommand, message string) {
		command = &ExploreCommand{}
		message = "No encounters found."
		WithNoEncountersFoundErrorMessage(message)(command)
		return command, message
	}

	t.Run("should set the no encounters found error message for the ExploreCommand instance.", func(t *testing.T) {
		if command, message := runWithNoEncountersFoundErrorMessageTests(); command.noEncountersFoundErrorMessage != message {
			t.Errorf("Expected noEncountersFoundErrorMessage to be %s, but got %s", message, command.noEncountersFoundErrorMessage)
		}
	})
}

func TestWithNoEnteredArgsErrorMessage(t *testing.T) {
	runWithNoEnteredArgsErrorMessageTests := func() (command *ExploreCommand, message string) {
		command = &ExploreCommand{}
		message = "No entered arguments."
		WithNoEnteredArgsErrorMessage(message)(command)
		return command, message
	}

	t.Run("should set the no entered arguments error message for the ExploreCommand instance.", func(t *testing.T) {
		if command, message := runWithNoEnteredArgsErrorMessageTests(); command.noEnteredArgsErrorMessage != message {
			t.Errorf("Expected noEnteredArgsErrorMessage to be %s, but got %s", message, command.noEnteredArgsErrorMessage)
		}
	})
}
