import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';

type Theme = 'light' | 'dark';

interface ThemeContextType {
  theme: Theme;
  toggleTheme: () => void;
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

export const ThemeProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  // Try to get the theme from localStorage, default to 'light'
  const [theme, setTheme] = useState<Theme>(() => {
    const savedTheme = localStorage.getItem('theme');
    return (savedTheme as Theme) || 'light';
  });

  // Update localStorage and meta tags when theme changes
  useEffect(() => {
    localStorage.setItem('theme', theme);
    // Apply theme class to the document body
    document.body.setAttribute('data-theme', theme);

    // Update status bar color based on theme
    const themeColor = theme === 'dark' ? '#4fd1c5' : '#202C83';

    // Update meta tags for status bar color
    const metaThemeColor = document.querySelector('meta[name="theme-color"]');
    const metaAppleStatusBar = document.querySelector(
      'meta[name="apple-mobile-web-app-status-bar-style"]',
    );

    if (metaThemeColor) {
      metaThemeColor.setAttribute('content', themeColor);
    }

    if (metaAppleStatusBar) {
      metaAppleStatusBar.setAttribute('content', themeColor);
    }
  }, [theme]);

  const toggleTheme = () => {
    setTheme(prevTheme => (prevTheme === 'light' ? 'dark' : 'light'));
  };

  return <ThemeContext.Provider value={{ theme, toggleTheme }}>{children}</ThemeContext.Provider>;
};

export const useTheme = (): ThemeContextType => {
  const context = useContext(ThemeContext);
  if (context === undefined) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  return context;
};
