import  { useState, useEffect } from 'react';
import TableroUsuario from "./tablero-usuario.tsx";
import Profile from "./profile.tsx";
import {useParams} from "react-router-dom";
import api from "../routes/api.tsx";

const ListaTableroUsuario = () => {
    const [usuarios, setUsuarios] = useState([]);
    const { userId } = useParams();

    useEffect(() => {
        api.get(`lists/userid/${userId}`)
            .then(response => {
                setUsuarios(response.data);
            })
            .catch(error => {
                console.error('Error fetching user lists:', error);
            });
    }, [userId]);

    return (
        <>
            <TableroUsuario usuarios={usuarios} />
        </>
    );
}


export default ListaTableroUsuario;