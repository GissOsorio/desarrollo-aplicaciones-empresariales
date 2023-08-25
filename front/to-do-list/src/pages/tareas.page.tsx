import {useParams} from "react-router-dom";
import KanbanBoard from "../components/KanbanBoard.tsx";
import TodoList from "../components/todolist-components/ToDoList.tsx";

const TareasPage = () => {
    let { idTablero } = useParams()

    return (
        <section className="Blog">
            <p className="white">Tareas Page {idTablero}</p>
            <TodoList/>
        </section>
    )
}

export default TareasPage