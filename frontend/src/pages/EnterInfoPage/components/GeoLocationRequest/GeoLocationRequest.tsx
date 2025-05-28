import { Settings, GpsFixed, HelpOutline } from '@mui/icons-material';
import { Dialog, Button, Typography, CircularProgress } from '@mui/material';
import { useState, useEffect } from 'react';

import { PWAGeolocationHelper } from '../PWAGeolocationHelper';

import styles from './geoLocationRequest.module.css';

interface GeoLocationRequestProps {
  onLocationReceived?: (coords: { lat: number; lng: number }) => void;
}

const GeoLocationRequest: React.FC<GeoLocationRequestProps> = ({ onLocationReceived }) => {
  const [showHelp, setShowHelp] = useState(false);
  const [showPWAHelper, setShowPWAHelper] = useState(false);
  const [isMobile, setIsMobile] = useState(false);
  const [isIOS, setIsIOS] = useState(false);
  const [isPWA, setIsPWA] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [locationStatus, setLocationStatus] = useState<'initial' | 'success' | 'error' | 'denied'>(
    'initial',
  );

  // Detect platform on component mount
  useEffect(() => {
    detectPlatform();
  }, []);

  // Determine if the app is running as a PWA
  useEffect(() => {
    // Check if the app is running in standalone mode (installed as PWA)
    const isInStandaloneMode = () =>
      window.matchMedia('(display-mode: standalone)').matches ||
      (window.navigator as any).standalone ||
      document.referrer.includes('android-app://');

    setIsPWA(isInStandaloneMode());
  }, []);

  // Detect user's platform
  const detectPlatform = () => {
    const userAgent = navigator.userAgent.toLowerCase();
    setIsMobile(/android|webos|iphone|ipad|ipod|blackberry|iemobile|opera mini/i.test(userAgent));
    setIsIOS(/iphone|ipad|ipod/.test(userAgent));
  };

  // Request geolocation permission
  const handleGeoRequest = () => {
    setIsLoading(true);
    setLocationStatus('initial');

    // Check if geolocation is supported
    if (!navigator.geolocation) {
      setLocationStatus('error');
      setIsLoading(false);
      setShowHelp(true);
      return;
    }

    navigator.geolocation.getCurrentPosition(
      position => {
        // Success handler
        setIsLoading(false);
        setLocationStatus('success');

        const coords = {
          lat: position.coords.latitude,
          lng: position.coords.longitude,
        };

        // Store coordinates in localStorage for persistence
        localStorage.setItem('userLocation', JSON.stringify(coords));

        // Mark that user has allowed geolocation
        localStorage.setItem('userLocationAllowed', 'true');

        // Call the callback if provided
        if (onLocationReceived) {
          onLocationReceived(coords);
        }
      },
      error => {
        setIsLoading(false);

        if (error.code === error.PERMISSION_DENIED) {
          setLocationStatus('denied');
          setShowHelp(true);
        } else {
          setLocationStatus('error');
          setShowHelp(true);
        }
      },
      {
        enableHighAccuracy: true,
        timeout: 10000,
        maximumAge: 0, // Always get fresh location
      },
    );
  };

  // Get platform-specific instructions
  const getInstructions = () => {
    if (isPWA && isIOS) {
      return [
        '1. Закройте приложение',
        '2. Откройте "Настройки" на вашем устройстве',
        '3. Перейдите в "Конфиденциальность" > "Службы геолокации"',
        '4. Найдите наше приложение в списке',
        '5. Включите доступ к геолокации',
        '6. Перезапустите приложение',
      ];
    } else if (isIOS) {
      return [
        '1. Откройте "Настройки" на вашем устройстве',
        '2. Перейдите в "Safari" > "Настройки сайтов" > "Местоположение"',
        '3. Найдите наш сайт и выберите "Разрешить"',
        '4. Вернитесь в браузер и обновите страницу',
      ];
    } else if (isPWA && !isIOS) {
      return [
        '1. Закройте приложение',
        '2. Откройте "Настройки" на вашем устройстве',
        '3. Перейдите в "Приложения" > Найдите наше приложение',
        '4. Выберите "Разрешения" > "Местоположение"',
        '5. Включите доступ к геолокации',
        '6. Перезапустите приложение',
      ];
    } else if (isMobile && !isIOS) {
      return [
        '1. Нажмите на иконку замка или три точки в адресной строке',
        '2. Выберите "Настройки сайта" или "Разрешения"',
        '3. Найдите пункт "Местоположение"',
        '4. Выберите "Разрешить"',
        '5. Обновите страницу',
      ];
    }

    // Desktop instructions
    return [
      '1. Нажмите на иконку замка в адресной строке',
      '2. Выберите "Настройки сайта"',
      '3. Найдите пункт "Местоположение"',
      '4. Выберите "Разрешить"',
      '5. Обновите страницу',
    ];
  };

  // Open device settings
  const openSettings = () => {
    try {
      if (isIOS) {
        // iOS settings URL schemes
        // Try multiple schemes as browser support varies
        if (isPWA) {
          // For PWA on iOS
          window.open('app-settings:', '_blank');
        } else {
          // For Safari
          window.location.href = 'App-Prefs:Privacy&path=LOCATION';

          // Fallback with timeout
          setTimeout(() => {
            window.location.href = 'prefs:root=LOCATION_SERVICES';
          }, 300);
        }
      } else if (!isIOS && isMobile) {
        // Android settings
        // Try multiple approaches

        // Intent URL for Android
        window.location.href = 'intent://settings/location#Intent;scheme=android-app;end';

        // Fallbacks with timeout
        setTimeout(() => {
          // Standard settings URL
          window.location.href = 'android.settings.LOCATION_SOURCE_SETTINGS';
        }, 300);

        setTimeout(() => {
          // Generic settings as last resort
          window.open('settings://', '_system');
        }, 600);
      } else {
        // For desktop, just close the dialog and retry
        setShowHelp(false);
      }
    } catch {
      // Show manual instructions if automatic opening fails
      alert('Пожалуйста, откройте настройки геолокации вручную, следуя инструкциям на экране.');
    }
  };

  return (
    <div className={styles.geoContainer}>
      <Button
        variant="contained"
        color="primary"
        startIcon={isLoading ? <CircularProgress size={20} color="inherit" /> : <GpsFixed />}
        onClick={handleGeoRequest}
        disabled={isLoading}
        className={styles.geoButton}
      >
        {isLoading ? 'Получение геопозиции...' : 'Разрешить геолокацию'}
      </Button>

      {locationStatus === 'success' && (
        <Typography className={styles.successMessage}>Геолокация успешно получена ✓</Typography>
      )}

      <Dialog open={showHelp} onClose={() => setShowHelp(false)}>
        <div className={styles.dialogContent}>
          <Typography variant="h6" gutterBottom>
            🛠 Как включить геолокацию:
          </Typography>

          {getInstructions().map((step, index) => (
            <Typography key={index} paragraph className={styles.instructionStep}>
              {step}
            </Typography>
          ))}

          <div className={styles.dialogButtons}>
            <Button
              variant="contained"
              color="primary"
              startIcon={<GpsFixed />}
              onClick={() => {
                setShowHelp(false);
                handleGeoRequest();
              }}
              className={styles.retryButton}
            >
              Попробовать снова
            </Button>

            {(isMobile || isPWA) && (
              <>
                <Button
                  variant="outlined"
                  startIcon={<Settings />}
                  onClick={openSettings}
                  className={styles.settingsButton}
                >
                  Открыть настройки
                </Button>

                {isPWA && (
                  <Button
                    variant="text"
                    startIcon={<HelpOutline />}
                    onClick={() => setShowPWAHelper(true)}
                    className={styles.helpButton}
                  >
                    Подробная инструкция
                  </Button>
                )}

                <Typography variant="caption" className={styles.settingsNote}>
                  Примечание: Автоматическое открытие настроек может не работать на некоторых
                  устройствах. В этом случае, пожалуйста, откройте настройки вручную.
                </Typography>
              </>
            )}
          </div>
        </div>
      </Dialog>

      {/* PWA-specific helper with detailed instructions */}
      <PWAGeolocationHelper
        open={showPWAHelper}
        onClose={() => setShowPWAHelper(false)}
        isIOS={isIOS}
      />
    </div>
  );
};

export default GeoLocationRequest;
