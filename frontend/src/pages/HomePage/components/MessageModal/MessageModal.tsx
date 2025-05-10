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
// @ts-ignore
import { forwardRef, useState } from 'react';

// @ts-nocheck

// const Transition = forwardRef(function Transition(props, ref) {
//   return <Slide direction="up" ref={ref} {...props} />;
// });

// @ts-ignore
export const MessageModal = ({ isOpen, onClose }) => {
  const [message, setMessage] = useState('');

  const handleSubmit = () => {
    if (!message.trim()) return;

    // Здесь обработка отправки сообщения
    console.log('Отправлено сообщение:', message);
    onClose();
  };

  return (
    <Dialog
      open={isOpen}
      onClose={onClose}
      // TransitionComponent={Transition}
      fullWidth
      maxWidth="sm"
      PaperProps={{
        style: {
          position: 'fixed',
          bottom: 0,
          margin: 0,
          borderRadius: '12px 12px 0 0',
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
