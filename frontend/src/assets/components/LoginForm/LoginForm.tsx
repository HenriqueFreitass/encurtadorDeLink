import React from "react";
import "./LoginForm.css";

function LoginForm() {
  return (
    <div className="mt-4 p-4 bg-gray-100 rounded">
      <h2 className="text-lg font-semibold">Login</h2>
      <form>
        <input
          type="email"
          placeholder="Email"
          className="block w-full p-2 mt-2 border rounded"
        />
        <input
          type="password"
          placeholder="Senha"
          className="block w-full p-2 mt-2 border rounded"
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