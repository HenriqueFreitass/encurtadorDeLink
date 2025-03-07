import React, { useState } from "react";
import "./CadastroForm.css";
import axios from "axios";
import { useNavigate } from 'react-router-dom';

function CadastroForm() {

  const [formData, setFormData] = useState({
    Email: "",
    Password: "",
    Name: "",
  });

    const navigate = useNavigate();
    
  // Função para atualizar o estado quando os campos do formulário mudam
  const handleInputChange = (e: { target: { name: any; value: any; }; }) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault(); // Impede o comportamento padrão do formulário
  
    try {
      // Faz a requisição HTTP POST para o backend
      const response = await axios.post("http://localhost:8080/users/", formData);
  
      // Lida com a resposta do backend
      if (response.status === 201) {
        localStorage.setItem('userId', response.data.Id);
        navigate('/login'); 

      }else if( response.status === 500){
        alert("Já existe um usuário cadastrado com esse email")
      }

    } catch (error) {
      console.error("Erro na requisição:", error);
    }
  };

  return (
    <div className="mt-4 p-4 bg-gray-100 rounded">
      <h2 className="text-lg font-semibold">Cadastrar</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="Name"
          placeholder="Nome"
          className="block w-full p-2 mt-2 border rounded"
          value={formData.Name}
          onChange={handleInputChange}
          required
        />
        <input
          type="email"
          name="Email"
          placeholder="Email"
          className="block w-full p-2 mt-2 border rounded"
          value={formData.Email}
          onChange={handleInputChange}
          required
        />
        <input
          type="password"
          name="Password"
          placeholder="Senha"
          className="block w-full p-2 mt-2 border rounded"
          value={formData.Password}
          onChange={handleInputChange}
          required
        />
        <button
          type="submit"
          className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Cadastrar
        </button>
      </form>
    </div>
  );
}

export default CadastroForm;