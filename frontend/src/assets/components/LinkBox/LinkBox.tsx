import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { faEye } from '@fortawesome/free-solid-svg-icons';
import "./LinkBox.css"

function LinkBox(){

    return(
        <div className="bg-white min-w-[300px] p-6 flex w-3/4 border-2 border-solid rounded-md">
        <div className="informations flex flex-col">
            <h1 className="text-2xl font-semibold">Youtube</h1>
            <a className="hover:underline" href="">bit.ly/ABC123</a>
            <a className="hover:underline" href="">www.youtube.com</a>
        </div>
        <div className="actions flex flex-col w-2/3 items-end justify-between">
            <div className="delete-btn bg-red-500 rounded-md size-7 flex justify-center items-center cursor-pointer">
                <FontAwesomeIcon icon={faTrash} />
            </div>
            <div className="flex items-center">
                <span className='block mr-1'>1</span>
                <FontAwesomeIcon icon={faEye} className='block' />
            </div>
        </div>
    </div>
    )
}

export default LinkBox