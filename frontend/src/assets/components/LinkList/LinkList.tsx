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
                console.log(error)
                setLinks([]);
            });
    }, []);

    return (
        <div className="gap-4 grid grid-cols-2 px-8 mt-16">
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
                    />
                ))
            )}
        </div>
    );
}

export default LinkList;
