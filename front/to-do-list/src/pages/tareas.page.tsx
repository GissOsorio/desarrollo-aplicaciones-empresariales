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
                console.log('response data', response.data);
            })
            .catch(error => {
                console.error('Error fetching user lists:', error);
            });
    }, [idTablero]);
    return (
        <section className="Blog">
            <p >{ tablero.name}</p>
            <TodoList tareas={tareas} tableroId={idTablero}/>
        </section>
    )
}

export default TareasPage