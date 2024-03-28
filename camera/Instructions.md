## Create your own camera module

Viam provides built-in support for a variety of different components and services, but you can add support for your own custom resources by creating a module.

This is a module template that implements one camera. You can use it as a starting point to write a module that adds support for the camera you need.

### Instructions

#### 1. Copy this folder into your workspace

To begin copy the [`camera`](../) folder into another folder on your machine and initialize git. You can rename the folder if you wish to describe your custom camera better.

Inside your copied folder, you have the following files:

- [cmd](./cmd/)
  - [module](./module/)
    - [cmd.go](./cmd.go): The entry point main program file, which imports the API for the model, defines the `main()` function that registers the model with `viam-server` and creates and starts the module.
- [.gitignore](./.gitignore): A sensible default `.gitignore` file that ensures that no binaries are added to git.
- [customcamera.go](./customcamera.go): A template for a custom camera implementation.
- [go.mod](./go.mod): Your go dependencies. For more information see [the Go documenation](https://go.dev/doc/tutorial/create-module).
- [go.sum](./go.sum): Checksums for the exact contents of each dependency at the time it is added to your module. For more information see [the Go documentation](https://go.dev/doc/tutorial/create-module).
- [Makefile](./Makefile): A default makefile with command `make module`, `make bin/customcamera`, `make lint` and more.
- [README.md](./README.md): A readme template for your module.

#### 2. Update `cmd/module/cmd.go`

Update the entry point file where there are TODOs. Each TODO explains what you need to do. Your updates will change the name `customcamera` to something more descriptive for the custom camera you are implementing, for example `raspberrypicam`.

#### 3. Update `cmd/module/customcamera.go`

Update the custom camera implementation where there are TODOs. Each TODO explains what you need to do. Your updates will:

- Give a name to the model. For more information, see the [docs on model naming](https://docs.viam.com/registry/create/#name-your-new-resource-model).
- Add any configurable values for the camera's attributes field. For more information, see the [docs on component configuration](https://docs.viam.com/build/configure/#components). Additionally you will update the `Validate()` function to validate that config values meet requirements, for example that required fields exist.
- Update the constructor for your camera as well as the `customCamera` struct.
- Update the implementation of the Camera API:
  - `Images()`
  - `Stream()`
  - `NextPointCloud()`
  - `DoCommand()`
  - `Properties()`
  - `Projector()`
  - `Close()`

#### 4. Compile and build

If you renamed your package and are building locally you need to specify that your package name should be replaced with the local version of the package:

```
go mod edit -replace github.com/viam-labs/module-templates-golang/camera/customcamera=./customcamera
```

Then run the following commands to update the `rdk` to the latest build and synchronize and update the project's go dependencies:

```sh
make updaterdk
```

Next, run `make bin/customcamera` to build the executable.

#### 5. Test your module

##### Test just your camera

The simplest way of testing your camera, is to run a version of `viam-server` with just your camera configured.

1. Edit the [`cmd/remote/cmd.go` file](./cmd/remote/cmd.go). The code there should configure and run a robot with your modular resource configured.
1. Run `make bin/remoteserver` to compile the [`cmd/remote/cmd.go` file](./cmd/module/cmd.go).
1. Run `./bin/remoteserver my_camera_name` and specify any additional commandline arguments that your code passes  to the modular resource attributes. This will instantiate a simplified instance of `viam-server` with your modular resource.

##### Test your camera with an operational machine

If you want to test your module with a lot of other resources you can:

- Test your module by following the steps in [Test just your camera](#test-just-your-camera) to deploy your module and use `localhost:8083` to [add it as a remote](https://docs.viam.com/build/configure/parts-and-remotes/#configure-a-remote). If you make changes to your code, you will only need to restart the remote that is running your camera.
- Test your module by adding it as [a local module](https://docs.viam.com/registry/configure/#add-a-local-module). If you make changes to your code, you will need to restart `viam-server` for the changes to take effect.

#### 6. Upload your module to the registry

To upload the module to the registry, which will allow you to add it to machines without manually moving the files to each machine and specifying the file path, follow the instructions to [Upload your Own Modules to the Viam Registry](https://docs.viam.com/registry/upload/).

#### 7. Update your Readme

Update the [README.md](./README.md) file with the configuration instructions for your module.