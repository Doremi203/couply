import React, { ReactNode } from 'react'
import styles from './iconButton.module.css'

interface CircleIconButtonProps {
  onClick: () => void
  className?: string
  iconClassName?: string
  children: ReactNode
}

export const IconButton = ({onClick, className, iconClassName, children}: CircleIconButtonProps) => {
  return (
    <div className={`${styles.iconCircle} ${className || ''}`} onClick={onClick}>
      <div className={`${styles.iconContainer} ${iconClassName || ''}`}>
        {children}
      </div>
    </div>
  )
}

export default IconButton;