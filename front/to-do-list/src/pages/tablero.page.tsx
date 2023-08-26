import {useAuth0} from "@auth0/auth0-react";
import ListaTableroUsuario from "../components/lista-tablero-usuario.tsx";

const Tablero = () => {
    const {isAuthenticated, user} = useAuth0()
    return <>
        <div className="container-titulo-tareas">
            <div className="text-container">
                <h3 className="titulo" >Tablero del usuario {user?.name}</h3>
                <span className="detalles">Organiza tus tableros ✍️ </span>
            </div>

            <p className="icono-lista">
                📎
            </p>
        </div>
        <ListaTableroUsuario/>
    </>

}

export default Tablero