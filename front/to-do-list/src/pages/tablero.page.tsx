import {useAuth0} from "@auth0/auth0-react";
import ListaTableroUsuario from "../components/lista-tablero-usuario.tsx";

const Home = () => {
    const {isAuthenticated, user} = useAuth0()
    return <>
        <h1> Tablero del usuario {user?.name}</h1>
        <ListaTableroUsuario/>
    </>

}

export default Home