import { ReportOutlined } from '@mui/icons-material';
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  FormControlLabel,
  Radio,
  RadioGroup,
  TextField,
} from '@mui/material';
import { useEffect, useState } from 'react';

const reportReasons = [
  'Спам',
  'Фейковый профиль',
  'Оскорбительное поведение',
  'Неприемлемый контент',
  'Возраст',
  'Другая причина',
];

//@ts-ignore
export const ComplaintModal = ({ isOpen, onClose }) => {
  const [selectedReason, setSelectedReason] = useState('');
  const [customText, setCustomText] = useState('');

  useEffect(() => {
    if (!isOpen) {
      setSelectedReason('');
      setCustomText('');
    }
  }, [isOpen]);

  const handleSubmit = () => {
    if (!selectedReason && !customText) return;

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
      sx={{
        '& .MuiBackdrop-root': {
          backgroundColor: 'rgba(0, 0, 0, 0.5)',
        },
      }}
    >
      <DialogTitle
        sx={{
          display: 'flex',
          alignItems: 'center',
          gap: 1,
          backgroundColor: theme => theme.palette.background.paper,
          boxShadow: 1,
        }}
      >
        <ReportOutlined color="error" />
        Пожаловаться
      </DialogTitle>

      <DialogContent>
        <Box sx={{ mt: 1 }}>
          <RadioGroup value={selectedReason} onChange={e => setSelectedReason(e.target.value)}>
            {reportReasons.map(reason => (
              <FormControlLabel
                key={reason}
                value={reason}
                control={<Radio size="small" />}
                label={reason}
                sx={{ mb: 1 }}
              />
            ))}
          </RadioGroup>

          <Divider sx={{ my: 2 }} />

          <TextField
            fullWidth
            multiline
            minRows={3}
            maxRows={6}
            value={customText}
            onChange={e => setCustomText(e.target.value)}
            placeholder="Опишите подробнее (необязательно)"
            variant="outlined"
          />
        </Box>
      </DialogContent>

      <DialogActions sx={{ p: 2 }}>
        <Button onClick={onClose} color="inherit">
          Отмена
        </Button>
        <Button variant="outlined" onClick={handleSubmit} disabled={!selectedReason && !customText}>
          Отправить жалобу
        </Button>
      </DialogActions>
    </Dialog>
  );
};
