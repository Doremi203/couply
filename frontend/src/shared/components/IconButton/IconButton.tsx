import { ReactNode } from 'react'
import styles from './iconButton.module.css'

interface CircleIconButtonProps {
  onClick: () => void
  className?: string
  iconClassName?: string
  children: ReactNode
  touchFriendly?: boolean // Add option for larger touch area
}

export const IconButton = ({
  onClick,
  className,
  iconClassName,
  children,
  touchFriendly = false
}: CircleIconButtonProps) => {
  return (
    <div
      className={`
        ${styles.iconCircle}
        ${touchFriendly ? styles.touchFriendly : ''}
        ${className || ''}
      `}
      onClick={onClick}
    >
      <div className={`${styles.iconContainer} ${iconClassName || ''}`}>
        {children}
      </div>
    </div>
  )
}

export default IconButton;