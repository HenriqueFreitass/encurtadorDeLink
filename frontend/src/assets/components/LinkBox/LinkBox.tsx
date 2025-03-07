import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash, faEye } from '@fortawesome/free-solid-svg-icons';
import "./LinkBox.css";

type LinkBoxProps = {
    siteName: string;
    originalLink: string;
    newUrl: string;
    views: number;
};


function LinkBox({ siteName, originalLink, newUrl, views }: LinkBoxProps) {
    return (
        <div className="overflow-hidden bg-white min-w-[200px] h-full p-6 flex w-4/4 border-2 border-solkey justify-between rounded-md">
            <div className="informations flex flex-col">
                <h1 className="text-2xl font-semibold">{siteName}</h1>
                <a className="hover:underline text-blue-500" href={newUrl} target="_blank" rel="noopener noreferrer">
                    {newUrl}
                </a>
                <a className="hover:underline text-gray-600" href={originalLink} target="_blank" rel="noopener noreferrer">
                    {originalLink}
                </a>
            </div>
            <div className="actions flex flex-col items-end justify-between ">
                <div className="delete-btn bg-red-500 text-white rounded-md size-7 flex justify-center items-center cursor-pointer">
                    <FontAwesomeIcon icon={faTrash} />
                </div>
                <div className="flex items-center">
                    <span className='block mr-1'>{views}</span>
                    <FontAwesomeIcon icon={faEye} className='block' />
                </div>
            </div>
        </div>
    );
}

export default LinkBox;
