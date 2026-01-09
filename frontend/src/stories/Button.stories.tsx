import type { Meta, StoryObj } from '@storybook/react';
import { Button } from '@/components/ui';

const meta = {
  title: 'UI/Button',
  component: Button,
  tags: ['autodocs'],
  argTypes: {
    children: { control: 'text' },
    className: { control: 'text' },
  },
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: 'Button',
  },
};

export const Secondary: Story = {
  args: {
    children: 'Secondary',
    className: 'bg-secondary text-secondary-foreground hover:bg-secondary/80',
  },
};

export const Destructive: Story = {
    args: {
        children: 'Destructive',
        className: 'bg-destructive text-destructive-foreground hover:bg-destructive/90',
    }
}
