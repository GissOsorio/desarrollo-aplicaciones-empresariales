import {useParams} from "react-router-dom";
import KanbanBoard from "../components/KanbanBoard.tsx";
import TodoList from "../components/todolist-components/ToDoList.tsx";
import {useEffect, useState} from "react";
import api from "../routes/api.tsx";

const TareasPage = () => {
    let { idTablero } = useParams()
    const [tareas, setTareas] = useState([]);
    const [tablero, setTablero] = useState({})
    useEffect(() => {
        api.get(`lists/${idTablero}`)
            .then(response => {
                setTablero(response.data);
            })
            .catch(error => {
                console.error('Error al obtener listas de usuario:', error);
            });
    }, [idTablero]);
    useEffect(() => {
        api.get(`elements/listid/${idTablero}`)
            .then(response => {
                setTareas(response.data);
            })
            .catch(error => {
                console.error('Error fetching user lists:', error);
            });
    }, [idTablero]);
    return (
        <section className="Blog">
            <div className="container-titulo-tareas">
                <div className="text-container">
                    <h3 className="titulo" >{ tablero.name }</h3>
                    <span className="detalles">Organiza y dale un seguimiento a tus actividades ✍️ </span>
                </div>

               <p className="icono-lista">
                   📖
               </p>
            </div>
            <TodoList tareas={tareas} tableroId={idTablero} onSetTareas={setTareas}/>
        </section>
    )
}

export default TareasPage