import React from "react";
import "./Header.css";

interface HeaderProps {
  onLoginClick: () => void;
  onRegisterClick: () => void;
}

function Header({ onLoginClick, onRegisterClick }: HeaderProps) {
  return (
    <header className="h-28 bg-white">
      <nav className="flex items-center justify-between p-4 h-28">
          <img
            id="logo"
            src="././imagens/logo.png"
            alt="Logo"
          />
        <div className="flex gap-4">
          <button
            className="px-4 py-2 bg-blue-500 text-white rounded font-semibold hover:bg-blue-600 transition-colors cursor-pointer"
            onClick={onLoginClick}
          >
            Login
          </button>
          <button
            className="px-4 py-2 bg-green-500 text-white rounded font-semibold hover:bg-green-600 transition-colors cursor-pointer"
            onClick={onRegisterClick}
          >
            Cadastrar
          </button>
        </div>
      </nav>
    </header>
  );
}

export default Header;