import { Droppable } from 'react-beautiful-dnd';
import TodoItem from "./toDoItem.tsx";
import {useState} from "react";


const Section = ({ section, todos, changeStatus, addTask, sectionClass, todosFull }) => {
    const [newTaskContent, setNewTaskContent] = useState('');

    const handleAddTask = () => {
        if (newTaskContent.trim() !== '') {
            addTask(newTaskContent, section);
            setNewTaskContent('');
        }
    };

    return (
        <div className={`section ${sectionClass}`}>
            <h2 className="column-header">{section.toUpperCase()}</h2>
            <Droppable droppableId={section}>
                {(provided) => (
                    <div>
                        {todos.length > 0 ? (
                            <ul
                                {...provided.droppableProps}
                                ref={provided.innerRef}
                                className="todo-list"
                            >
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
                        ) : (
                            <p>No hay Tareas</p>
                        )}
                        {section === 'todo' && (
                            <div className="add-task">
                                <input
                                    type="text"
                                    placeholder="Nueva tarea"
                                    value={newTaskContent}
                                    onChange={(e) => setNewTaskContent(e.target.value)}
                                />
                                <button
                                    onClick={handleAddTask}
                                    className="btn-success"
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        className="icon-add"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke="currentColor"
                                    >
                                        <path
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="2"
                                            d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                                        />
                                    </svg>
                                    Agregar
                                </button>
                            </div>
                        )}
                    </div>
                )}
            </Droppable>
        </div>
    );
};

export default Section;




