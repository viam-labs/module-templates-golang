> [!NOTE]
> This is a module template. Read the [instructions](./Instructions.md) for a step by step guide on how to create your module.


# <INSERT NAME> modular resource

This module implements...
With this model, you can...

## Requirements

_Add instructions here for any requirements._

## Build and run

To use this module, follow the instructions to [add a module from the Viam Registry](https://docs.viam.com/registry/configure/#add-a-modular-resource-from-the-viam-registry) and select the `<INSERT API NAMESPACE>:<INSERT API NAME>:<INSERT MODEL>` model from the [`<INSERT MODEL>` module](https://app.viam.com/module/<INSERT API NAMESPACE>/<INSERT MODEL>).

## Configure your sensor

> [!NOTE]
> Before configuring your sensor you must [create a machine](https://docs.viam.com/manage/fleet/machines/#add-a-new-machine).

Navigate to the **Config** tab of your machine's page in [the Viam app](https://app.viam.com/).
Click on the **Components** subtab and click **Create component**.
Select the `<INSERT API NAME>` type, then select the `<INSERT MODEL>` model.
Click **Add module**, then enter a name for your sensor and click **Create**.

On the new component panel, copy and paste the following attribute template into your sensor’s **Attributes** box:

```json
{
  <INSERT SAMPLE ATTRIBUTES>
}
```

> [!NOTE]
> For more information, see [Configure a Machine](https://docs.viam.com/manage/configuration/).

### Attributes

The following attributes are available for `<INSERT API NAMESPACE>:<INSERT API NAME>:<INSERT MODEL>` sensor's:

| Name    | Type   | Inclusion    | Description |
| ------- | ------ | ------------ | ----------- |
| `todo1` | string | **Required** | TODO        |
| `todo2` | string | Optional     | TODO        |

### Example configuration

```json
{
  <INSERT SAMPLE CONFIGURATION(S)>
}
```

### Next steps

_Add any additional information you want readers to know and direct them towards what to do next with this module._
_For example:_

- To test your...
- To write code against your...

## Troubleshooting

_Add troubleshooting notes here._