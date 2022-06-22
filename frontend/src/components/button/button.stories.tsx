import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { Button } from './button';

export default {
  component: Button,
} as ComponentMeta<typeof Button>;

const Template: ComponentStory<typeof Button> = (args) => (
  <Button {...args}/>
);

export const Primary = Template.bind({});
Primary.args = {
  variant: 'primary',
  children: ('Test'),
};

export const Secondary = Template.bind({});
Secondary.args = {
  variant: 'secondary',
  children: ('Test'),
};

export const Error = Template.bind({});
Error.args = {
  variant: 'error',
  children: ('Test'),
};
