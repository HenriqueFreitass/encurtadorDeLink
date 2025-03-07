import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSignOut } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';

function HeaderInventario() {
  const navigate = useNavigate();
  const [userName, setUserName] = useState<string>(''); // Estado para armazenar o nome do usuário

  useEffect(() => {
    const userId = localStorage.getItem('userId');
    if (userId) {
      fetch(`http://localhost:8080/users/${userId}`)
        .then((response) => response.json())
        .then((data) => {
            setUserName(data.Name);
        })
        .catch((error) => {
          console.error('Erro ao buscar o nome do usuário:', error);
        });
    }
  }, []);

  const handleLogout = () => {
    navigate('/');
  };

  return (
    <header className="flex justify-between items-center bg-white border-solid border-b-2 border-black h-16 pl-4">
      <h2 className="text-xl font-bold">Olá {userName || 'Usuário'}</h2> {/* Exibe o nome do usuário ou "Usuário" caso esteja vazio */}
      <div
        className="btn-sair flex justify-center items-center mr-4 bg-red-500 p-2 rounded-md border-2 cursor-pointer"
        onClick={handleLogout}
      >
        <FontAwesomeIcon icon={faSignOut} />
      </div>
    </header>
  );
}

export default HeaderInventario;
