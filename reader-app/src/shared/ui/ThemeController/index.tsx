import { FC } from "react";

export type ThemeControllerProps = {
  onToggle: (value: boolean) => void;
  checked: boolean;
};

export const ThemeController: FC<ThemeControllerProps> = ({ checked, onToggle }) => {
  return (
    <input
      checked={checked}
      type="checkbox"
      className="toggle theme-controller"
      onChange={(event) => {
        onToggle(event.target.checked);
      }}
    />
  );
};
