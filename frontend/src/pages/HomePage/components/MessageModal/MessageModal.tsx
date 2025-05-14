import { SendOutlined } from '@mui/icons-material';
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
} from '@mui/material';
import { useState } from 'react';

import { useLikeUserMutation } from '../../../../entities/matches';

// @ts-ignore
export const MessageModal = ({ isOpen, onClose, targetUserId }) => {
  const [likeUser] = useLikeUserMutation();
  const [message, setMessage] = useState('');

  const handleSubmit = async () => {
    if (!message.trim()) return;

    await likeUser({ targetUserId, message });

    console.log('Отправлено сообщение:', message);
    onClose();
  };

  return (
    <Dialog
      open={isOpen}
      onClose={onClose}
      fullWidth
      maxWidth="sm"
      PaperProps={{
        sx: {
          position: 'fixed',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          margin: 0,
          borderRadius: '12px',
          maxHeight: '90vh',
          overflow: 'auto',
        },
      }}
    >
      <DialogTitle sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
        <SendOutlined color="primary" />
        Отправить сообщение
      </DialogTitle>

      <DialogContent>
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
              },
            }}
          />
        </Box>
      </DialogContent>

      <DialogActions sx={{ p: 2 }}>
        <Button onClick={onClose} color="inherit">
          Отмена
        </Button>
        <Button
          variant="contained"
          onClick={handleSubmit}
          disabled={!message.trim()}
          startIcon={<SendOutlined />}
        >
          Отправить
        </Button>
      </DialogActions>
    </Dialog>
  );
};
