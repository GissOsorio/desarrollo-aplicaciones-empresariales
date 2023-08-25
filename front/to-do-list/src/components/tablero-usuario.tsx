import {Link} from "react-router-dom";
const tablerosUsuarioEjemplo = [
    { id: 1, nombre: 'Tablero 1', descripcion: 'Descripción del Tablero 1', estado: 'Activo' },
    { id: 2, nombre: 'Tablero 2', descripcion: 'Descripción del Tablero 2', estado: 'Inactivo' },
    { id: 3, nombre: 'Tablero 3', descripcion: 'Descripción del Tablero 3', estado: 'Pendiente' },
    // ... puedes agregar más objetos según tus necesidades
];
const TableroUsuario = ({ tablerosUsuario=tablerosUsuarioEjemplo }) => (
    <div>
        <div className="md:px-32 py-8 w-full">
            {tablerosUsuario.length === 0 ? (
                <p>No hay Datos</p>
            ) : (
                <div className="shadow overflow-hidden rounded border-b border-gray-200">
                    <table className="min-w-full bg-white">
                        <thead className="bg-gray-800 text-white">
                        <tr>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Número</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Nombre</th>
                            <th className="w-2/5 text-left py-3 px-4 uppercase font-semibold text-sm">Descripción</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Estado</th>
                            <th className="w-1/5 text-left py-3 px-4 uppercase font-semibold text-sm">Acciones</th>
                        </tr>
                        </thead>
                        <tbody className="text-gray-700">
                        {tablerosUsuario.map((tableroUsuario, index) => (
                            <tr key={tableroUsuario.id} className={index % 2 === 0 ? "bg-gray-100" : ""}>
                                <td className="w-1/5 text-left py-3 px-4">{index + 1}</td>
                                <td className="w-1/5 text-left py-3 px-4">{tableroUsuario.nombre}</td>
                                <td className="w-2/5 text-left py-3 px-4">{tableroUsuario.descripcion}</td>
                                <td className="w-1/5 text-left py-3 px-4">{tableroUsuario.estado}</td>
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
);

export default TableroUsuario;