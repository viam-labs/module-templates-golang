package main

import (
    "context"
    "os"
    "strconv"

    "go.viam.com/rdk/module"

    "go.viam.com/rdk/components/camera"
    "go.viam.com/rdk/config"
    "go.viam.com/rdk/logging"
    "go.viam.com/rdk/resource"
    robotimpl "go.viam.com/rdk/robot/impl"
    "go.viam.com/rdk/robot/web"
    rdkutils "go.viam.com/rdk/utils"
    "go.viam.com/utils"

    "github.com/viam-labs/module-templates-golang/camera"
)

func main() {
    // NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
    // "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
    utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("testcamera"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {

    netconfig := config.NetworkConfig{}
    netconfig.BindAddress = "0.0.0.0:8083"

    if err := netconfig.Validate(""); err != nil {
        return err
    }

    arg_2_converted, _ := strconv.Atoi(os.Args[2])

	// Update the Attributes and ConvertedAttributes with the attributes your modular resource should receive
    conf := &config.Config{
        Network: netconfig,
        Components: []resource.Config{
            {
                Name:  os.Args[1],
                API:   camera.API,
                Model: customcamera.Model,
                Attributes: rdkutils.AttributeMap{
                    "one": arg_2_converted,
                    "two": os.Args[3],
                },
                ConvertedAttributes: &customcamera.Config{
                    ArgumentOne: arg_2_converted,
                    ArgumentTwo: os.Args[3],
                },
            },
        },
    }

    myRobot, err := robotimpl.New(ctx, conf, logger)
    if err != nil {
        return err
    }

    return web.RunWebWithConfig(ctx, myRobot, conf, logger)
}