import React, { useState } from "react";
import axios from "axios";

function TxtEncurtar(){

          const [newUrl, setNewUrl] = useState({
            url: "",
            id: localStorage.getItem('userId'),
          });

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setNewUrl({
            ...newUrl,
          [name]: value,
        });
      };

      const handleShortenLink = async () =>{
        try {
            const response = await axios.post("http://localhost:8080/users/shorten", newUrl);

            if (response.status === 200) {
              alert("Login realizado com sucesso!");
              console.log(response.data)
            }
          } catch (error) {
            console.error("Erro no login:", error);
            alert("Email ou senha inválidos!");
          }
      }

    return(
        <div className="mt-4">
            <p>Crie um novo link encurtado aqui</p>
            <input type="text" name="url" id="url" placeholder="Link Encurtado" className="rounded-md pl-3 pr-3 pt-1 pb-1"
            onChange={handleInputChange} />
            <div className="bg-indigo-500 w-32 h-8"
            onClick={handleShortenLink}></div>
        </div>

    )
}
export default TxtEncurtar