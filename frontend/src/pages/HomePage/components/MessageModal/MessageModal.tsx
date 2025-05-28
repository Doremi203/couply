import { SendOutlined } from '@mui/icons-material';
import CloseIcon from '@mui/icons-material/Close';
import { Box, Button, TextField } from '@mui/material';
import React, { useState } from 'react';
import { createPortal } from 'react-dom';

import { useLikeUserMutation } from '../../../../entities/matches';

import styles from './messageModal.module.css';

interface MessageModalProps {
  isOpen: boolean;
  onClose: () => void;
  targetUserId: string;
}

export const MessageModal: React.FC<MessageModalProps> = ({ isOpen, onClose, targetUserId }) => {
  const [likeUser] = useLikeUserMutation();
  const [message, setMessage] = useState('');

  const handleSubmit = async () => {
    if (!message.trim()) return;

    await likeUser({ targetUserId, message });
    onClose();
    setMessage('');
  };

  const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (e.target === e.currentTarget) {
      onClose();
    }
  };

  if (!isOpen) return null;

  return createPortal(
    <div className={styles.modalOverlay} onClick={handleOverlayClick}>
      <div className={styles.modalContent}>
        <button className={styles.closeButton} onClick={onClose}>
          <CloseIcon />
        </button>

        <div className={styles.decorations}>
          <div className={styles.star1} />
          <div className={styles.star2} />
          <div className={styles.circle} />
          <div className={styles.wave1} />
        </div>

        <h2 className={styles.title}>
          <SendOutlined style={{ marginRight: '8px' }} />
          Отправить сообщение
        </h2>

        <div className={styles.messageContent}>
          <Box sx={{ mt: 2 }}>
            <TextField
              fullWidth
              multiline
              autoFocus
              minRows={3}
              maxRows={6}
              value={message}
              onChange={e => setMessage(e.target.value)}
              placeholder="Введите ваше сообщение..."
              variant="outlined"
              sx={{
                '& .MuiOutlinedInput-root': {
                  borderRadius: '8px',
                  backgroundColor: 'rgba(255, 255, 255, 0.1)',
                  '& fieldset': {
                    borderColor: 'rgba(255, 255, 255, 0.3)',
                  },
                  '&:hover fieldset': {
                    borderColor: 'rgba(255, 255, 255, 0.5)',
                  },
                  '&.Mui-focused fieldset': {
                    borderColor: 'rgba(255, 255, 255, 0.8)',
                  },
                },
                '& .MuiInputBase-input': {
                  color: 'white',
                },
                '& .MuiInputLabel-root': {
                  color: 'rgba(255, 255, 255, 0.7)',
                },
                fontFamily: 'Jost',
              }}
            />
          </Box>
        </div>

        <div className={styles.actions}>
          <Button
            onClick={onClose}
            sx={{
              color: 'white',
              border: '2px solid rgba(255, 255, 255, 0.3)',
              '&:hover': {
                border: '2px solid rgba(255, 255, 255, 0.5)',
                backgroundColor: 'rgba(255, 255, 255, 0.1)',
              },
              fontFamily: 'Jost',
            }}
          >
            Отмена
          </Button>
          <Button
            variant="contained"
            onClick={handleSubmit}
            disabled={!message.trim()}
            startIcon={<SendOutlined />}
            sx={{
              background: 'rgba(255, 255, 255, 0.2)',
              '&:hover': {
                background: 'rgba(255, 255, 255, 0.3)',
              },
              '&.Mui-disabled': {
                background: 'rgba(255, 255, 255, 0.1)',
                color: 'rgba(255, 255, 255, 0.5)',
              },
              fontFamily: 'Jost',
            }}
          >
            Отправить
          </Button>
        </div>
      </div>
    </div>,
    document.body,
  );
};
