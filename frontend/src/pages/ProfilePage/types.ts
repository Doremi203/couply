export interface ProfileData {
  name: string;
  age: number;
  phone: string;
  dateOfBirth: string;
  email: string;
  gender: string;
  interests: string[];
  about: string;
  music: string[];
  movies: string[];
  books: string[];
  hobbies: string[];
  isHidden: boolean;
  photos: (string | { url: string })[];
  bio?: string;
}

export interface ActivityItem {
  type: string;
  user: string;
  date: string;
}
