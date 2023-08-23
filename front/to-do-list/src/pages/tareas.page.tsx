import {useParams} from "react-router-dom";
import KanbanBoard from "../components/KanbanBoard.tsx";

const TareasPage = () => {
    let { idTablero } = useParams()

    return (
        <section className="Blog">
            <p className="white">Tareas Page {idTablero}</p>
            <KanbanBoard/>
        </section>
    )
}

export default TareasPage