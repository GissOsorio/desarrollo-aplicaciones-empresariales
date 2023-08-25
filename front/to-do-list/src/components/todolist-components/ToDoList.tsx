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
    const [serverUrl, setServerUrl] = useState('http://localhost:5000');

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
    const addTask = (content, section) => {
        const newTask = {
            id: String(new Date().getTime()), // Generar ID único (puedes usar una librería como uuid)
            content,
            status: section,
        };

        setTodos([...todos, newTask]);
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
    const sectionClasses = {
        todo: 'todo-section',
        doing: 'doing-section',
        done: 'done-section'
    };

    return (
        <DragDropContext onDragEnd={handleDragEnd}>
            <div className="sections-container">
                {sections.map((section) => (
                    <Section
                        key={section}
                        section={section}
                        todos={todos.filter((todo) => todo.status === section)}
                        changeStatus={changeStatus}
                        addTask={addTask}
                        sectionClass={sectionClasses[section]}
                    />
                ))}
            </div>
        </DragDropContext>
    );
};

export default TodoList;