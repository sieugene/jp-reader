import { ThemeControllerProps } from "@/shared/ui/ThemeController";
import { useEffect, useState } from "react";

enum Theme {
  "light" = "light",
  "dark" = "dark",
}
const STORAGE_THEME_KEY = "theme";

const getTheme = (checked: boolean) => {
  return checked ? Theme.dark : Theme.light;
};

export const useThemeController = (): ThemeControllerProps => {
  const [theme, setTheme] = useState<Theme>(
    (localStorage.getItem(STORAGE_THEME_KEY) as Theme) || Theme.light,
  );
  const checked = theme === Theme.dark;

  const toggleTheme = (value: boolean) => {
    const nextTheme = getTheme(value);

    localStorage.setItem(STORAGE_THEME_KEY, nextTheme);
    setTheme(nextTheme);
  };

  useEffect(() => {
    document.documentElement.setAttribute("data-theme", theme);
  }, [theme]);

  return {
    checked: checked,
    onToggle: (value) => {
      toggleTheme(value);
    },
  };
};
