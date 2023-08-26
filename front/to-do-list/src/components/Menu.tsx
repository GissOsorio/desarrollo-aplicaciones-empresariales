import {Link} from "react-router-dom";

const Menu = ({ userId }) => { // Asegúrate de recibir userId como una prop
    return (
        <div className="p-4 sm:w-1/2 lg:w-1/3 w-full hover:scale-105 duration-500">
            <div className="flex items-center justify-between p-4 rounded-lg bg-white shadow-indigo-50 shadow-md">
                <div>
                    <h2 className="text-gray-900 text-lg font-bold">Lista de tareas</h2>
                    <p className="text-sm font-semibold text-gray-400">Maneje todas sus listas de tareas</p>
                    <button className="text-sm mt-6 px-4 py-2 bg-yellow-400 text-white rounded-lg tracking-wider hover:bg-yellow-300 outline-none">
                        <Link to={`tablero-usuario/${userId}`}>Ver mis tableros</Link>
                    </button>
                </div>
                <div className="bg-gradient-to-tr from-yellow-500 to-yellow-400 w-32 h-32 rounded-full shadow-2xl shadow-yellow-400 border-white border-dashed border-2 flex justify-center items-center">
                    <div>
                        <h1 className="text-white text-2xl">Tareas</h1>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Menu;
