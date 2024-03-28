// Package main is a module which serves the custommovementsensor custom model.
package main

import (
    "context"

    "go.viam.com/rdk/components/movementsensor"
    "go.viam.com/rdk/logging"
    "go.viam.com/rdk/module"
    "go.viam.com/utils"

    // Import your local package "custommovementsensor"
    // TODO: Update this path if your custom resource is in a different location,
    // or has a different name:
    "github.com/viam-labs/module-templates-golang/movementsensor"
)

func main() {
    // NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
    // "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
    // TODO: Change the name of the logger from custommovementsensor to the name of the module your are creating
    utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("custommovementsensor"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
    myModule, err := module.NewModuleFromArgs(ctx, logger)
    if err != nil {
        return err
    }

    // Adds the preregistered movementsensor component API to the module for the new model.
    // TODO: Update the name of your package custommovementsensor
    err = myModule.AddModelFromRegistry(ctx, movementsensor.API, custommovementsensor.Model)
    if err != nil {
        return err
    }

    err = myModule.Start(ctx)
    defer myModule.Close(ctx)
    if err != nil {
        return err
    }
    <-ctx.Done()
    return nil
}
