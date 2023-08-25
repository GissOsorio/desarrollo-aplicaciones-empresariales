import  { useState, useEffect } from 'react';
import TableroUsuario from "./tablero-usuario.tsx";
import Profile from "./profile.tsx";

const ListaTableroUsuario = () => {
    const [usuarios, setUsuarios] = useState([]);

    useEffect(() => {
        // Realizar la solicitud a la API al montar el componente
        fetch('https://api.example.com/users')
            .then(response => response.json())
            .then(data => setUsuarios(data))
            .catch(error => console.error('Error fetching data:', error));
    }, []);

    return (
        <>
            <TableroUsuario usuarios={usuarios} />
        </>

    );
}

export default ListaTableroUsuario;