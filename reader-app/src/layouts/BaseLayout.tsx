import { NavbarWidget, FooterWidget } from "@/widgets";
import { ReactNode } from "react";

interface BaseLayoutProps {
  children: ReactNode;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  return (
    <div className="flex flex-col min-h-screen">
      <NavbarWidget />
      <main className="flex-grow container mx-auto p-4">{children}</main>
      <FooterWidget />
    </div>
  );
}
