import MessageOutlinedIcon from '@mui/icons-material/MessageOutlined';

import styles from './messageButton.module.css';

interface BackCircleButtonProps {
  onClick: () => void;
  className?: string;
}

export const MessageButton = ({ onClick, className }: BackCircleButtonProps) => {
  return (
    <div className={`${styles.circle} ${className || ''}`} onClick={onClick}>
      <MessageOutlinedIcon className={styles.icon} />
    </div>
  );
};

export default MessageButton;
