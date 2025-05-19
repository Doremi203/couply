import { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';

import { useUpdateUserMutation } from '../../../entities/user/api/userApi';
import { getUserId } from '../../../entities/user/model/userSlice';

interface Coordinates {
  lat: number;
  lng: number;
}

interface UseGeolocationReturn {
  coordinates: Coordinates | null;
  error: string | null;
  isLoading: boolean;
  updateUserLocation: () => Promise<boolean>;
}

export const useGeolocation = (): UseGeolocationReturn => {
  const [coordinates, setCoordinates] = useState<Coordinates | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const userId = useSelector(getUserId);
  const [updateUser] = useUpdateUserMutation();

  // Initialize with stored coordinates if they exist
  useEffect(() => {
    const storedLocation = localStorage.getItem('userLocation');
    if (storedLocation) {
      try {
        setCoordinates(JSON.parse(storedLocation));
      } catch (e) {
        console.error('Failed to parse stored location', e);
      }
    }
  }, []);

  const getPosition = (): Promise<GeolocationPosition> => {
    return new Promise((resolve, reject) => {
      if (!navigator.geolocation) {
        reject(new Error('Geolocation is not supported by your browser'));
        return;
      }

      navigator.geolocation.getCurrentPosition(resolve, reject, {
        enableHighAccuracy: true,
        timeout: 10000,
        maximumAge: 0,
      });
    });
  };

  const updateUserLocation = async (): Promise<boolean> => {
    // Check if geolocation is permitted and user is logged in
    const locationAllowed = localStorage.getItem('userLocationAllowed') === 'true';
    const isAuthenticated = !!userId && !!localStorage.getItem('token');

    if (!locationAllowed || !isAuthenticated) {
      setError('Geolocation not allowed or user not authenticated');
      return false;
    }

    setIsLoading(true);
    setError(null);

    try {
      const position = await getPosition();
      const newCoords = {
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      };

      // Save to state and localStorage
      setCoordinates(newCoords);
      localStorage.setItem('userLocation', JSON.stringify(newCoords));

      // Update user in backend
      if (userId) {
        await updateUser({
          id: userId,
          data: {
            location: JSON.stringify(newCoords),
          },
        });
      }

      setIsLoading(false);
      return true;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Unknown error occurred';
      setError(errorMessage);
      setIsLoading(false);
      return false;
    }
  };

  return {
    coordinates,
    error,
    isLoading,
    updateUserLocation,
  };
};
