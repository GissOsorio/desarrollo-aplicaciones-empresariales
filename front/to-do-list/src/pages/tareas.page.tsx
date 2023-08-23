import {useParams} from "react-router-dom";

const TareasPage = () => {
    let { idTablero } = useParams()

    return (
        <section className="Blog">
            <p>Tareas Page {idTablero}</p>
        </section>
    )
}

export default TareasPage