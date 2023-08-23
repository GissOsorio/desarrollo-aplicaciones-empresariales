import LoginButton from "../components/login-button.tsx";
import LogoutButton from "../components/logout-button.tsx";
import Profile from "../components/profile.tsx";
import {useAuth0} from "@auth0/auth0-react";

const HomePage = () => {
    const {isAuthenticated} = useAuth0()
    return <>
        <h1>Aplicaciones empresariales</h1>
        {!isAuthenticated ? <LoginButton/> : null}
        {isAuthenticated ? <LogoutButton/> : null}
        {isAuthenticated ? <Profile/> : null}
    </>

}

export default HomePage
