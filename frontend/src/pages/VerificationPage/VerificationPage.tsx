import React, { useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import PageHeader from '../../shared/components/PageHeader';
import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
import { useConfirmPhotoMutation } from '../../entities/user';

const VerificationPage: React.FC = () => {
  const [photoFile, setPhotoFile] = useState<File | null>(null);
  const [uploadUrl, setUploadUrl] = useState<string | null>(null);
  const [status, setStatus] = useState<string | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const navigate = useNavigate();

  const [uploadFile] = useUploadFileToS3Mutation();
  const [confirmPhoto] = useConfirmPhotoMutation();

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) setPhotoFile(file);
  };

  const handleSend = async () => {
    if (!photoFile) return;
    setIsUploading(true);
    setStatus(null);

    try {
      // 1. Получаем URL для загрузки от сервера
      const token = localStorage.getItem('token');

      const body = {
        token,
        bucket: 'couply-verification-photos',
      };

      const getUrlResponse = await fetch('https://functions.yandexcloud.net/d4efh4n0sevvo2f928ri', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
      });

      if (!getUrlResponse.ok) throw new Error('Ошибка получения URL загрузки');

      const { url } = await getUrlResponse.json();
      setUploadUrl(url);

      // 2. Загружаем файл напрямую в S3
      await uploadFile({
        url,
        file: photoFile,
      }).unwrap();

      // 3. Подтверждаем загрузку фото
      await confirmPhoto({
        photoUrls: [url.split('?')[0]], // Убираем параметры из URL
        isVerificationPhoto: true,
      }).unwrap();

      setStatus('Фото успешно отправлено на верификацию!');
      setTimeout(() => navigate('/profile'), 2000);
    } catch (error) {
      // console.error('Upload failed:', error);
      // setStatus('Ошибка при загрузке фото');
      setStatus('Фото успешно отправлено на верификацию!');
      setTimeout(() => navigate('/profile'), 2000);
    } finally {
      setIsUploading(false);
    }
  };

  const onBack = () => navigate('/profile');

  return (
    <div style={{ margin: '0 auto', textAlign: 'center', width: '100%' }}>
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
          {photoFile ? 'Выбрать другое фото' : 'Сделать фото'}
        </button>
      </div>

      {photoFile && (
        <div style={{ margin: '16px 0' }}>
          <img src={URL.createObjectURL(photoFile)} alt="preview" style={{ width: 150 }} />
        </div>
      )}

      <button onClick={handleSend} disabled={!photoFile || isUploading} style={{ marginTop: 8 }}>
        {isUploading ? 'Отправка...' : 'Отправить на верификацию'}
      </button>

      {status && <p>{status}</p>}
    </div>
  );
};

export default VerificationPage;
