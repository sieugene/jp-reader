import { Navbar } from "@/shared/ui/Navbar";
import { useThemeController } from "../hooks/useThemeController";

export const NavbarWidget = () => {
  const theme = useThemeController();
  return (
    <>
      <Navbar theme={theme} />
    </>
  );
};
