import { Settings, GpsFixed } from '@mui/icons-material';
import { Dialog, Button, Typography } from '@mui/material';
import { useState } from 'react';

const GeoLocationRequest = () => {
  const [showHelp, setShowHelp] = useState(false);
  const [isMobile, setIsMobile] = useState(false);
  const [isIOS, setIsIOS] = useState(false);

  const detectPlatform = () => {
    const userAgent = navigator.userAgent.toLowerCase();
    setIsMobile(/android|webos|iphone|ipad|ipod|blackberry|iemobile|opera mini/i.test(userAgent));
    setIsIOS(/iphone|ipad|ipod/.test(userAgent));
  };

  const handleGeoRequest = () => {
    detectPlatform();

    navigator.geolocation.getCurrentPosition(
      position => {
        // Обработка успешного получения координат
        // console.log(position.coords);
      },
      error => {
        if (error.code === error.PERMISSION_DENIED) {
          setShowHelp(true);
        }
      },
      { enableHighAccuracy: true, timeout: 10000 },
    );
  };

  const getInstructions = () => {
    if (isIOS) {
      return [
        '1. Откройте "Настройки"',
        '2. Перейдите в "Конфиденциальность"',
        '3. Выберите "Службы геолокации"',
        '4. Включите для этого сайта',
      ];
    }

    if (isMobile) {
      return [
        '1. Нажмите на три точки в правом верхнем углу',
        '2. Выберите "Настройки"',
        '3. Перейдите в "Местоположение"',
        '4. Включите доступ для этого сайта',
      ];
    }

    return [
      '1. Нажмите на иконку замка в адресной строке',
      '2. Выберите "Настройки сайта"',
      '3. Найдите пункт "Местоположение"',
      '4. Выберите "Разрешить"',
    ];
  };

  return (
    <div>
      <Button variant="contained" startIcon={<GpsFixed />} onClick={handleGeoRequest}>
        Разрешить геолокацию
      </Button>

      <Dialog open={showHelp} onClose={() => setShowHelp(false)}>
        <div style={{ padding: '20px', maxWidth: '400px' }}>
          <Typography variant="h6" gutterBottom>
            🛠 Как включить геолокацию:
          </Typography>

          {getInstructions().map((step, index) => (
            <Typography key={index} paragraph>
              {step}
            </Typography>
          ))}

          <div style={{ display: 'flex', gap: '10px', marginTop: '20px' }}>
            <Button
              variant="contained"
              color="primary"
              startIcon={<GpsFixed />}
              onClick={() => {
                setShowHelp(false);
                handleGeoRequest();
              }}
            >
              Попробовать снова
            </Button>

            {isMobile && (
              <Button
                variant="outlined"
                startIcon={<Settings />}
                onClick={() => {
                  if (isIOS) {
                    window.location.href = 'App-Prefs:Privacy&path=LOCATION';
                  } else {
                    window.location.href = 'android.settings.LOCATION_SOURCE_SETTINGS';
                  }
                }}
              >
                Открыть настройки
              </Button>
            )}
          </div>
        </div>
      </Dialog>
    </div>
  );
};

export default GeoLocationRequest;
