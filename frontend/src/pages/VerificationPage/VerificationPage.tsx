import React, { useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import PageHeader from '../../shared/components/PageHeader';

const VerificationPage: React.FC = () => {
  const [photo, setPhoto] = useState<File | null>(null);
  const [status, setStatus] = useState<string | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const navigate = useNavigate();

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) setPhoto(file);
  };

  const handleSend = async () => {
    if (!photo) return;
    setIsUploading(true);
    setStatus(null);
    const token = localStorage.getItem('token');
    if (!token) {
      setStatus('Токен не найден');
      setIsUploading(false);
      return;
    }
    const toBase64 = (file: File) =>
      new Promise<string>((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result as string);
        reader.onerror = () => reject();
      });
    let base64Photo = '';
    try {
      base64Photo = await toBase64(photo);
    } catch {
      setStatus('Ошибка чтения файла');
      setIsUploading(false);
      return;
    }
    const body = {
      token,
      bucket: 'couply-verification-photos',
      photo: base64Photo,
    };
    try {
      const response = await fetch('https://functions.yandexcloud.net/d4efh4n0sevvo2f928ri', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
      });
      if (response.ok) {
        setStatus('Фото успешно отправлено на верификацию!');
      } else {
        setStatus('Ошибка при отправке фото');
      }
    } catch {
      setStatus('Ошибка сети');
    } finally {
      setIsUploading(false);
    }
  };

  const onBack = () => {
    navigate('/profile');
  };

  return (
    <div style={{ margin: '0 auto', textAlign: 'center', width: '100%' }}>
      <button onClick={() => navigate(-1)} style={{ position: 'absolute', left: 16, top: 16 }}>
        Назад
      </button>
      <PageHeader onBack={onBack} title="Верификация" />
      <p style={{ marginTop: '15px', fontSize: '17px' }}>
        Сделайте фото с жестом виктори (пример ниже):
      </p>
      <img src="/peace.png" alt="Пример жеста" style={{ width: 200, marginBottom: 16 }} />
      <div>
        <input
          type="file"
          accept="image/*"
          capture="user"
          ref={fileInputRef}
          style={{ display: 'none' }}
          onChange={handleFileChange}
        />
        <button onClick={() => fileInputRef.current?.click()} disabled={isUploading}>
          {photo ? 'Выбрать другое фото' : 'Сделать фото'}
        </button>
      </div>
      {photo && (
        <div style={{ margin: '16px 0' }}>
          <img src={URL.createObjectURL(photo)} alt="preview" style={{ width: 150 }} />
        </div>
      )}
      <button onClick={handleSend} disabled={!photo || isUploading} style={{ marginTop: 8 }}>
        {isUploading ? 'Отправка...' : 'Отправить на верификацию'}
      </button>
      {status && <p>{status}</p>}
    </div>
  );
};

export default VerificationPage;
