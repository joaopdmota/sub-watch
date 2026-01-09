import type { Meta, StoryObj } from '@storybook/react';
import { KPICard } from '@/features/dashboard/components/KPICard';

const meta = {
  title: 'Features/Dashboard/KPICard',
  component: KPICard,
  tags: ['autodocs'],
} satisfies Meta<typeof KPICard>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    title: 'Total Monthly',
    value: 'R$ 250,50',
  },
};

export const WithSubtext: Story = {
  args: {
    title: 'Yearly Projection',
    value: 'R$ 3.000,00',
    subtext: '+12% from last year',
  },
};
