import {useParams} from "react-router-dom";
import KanbanBoard from "../components/KanbanBoard.tsx";
import TodoList from "../components/todolist-components/ToDoList.tsx";
import {useEffect, useState} from "react";
import api from "../routes/api.tsx";

const TareasPage = () => {
    let { idTablero } = useParams()
    const [tareas, setTareas] = useState([]);

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
            <p className="white">Tareas Page {idTablero}</p>
            <TodoList tareas={tareas}/>
        </section>
    )
}

export default TareasPage