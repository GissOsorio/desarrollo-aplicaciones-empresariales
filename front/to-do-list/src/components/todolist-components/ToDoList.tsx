import  { useState } from 'react';
import { DragDropContext } from 'react-beautiful-dnd';
import Section from "./section.tsx";

const TodoList = ({tareas, tableroId}) => {
    if (!Array.isArray(tareas) || tareas.length === 0) {
        return <p>No hay tareas disponibles.</p>;
    }
    const [todos, setTodos] = useState(tareas);

    const handleDragEnd = (result) => {

    };

    const changeStatus = (todoId) => {
        const updatedTodos = todos.map((todo) =>
            todo.id === todoId
                ? { ...todo, status: getNextStatus(todo.status) }
                : todo
        );
        console.log('upda', updatedTodos)
        console.log('upda333', JSON.stringify(updatedTodos.filter((element)=> element.id === todoId)[0]))
        fetch('http://localhost:8080/elements/status/'+todoId, {  // Enter your IP address here
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(updatedTodos.filter((element)=> element.id === todoId)[0]) // body data type must match "Content-Type" header
        })
        setTodos(updatedTodos);
    };
    const addTask = (content, section) => {

        const newTask = {
            content,
            status: section,
            listId: tableroId
        };
        fetch('http://localhost:8080/elements', {  // Enter your IP address here
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(newTask) // body data type must match "Content-Type" header
        })
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
                        todosFull={todos}
                        section={section}
                        todos={todos.filter((todo) => {
                            return todo.status === section;
                        })}
                        changeStatus={changeStatus}
                        addTask={addTask}
                        sectionClass={sectionClasses[section]}
                    >
                        {/* Render the section only if there are todos */}
                        {todos.filter((todo) => todo.status === section).length > 0 && (
                            <Section
                                key={section}
                                todosFull={todos}
                                section={section}
                                todos={todos.filter((todo) => todo.status === section)}
                                changeStatus={changeStatus}
                                addTask={addTask}
                                sectionClass={sectionClasses[section]}
                            />
                        )}
                    </Section>
                ))}
            </div>
        </DragDropContext>
    );
};

export default TodoList;