/**
 * Converts latitude and longitude coordinates to a human-readable location (city and country)
 * @param latitude - The latitude coordinate
 * @param longitude - The longitude coordinate
 * @returns Promise<string> - A promise that resolves to the location text (e.g., "New York, USA")
 */
export const getLocationFromCoordinates = async (
  latitude: number,
  longitude: number,
): Promise<string> => {
  try {
    const response = await fetch(
      `https://nominatim.openstreetmap.org/reverse?format=json&lat=${latitude}&lon=${longitude}&zoom=10&addressdetails=1`,
      {
        headers: {
          'Accept-Language': 'ru',
          'User-Agent': 'CouplyApp/1.0',
        },
      },
    );

    const data = await response.json();

    if (!data.address) {
      throw new Error('Location not found');
    }

    const { city, town, village, country } = data.address;
    const cityName = city || town || village || '';
    const countryName = country || '';

    return cityName && countryName ? `${cityName}, ${countryName}` : data.display_name;
  } catch {
    return 'Location unavailable';
  }
};
