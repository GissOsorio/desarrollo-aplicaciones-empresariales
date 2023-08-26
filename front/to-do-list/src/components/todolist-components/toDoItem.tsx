import { Draggable } from 'react-beautiful-dnd';

const TodoItem = ({ todo, index, changeStatus }) => (
    <Draggable draggableId={todo.id} index={index}>
        {(provided) => (
            <li
                {...provided.draggableProps}
                {...provided.dragHandleProps}
                ref={provided.innerRef}
                className="todo-item"
            >
                {todo.content}
                {todo.status !== 'done' && (
                    <button
                        onClick={() => changeStatus(todo.id)}
                        className="status-button"
                    >
                        {getStatusButtonLabel(todo.status)}
                    </button>
                )}
            </li>
        )}
    </Draggable>
);

const getStatusButtonLabel = (status) => {
    switch (status) {
        case 'todo':
            return 'Iniciar';
        case 'doing':
            return 'Terminar';
        case 'done':
            return 'Abrir';
        default:
            return 'Cambiar Estado';
    }
};

export default TodoItem;