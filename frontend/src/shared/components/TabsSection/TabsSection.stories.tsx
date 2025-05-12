import type { Meta, StoryObj } from '@storybook/react';
import { useState } from 'react';

import { TabsSection } from './TabsSection';

// Define a wrapper component to handle state for the tabs
const TabsSectionWrapper = <T extends string>({
  initialTab,
  tabs,
  tabLabels,
}: {
  initialTab: T;
  tabs: T[];
  tabLabels?: Record<T, string>;
}) => {
  const [activeTab, setActiveTab] = useState<T>(initialTab);

  return (
    <TabsSection
      tabs={tabs}
      activeTab={activeTab}
      onTabChange={setActiveTab}
      tabLabels={tabLabels}
    />
  );
};

const meta = {
  title: 'Shared/TabsSection',
  component: TabsSection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section with tabs for navigation between different content sections.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof TabsSection>;

export default meta;
type Story = StoryObj<typeof TabsSection>;

// Define some example tab types
type ProfileTabs = 'info' | 'photos' | 'interests';

export const ProfileTabsExample: Story = {
  render: () => {
    const tabs: ProfileTabs[] = ['info', 'photos', 'interests'];
    const tabLabels: Record<ProfileTabs, string> = {
      info: 'Info',
      photos: 'Photos',
      interests: 'Interests',
    };

    return <TabsSectionWrapper initialTab="info" tabs={tabs} tabLabels={tabLabels} />;
  },
};
