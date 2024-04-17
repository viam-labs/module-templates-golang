// Package customcamera implements a camera where all methods are unimplemented.
// It extends the built-in resource subtype camera and implements methods to handle resource construction and attribute configuration.

package customcamera

import (
    "context"
    "errors"

    "go.viam.com/rdk/components/camera"
    "go.viam.com/rdk/pointcloud"
    "go.viam.com/rdk/logging"
    "go.viam.com/rdk/resource"
    "go.viam.com/rdk/rimage/transform"
    "go.viam.com/rdk/gostream"
    "go.viam.com/utils"
)

// Here is where we define your new model's colon-delimited-triplet (viam-labs:go-module-templates-camera:customcamera)
// viam-labs = namespace, go-module-templates-camera = repo-name, customcamera = model name.
// TODO: Change model namespace, family (often the repo-name), and model. For more information see https://docs.viam.com/registry/create/#name-your-new-resource-model
var (
    Model            = resource.NewModel("viam-labs", "go-module-templates-camera", "customcamera")
    errUnimplemented = errors.New("unimplemented")
)

func init() {
    resource.RegisterComponent(camera.API, Model,
        resource.Registration[camera.Camera, *Config]{
            Constructor: newCustomCamera,
        },
    )
}

// TODO: Change the Config struct to contain any values that you would like to be able to configure from the attributes field in the component
// configuration. For more information see https://docs.viam.com/build/configure/#components
type Config struct {
    ArgumentOne int    `json:"one"`
    ArgumentTwo string `json:"two"`
}

// Validate validates the config and returns implicit dependencies.
// TODO: Change the Validate function to validate any config variables.
func (cfg *Config) Validate(path string) ([]string, error) {
    if cfg.ArgumentOne == 0 {
        return nil, utils.NewConfigValidationFieldRequiredError(path, "one")
    }

    if cfg.ArgumentTwo == "" {
        return nil, utils.NewConfigValidationFieldRequiredError(path, "two")
    }

    // TODO: return implicit dependencies if needed as the first value
    return []string{}, nil
}

// Constructor for a custom camera that creates and returns a customCamera.
// TODO: update the customCamera struct and the initialization.
func newCustomCamera(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (camera.Camera, error) {
    // This takes the generic resource.Config passed down from the parent and converts it to the
    // model-specific (aka "native") Config structure defined above, making it easier to directly access attributes.
    conf, err := resource.NativeConfig[*Config](rawConf)
    if err != nil {
        return nil, err
    }

    // Create a cancelable context for custom camera
    cancelCtx, cancelFunc := context.WithCancel(context.Background())

    c := &customCamera{
        name:        rawConf.ResourceName(),
        logger:      logger,
        cfg:         conf,
        cancelCtx:   cancelCtx,
        cancelFunc:  cancelFunc,
    }

    // TODO: If your custom component has dependencies, perform any checks you need to on them.

    // The Reconfigure() method changes the values on the customCamera based on the attributes in the component config
    if err := c.Reconfigure(ctx, deps, rawConf); err != nil {
        logger.Error("Error configuring module with ", err)
        return nil, err
    }

    return c, nil
}

// TODO: update the customCamera struct with any fields you require.
type customCamera struct {
    name   resource.Name
    logger logging.Logger
    cfg    *Config

    cancelCtx  context.Context
    cancelFunc func()

    argumentOne int
    argumentTwo string
}

func (c *customCamera) Name() resource.Name {
    return c.name
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the camera, throw a NewMustBuildError.
func (c *customCamera) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
    cameraConfig, err := resource.NativeConfig[*Config](conf)
    if err != nil {
        c.logger.Warn("Error reconfiguring module with ", err)
        return err
    }

    c.argumentOne = cameraConfig.ArgumentOne
    c.argumentTwo = cameraConfig.ArgumentTwo
    c.name = conf.ResourceName()
    c.logger.Info("one is now configured to: ", c.argumentOne)
    c.logger.Info("two is now configured to ", c.argumentTwo)

    return nil
}

// Images is for getting simultaneous images from different sensors.
// If the underlying source did not specify an Images function, a default is applied.
// The default should return a list of 1 image from ReadImage, and the current time.
func (c *customCamera) Images(ctx context.Context) ([]camera.NamedImage, resource.ResponseMetadata, error) {
    // TODO: Obtain and return simultaneous images from different sensors.
    c.logger.Error("Images method unimplemented")
    return nil, resource.ResponseMetadata{}, errUnimplemented
}

func (c *customCamera) Stream(ctx context.Context, errHandlers ...gostream.ErrorHandler) (gostream.VideoStream, error) {
    // TODO: Obtain and return the camera stream.
    c.logger.Error("Stream method unimplemented")
    return nil, errUnimplemented
}

// NextPointCloud returns the next PointCloud from the camera, or will error if not supported.
func (c *customCamera) NextPointCloud(ctx context.Context) (pointcloud.PointCloud, error) {
    // TODO: Obtain and return the next PointCloud.
    c.logger.Error("NextPointCloud method unimplemented")
    return nil, errUnimplemented
}

// DoCommand is a place to add additional commands to extend the camera API. This is optional.
func (c *customCamera) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
    c.logger.Error("DoCommand method unimplemented")
    return nil, errUnimplemented
}

func (c *customCamera) Properties(ctx context.Context) (camera.Properties, error) {
    // TODO: Obtain and return an object wrapping the camera's properties.
    c.logger.Error("Properties method unimplemented")
    return camera.Properties{}, errUnimplemented
}

func (c *customCamera) Projector(ctx context.Context) (transform.Projector, error) {
    // TODO: Obtain and return an object wrapping the camera's properties.
    c.logger.Error("Projector method unimplemented")
    return nil, errUnimplemented
}

// Close closes the underlying generic.
func (c *customCamera) Close(ctx context.Context) error {
    c.cancelFunc()
    return nil
}
