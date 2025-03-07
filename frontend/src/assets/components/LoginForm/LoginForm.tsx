import React, { useState } from "react";
import axios from "axios";
import "./LoginForm.css";
import { useNavigate } from 'react-router-dom';

function LoginForm() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const navigate = useNavigate();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  // Envia os dados para o servidor ao clicar em "Entrar"
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // Impede recarregamento da página

    try {
      const response = await axios.post("http://localhost:8080/users/login", formData);

      if (response.status === 200) {
        alert("Login realizado com sucesso!");
        // console.log(response.data.user.Email) Email do usuário autentificado
        localStorage.setItem('userId', response.data.user.Id);
        navigate('/login'); 
      }
    } catch (error) {
      console.error("Erro no login:", error);
      alert("Email ou senha inválidos!");
    }
  };

  return (
    <div className="mt-4 p-4 bg-gray-100 rounded">
      <h2 className="text-lg font-semibold">Login</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          name="email"
          placeholder="Email"
          className="block w-full p-2 mt-2 border rounded"
          value={formData.email}
          onChange={handleInputChange}
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Senha"
          className="block w-full p-2 mt-2 border rounded"
          value={formData.password}
          onChange={handleInputChange}
          required
        />
        <button
          type="submit"
          className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Entrar
        </button>
      </form>
    </div>
  );
}

export default LoginForm;
