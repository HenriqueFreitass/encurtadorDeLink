import React, { useState } from "react";
import Header from "./assets//components/Header/Header";
import LoginForm from "./assets//components/LoginForm/LoginForm";
import CadastroForm from "./assets/components/CadastroForm/CadastroForm";

function Inicial() {
  const [showLogin, setShowLogin] = useState(false);
  const [showRegister, setShowRegister] = useState(false);

  const handleLoginClick = () => {
    setShowLogin(true);
    setShowRegister(false);
  };

  const handleRegisterClick = () => {
    setShowRegister(true);
    setShowLogin(false);
  };

  return (
    <div>
      <Header onLoginClick={handleLoginClick} onRegisterClick={handleRegisterClick} />
      <div className="container mx-auto p-4 mt-8"> {/* Adicionei mt-8 aqui */}
        {/* Parágrafo da descrição */}
        <p className="text-center text-2xl font-bold text-gray-800 mb-4">
          Encurte, Compartilhe e Acompanhe!
        </p>
        <p className="text-center text-lg text-gray-600 mb-8">
          Transforme links longos em URLs curtas e fáceis de compartilhar. Simples, rápido e eficiente!
        </p>

        {/* Formulários ou outros componentes */}
        {/* Exemplo: */}
        {/* {showLogin && <LoginForm />} */}
        {/* {showRegister && <CadastroForm />} */}
      </div>
      <div className="container mx-auto p-4">
        {showLogin && <LoginForm />}
        {showRegister && <CadastroForm />}
      </div>
    </div>
  );
}

export default Inicial;