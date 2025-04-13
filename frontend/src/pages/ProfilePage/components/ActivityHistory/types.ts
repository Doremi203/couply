export interface ActivityItem {
  type: string;
  user: string;
  date: string;
}

export interface ActivityHistoryProps {
  activityHistory: ActivityItem[];
  onBack: () => void;
  formatDate: (dateString: string) => string;
}
