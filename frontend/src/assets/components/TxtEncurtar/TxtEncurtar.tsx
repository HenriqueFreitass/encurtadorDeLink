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
        console.log(newUrl)
        if(newUrl.url != ""){
          try {
            const response = await axios.post("http://localhost:8080/shorten", newUrl);
            if (response.status === 200) {
              setNewUrl({
                ...newUrl,
                ["url"]:"",
            });
            alert("Link criado com sucesso!");
            window.location.reload();
            }
          } catch (error) {
            console.error("Erro no login:", error);
            alert("Erro ao criar o link encurtado");
          }
        }else{
          alert("Preencha o campo do link primeiramente")
        }

      }

    return(
        <div className="flex flex-col items-center m-auto mt-4 w-3/5 min-w-64">
            <p>Crie um novo link encurtado aqui</p>
            <input type="text" name="url" id="url" placeholder="Link Original" value={newUrl.url} className="w-3/4 mb-4 mt-4 rounded-md pl-3 pr-3 pt-1 pb-1"
            onChange={handleInputChange} />
            <button className="bg-indigo-500 w-32 h-8 rounded-md text-white"
            onClick={handleShortenLink}>Criar</button>
        </div>

    )
}
export default TxtEncurtar