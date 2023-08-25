import { Droppable } from 'react-beautiful-dnd';
import TodoItem from "./toDoItem.tsx";


const Section = ({ section, todos, changeStatus }) => (
    <div className="section">
        <h2>{section.toUpperCase()}</h2>
        <Droppable droppableId={section}>
            {(provided) => (
                <ul {...provided.droppableProps} ref={provided.innerRef} className="todo-list">
                    {todos.map((todo, index) => (
                        <TodoItem
                            key={todo.id}
                            todo={todo}
                            index={index}
                            changeStatus={changeStatus}
                        />
                    ))}
                    {provided.placeholder}
                </ul>
            )}
        </Droppable>
    </div>
);

export default Section;






