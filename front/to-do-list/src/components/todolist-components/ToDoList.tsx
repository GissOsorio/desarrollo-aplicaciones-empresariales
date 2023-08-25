import  { useState } from 'react';
import { DragDropContext } from 'react-beautiful-dnd';
import Section from "./section.tsx";

const initialTodos = [
    { id: '1', content: 'Buy groceries', status: 'todo' },
    { id: '2', content: 'Finish homework', status: 'doing' },
    { id: '3', content: 'Go for a run', status: 'done' },
];

const TodoList = () => {
    const [todos, setTodos] = useState(initialTodos);

    const handleDragEnd = (result) => {
        if (!result.destination) return;

        const updatedTodos = Array.from(todos);
        const [reorderedItem] = updatedTodos.splice(result.source.index, 1);
        updatedTodos.splice(result.destination.index, 0, reorderedItem);

        setTodos(updatedTodos);
    };

    const changeStatus = (todoId) => {
        const updatedTodos = todos.map((todo) =>
            todo.id === todoId
                ? { ...todo, status: getNextStatus(todo.status) }
                : todo
        );

        setTodos(updatedTodos);
    };

    const getNextStatus = (currentStatus) => {
        switch (currentStatus) {
            case 'todo':
                return 'doing';
            case 'doing':
                return 'done';
            default:
                return currentStatus;
        }
    };

    const sections = ['todo', 'doing', 'done'];

    return (
        <DragDropContext onDragEnd={handleDragEnd}>
            <div className="sections-container">
                {sections.map((section) => (
                    <Section
                        key={section}
                        section={section}
                        todos={todos.filter((todo) => todo.status === section)}
                        changeStatus={changeStatus}
                    />
                ))}
            </div>
        </DragDropContext>
    );
};

export default TodoList;