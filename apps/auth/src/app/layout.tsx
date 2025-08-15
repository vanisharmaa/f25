import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Finan-silly",
  description:
    "This auth micro-app will handle the user signup/login/logout functionality.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
