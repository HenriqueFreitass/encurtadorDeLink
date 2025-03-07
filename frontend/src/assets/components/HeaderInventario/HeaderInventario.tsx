import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSignOut } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';


function HeaderInventario(){

  const navigate = useNavigate();
    const handleLogout = () =>{
        navigate('/'); 
    }

    return(
        <header className="flex justify-between items-center bg-white border-solid border-b-2 border-black h-16 pl-4 ">
            <h2 className='text-xl bold'>Olá Lucas</h2>
            <div className="btn-sair flex justify-center items-center mr-4 bg-red-500 p-2 rounded-md border-2 cursor-pointer"
            onClick={handleLogout}>
                <FontAwesomeIcon icon={faSignOut} />
            </div>
        </header>
    )
}

export default HeaderInventario