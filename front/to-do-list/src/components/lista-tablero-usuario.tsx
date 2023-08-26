import  { useState, useEffect } from 'react';
import TableroUsuario from "./tablero-usuario.tsx";
import Profile from "./profile.tsx";
import {useParams} from "react-router-dom";
import api from "../routes/api.tsx";

const ListaTableroUsuario = () => {
    const [tablerosUsuario, setTableros] = useState([]);
    const { userId } = useParams();
    const agregarTablero = nuevoTablero => {
        setTableros(prevTableros => [...prevTableros, nuevoTablero]);
    };
    useEffect(() => {
        api.get(`lists/userid/${userId}`)
            .then(response => {
                setTableros(response.data);
                console.log('response data', response.data);
            })
            .catch(error => {
                console.error('Error fetching user lists:', error);
            });
    }, [userId]);

    if (tablerosUsuario ) {
        // Muestra un indicador de carga o simplemente no renderiza nada
        return (
            <>
                <TableroUsuario tablerosUsuario={tablerosUsuario} userId={userId} />
            </>
        );
    }
    return null


}

export default ListaTableroUsuario;