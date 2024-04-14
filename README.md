# Golang Module templates

Viam provides built-in support for a variety of different components and services. If your resource is not natively supported, you can add support for a custom resources by creating a module.

This repository provides templates for implementing the following components:

- [camera](./camera/)
- [movement sensor](./movementsensor/)
- [power sensor](./powersensor/)
- [sensor](./sensor/)
- [servo](./servo/)

You can use these templates as a starting point to write a module that adds support for the component you need.

If the resource you need to implement is not among the examples in this repository, you can still reference these examples and follow the instructions. For additional reference materials, review [existing modules in the Viam Registry](https://docs.viam.com/registry/) that implement a similar resource.

> [!NOTE]
> Before creating your own module, you can also check whether someone else has already created and shared a module for the component you need by [searching through the Viam Registry](https://docs.viam.com/registry/).