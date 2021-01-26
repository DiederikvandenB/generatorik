import { Meta, Story, ArgsTable, Canvas } from '@storybook/addon-docs/blocks';
import Component from './{{.ComponentName}}';

<Meta
  title="Components/{{.ComponentName}}"
  component={Component}
/>

# {{.ComponentName}}
<Canvas>
  <Story name="{{.ComponentName}}">
    {args => <Component {...args}>{args.label}</Component>}
  </Story>
</Canvas>

<ArgsTable story="{{.ComponentName}}" />

## Variants
