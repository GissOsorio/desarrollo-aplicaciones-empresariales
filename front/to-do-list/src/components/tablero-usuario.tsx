import {Link} from "react-router-dom";
import {useEffect, useState} from "react";
import api from "../routes/api.tsx";
import axios from "axios";

const TableroUsuario = ({ tablerosUsuario, userId, onSetTableros }) => {
const [nuevoTablero, setNuevoTablero] = useState({ name: '', status: 'open', userId });
const [tableros, setTableros] = useState([]);
const handleInputChange = (event) => {
    const { name, value } = event.target;
    setNuevoTablero((prevTablero) => ({
        ...prevTablero,
        [name]: value,
    }));
};

    const agregarTablero = () => {
        fetch('http://localhost:8080/lists', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(nuevoTablero)
        }).then(response => response.json())
            .then(newTablero => {
                const nuevoTableroLista = {
                    id: newTablero,
                    ...nuevoTablero
                };
                onSetTableros(prevTableros => [...prevTableros, nuevoTableroLista]);
                setNuevoTablero({ name: '', status: 'open', userId });
            });
    };


    useEffect(() => {
        api.get(`lists/${userId}`)
            .then(response => {
                setTableros(response.data);
            })
            .catch(error => {
                console.error('Error al obtener listas de usuario:', error);
            });
    }, [userId]);
   return <div>
        <div>
            <div className="p-4">
                <h2 className="text-lg font-semibold mb-2">Agregar Nuevo Tablero</h2>
                <div className="flex">
                    <input
                        type="text"
                        name="name"
                        value={nuevoTablero.name}
                        onChange={handleInputChange}
                        placeholder="Nombre del tablero"
                        className="border rounded px-2 py-1 mr-2"
                    />
                    <button
                        onClick={agregarTablero}
                        className="bg-blue-500 text-white px-2 py-1 rounded"
                    >
                        Agregar
                    </button>
                </div>
            </div>
        </div>
        <div className="md:px-32 py-8 w-full">
            {!tablerosUsuario ? (
                <p>No hay Datos</p>
            ) : (
                <div className="shadow overflow-hidden rounded border-b border-gray-200">
                    <table className="min-w-full bg-white">
                        <thead className="bg-gray-800 text-white">
                        <tr>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Número</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Nombre</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Estado</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Acciones</th>
                        </tr>
                        </thead>
                        <tbody className="text-gray-700">
                        {tablerosUsuario.map((tableroUsuario, index) => (
                            <tr key={tableroUsuario.id} className={index % 2 === 0 ? "bg-gray-100" : ""}>
                                <td className="w-1/5 text-left py-3 px-4">{index + 1}</td>
                                <td className="w-1/5 text-left py-3 px-4">{tableroUsuario.name}</td>
                                <td className="w-1/5 text-left py-3 px-4">{tableroUsuario.status}</td>
                                <td className="w-1/5 text-left py-3 px-4">
                                    <button className="text-orange-500 background-transparent font-bold uppercase px-8 py-3 outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button">
                                        <Link to={`/tareas/${tableroUsuario.id}`}>Ver tareas</Link>
                                    </button>

                                </td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    </div>
}

export default TableroUsuario;