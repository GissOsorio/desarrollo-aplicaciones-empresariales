import {useAuth0} from "@auth0/auth0-react";

const LogoutButton = ()=> {
    const {logout} = useAuth0();
    return(
        <a onClick={()=> logout()} className="inline-flex justify-center items-center py-3 px-5 text-base font-medium text-center text-white rounded-lg bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900">
            Salir
            <svg aria-hidden="true" className="w-4 h-4" fill="currentColor" style={{ marginLeft: '5px' }} viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd" d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 111.414 1.414L4.414 9H16a1 1 0 010 2H4.414l4.293 4.293a1 1 0 010 1.414z" clip-rule="evenodd" />
            </svg>
        </a>
    )
}
export default LogoutButton;