import { useEffect, useState } from "react";
import LinkBox from "../LinkBox/LinkBox";

type Link = {
    Id: string;
    SiteName: string;
    OriginalUrl: string;
    NewUrl: string;
    Views: number;
};

function LinkList() {
    const [links, setLinks] = useState<Link[]>([]);

    useEffect(() => {
        fetch(`http://localhost:8080/users/link/${localStorage.getItem("userId")}`)
            .then((response) => response.json())
            .then((data) => {
                if (Array.isArray(data.links)) {
                    setLinks(data.links);
                } else {
                    setLinks([]);
                }
            })
            .catch((error) => {
                console.log(error);
                setLinks([]);
            });
    }, []);

    const handleDelete = async (linkId: string) => {
        try {
            const response = await fetch(`http://localhost:8080/users/link/${linkId}`, {
                method: 'DELETE',
            });
            if (response.status === 200) {
                setLinks(links.filter(link => link.Id !== linkId)); // Remove o link excluído da lista
                alert("Link excluído com sucesso!");
            } else {
                alert("Erro ao excluir o link");
            }
        } catch (error) {
            console.error("Erro ao excluir o link:", error);
            alert("Erro ao excluir o link");
        }
    };

    return (
        <div className="gap-4 grid grid-cols-2 px-8 mt-8">
            {links.length === 0 ? (
                <div></div>
            ) : (
                links.map((link) => (
                    <LinkBox
                        key={link.Id}
                        siteName={link.SiteName || link.OriginalUrl}
                        originalLink={link.OriginalUrl}
                        newUrl={link.NewUrl}
                        views={link.Views}
                        onDelete={() => handleDelete(link.Id)} // Passando a função em formato correto
                    />
                ))
            )}
        </div>
    );
}

export default LinkList;
