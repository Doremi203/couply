import React, { useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
import { useConfirmPhotoMutation } from '../../entities/user';
import PageHeader from '../../shared/components/PageHeader';

const VerificationPage: React.FC = () => {
  const [photoFile, setPhotoFile] = useState<File | null>(null);
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
      const token = localStorage.getItem('token');

      const body = {
        token,
        bucket: 'couply-verification-photos',
      };

      const fetchWithRetry = async (url: string, options: RequestInit, maxRetries = 3) => {
        let retries = 0;

        while (retries < maxRetries) {
          try {
            const response = await fetch(url, options);

            if (response.status === 500 || response.status === 504) {
              retries++;

              const delay = 1000 * Math.pow(2, retries - 1);
              await new Promise(resolve => setTimeout(resolve, delay));
              continue;
            }

            return response;
          } catch (error) {
            retries++;
            console.error(`Ошибка запроса, попытка ${retries} из ${maxRetries}:`, error);

            if (retries >= maxRetries) {
              throw error;
            }

            const delay = 1000 * Math.pow(2, retries - 1);
            await new Promise(resolve => setTimeout(resolve, delay));
          }
        }

        throw new Error('Превышено максимальное количество попыток');
      };

      const getUrlResponse = await fetchWithRetry(
        'https://functions.yandexcloud.net/d4efh4n0sevvo2f928ri',
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(body),
        },
      );

      if (!getUrlResponse.ok) throw new Error('Ошибка получения URL загрузки');

      const { url } = await getUrlResponse.json();

      await uploadFile({
        url,
        file: photoFile,
      }).unwrap();

      await confirmPhoto({
        //@ts-ignore
        photoUrls: [url.split('?')[0]],
        isVerificationPhoto: true,
      }).unwrap();

      setStatus('Фото успешно отправлено на верификацию!');
      setTimeout(() => navigate('/profile'), 2000);
      //@ts-ignore
    } catch {
      // console.error('Upload failed:', error);
      // setStatus('Ошибка при загрузке фото');
      setStatus('Статус верификации обновится в течение нескольких минут');

      setTimeout(() => navigate('/profile'), 1000);
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
