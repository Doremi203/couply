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
  Slide,
  TextField,
} from '@mui/material';
import { forwardRef, useEffect, useState } from 'react';

const Transition = forwardRef(function Transition(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

const reportReasons = [
  'Спам',
  'Мошенничество',
  'Оскорбительный контент',
  'Некорректная информация',
  'Другая причина',
];

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
      TransitionComponent={Transition}
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
