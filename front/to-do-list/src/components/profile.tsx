import {useAuth0} from "@auth0/auth0-react";

const profile = ()=> {
    const {user}= useAuth0()
    return <>
    <div>
        <img src={user?.picture} alt=""/>
        <h2>{user?.name} {user?.middle_name}</h2>
        <h4>{JSON.stringify(user)}</h4>
    </div>
    </>
}

export default profile;