import LoginButton from "../components/login-button.tsx";
import LogoutButton from "../components/logout-button.tsx";
import {useAuth0} from "@auth0/auth0-react";
import Menu from "../components/Menu.tsx";

const HomePage = () => {
    const {isAuthenticated} = useAuth0()
    return <>
        <section className="bg-center bg-no-repeat bg-[url('https://flowbite.s3.amazonaws.com/docs/jumbotron/conference.jpg')] bg-gray-700 bg-blend-multiply">
            <div className="px-4 mx-auto max-w-screen-xl text-center py-24 lg:py-56">
                <h1 className="mb-4 text-4xl font-extrabold tracking-tight leading-none text-white md:text-5xl lg:text-6xl">Aplicaciones Empresariales</h1>
                <p className="mb-8 text-lg font-normal text-gray-300 lg:text-xl sm:px-16 lg:px-48">
                    Proyecto final del Mòdulo de Aplicaciones Empresariales. Powered by Go&React.😀
                </p>
                <div className="flex flex-col space-y-4 sm:flex-row sm:justify-center sm:space-y-0 sm:space-x-4">
                    {!isAuthenticated ? <LoginButton/> : null}
                </div>
                {isAuthenticated ? <LogoutButton/> : null}
                {isAuthenticated ?  <Menu/>: null}

            </div>

        </section>
        {/*<h1>Aplicaciones empresariales</h1>*/}

        {/*{isAuthenticated ? <Profile/> : null}*/}
    </>

}

export default HomePage
