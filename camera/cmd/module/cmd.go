// Package main is a module which serves the customcamera custom model.
package main

import (
    "context"

    "go.viam.com/rdk/components/camera"
    "go.viam.com/rdk/logging"
    "go.viam.com/rdk/module"
    "go.viam.com/utils"

    // Import your local package "customcamera"
    // TODO: Update this path if your custom resource is in a different location,
    // or has a different name:
    "github.com/viam-labs/module-templates-golang/camera"
)

func main() {
    // NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
    // "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
    // TODO: Change the name of the logger from customcamera to the name of the module your are creating
    utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("customcamera"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
    myModule, err := module.NewModuleFromArgs(ctx, logger)
    if err != nil {
        return err
    }

    // Adds the preregistered camera component API to the module for the new model.
    // TODO: Update the name of your package customcamera
    err = myModule.AddModelFromRegistry(ctx, camera.API, customcamera.Model)
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