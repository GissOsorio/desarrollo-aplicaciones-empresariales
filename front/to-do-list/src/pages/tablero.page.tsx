import {useAuth0} from "@auth0/auth0-react";

const Home = () => {
    const {isAuthenticated, user} = useAuth0()
    return <>
        <h1> Tablero del usuario {user?.name}</h1>

    </>

}

export default Home