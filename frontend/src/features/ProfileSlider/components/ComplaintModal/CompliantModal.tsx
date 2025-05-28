import CloseIcon from '@mui/icons-material/Close';
import { ReportOutlined } from '@mui/icons-material';
import { Button, FormControlLabel, Radio, RadioGroup, TextField } from '@mui/material';
import { useEffect, useState } from 'react';
import { createPortal } from 'react-dom';

import { useCreateComplaintMutation } from '../../../../entities/blocker';
import { reportOptions, reportToApi } from '../../../../entities/blocker/constants';

import styles from './complaintModal.module.css';

const reportReasons = Object.values(reportOptions);

export interface ComplaintModalProps {
  isOpen: boolean;
  onClose: () => void;
  targetUserId: string;
}

export const ComplaintModal = ({ isOpen, onClose, targetUserId }: ComplaintModalProps) => {
  const [selectedReason, setSelectedReason] = useState('');
  const [customText, setCustomText] = useState('');

  const [createComplaint] = useCreateComplaintMutation();

  useEffect(() => {
    if (!isOpen) {
      setSelectedReason('');
      setCustomText('');
    }
  }, [isOpen]);

  const handleSubmit = () => {
    if (!selectedReason && !customText) return;

    createComplaint({
      targetUserId: targetUserId,
      //@ts-ignore
      reasons: [reportToApi[selectedReason]],
      message: customText,
    });

    onClose();
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
          <div className={styles.wave1} />
        </div>

        <h2 className={styles.title}>
          <ReportOutlined />
          Пожаловаться
        </h2>

        <div className={styles.content}>
          <RadioGroup
            value={selectedReason}
            onChange={e => setSelectedReason(e.target.value)}
            className={styles.radioGroup}
          >
            {reportReasons.map(reason => (
              <FormControlLabel
                key={reason}
                value={reason}
                control={
                  <Radio
                    size="small"
                    sx={{
                      color: 'rgba(255, 255, 255, 0.7)',
                      '&.Mui-checked': {
                        color: 'white',
                      },
                    }}
                  />
                }
                label={reason}
                className={styles.radioLabel}
              />
            ))}
          </RadioGroup>

          <div className={styles.divider} />

          <TextField
            fullWidth
            multiline
            minRows={3}
            maxRows={6}
            value={customText}
            onChange={e => setCustomText(e.target.value)}
            placeholder="Опишите подробнее (необязательно)"
            variant="outlined"
            className={styles.textField}
          />
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
            disabled={!selectedReason && !customText}
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
            Отправить жалобу
          </Button>
        </div>
      </div>
    </div>,
    document.body,
  );
};
