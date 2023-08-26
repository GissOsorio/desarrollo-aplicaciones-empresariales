import  { useState } from 'react';
import { DragDropContext } from 'react-beautiful-dnd';
import Section from "./section.tsx";

const TodoList = ({tareas, tableroId, onSetTareas}) => {


    const handleDragEnd = (result) => {

    };

    const changeStatus = (todoId) => {
        const updatedTodo = tareas.find((todo) => todo.id === todoId);
        if (!updatedTodo) {
            console.log('notfound')
            return; // Exit if the todo is not found
        }

        const updatedStatus = getNextStatus(updatedTodo.status);

        fetch('http://localhost:8080/elements/status/' + todoId, {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({ status: updatedStatus })
        })
            .then((response) => response.json())
            .then(() => {
                const updatedTodos = tareas.map((todo) =>
                    todo.id === todoId ? { ...todo, status: updatedStatus } : todo
                );

                // Update the state with the updated todos
                onSetTareas(updatedTodos);
            })
            .catch((error) => {
                console.error('Error updating status:', error);
            });
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
        }).then(response=> response.json())
            .then(idTask=> {
                const newTaskList = {
                    id: idTask,
                    ...newTask
                }
                onSetTareas([...tareas, newTaskList]);
                // setTodos([...tareas, newTaskList]);
            })

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
    const handleAddTaskClick = () => {
        const inputElement = document.querySelector('.task-input');
        const taskContent = inputElement.value;
        if (taskContent) {
            addTask(taskContent, 'todo'); // Assuming 'todo' is the default section
            inputElement.value = ''; // Clear the input field
        }
    };

    return (
       <> {tareas?
           <DragDropContext onDragEnd={handleDragEnd}>
               <div className="sections-container">
                   {sections.map((section) => (
                       <Section
                           key={section}
                           section={section}
                           todos={tareas.filter((todo) => {
                               return todo.status === section;
                           })}
                           changeStatus={changeStatus}
                           addTask={addTask}
                           sectionClass={sectionClasses[section]}
                       >
                           {/* Render the section only if there are todos */}
                           {tareas.filter((todo) => todo.status === section).length > 0 && (
                               <Section
                                   key={section}
                                   section={section}
                                   todos={tareas.filter((todo) => todo.status === section)}
                                   changeStatus={changeStatus}
                                   addTask={addTask}
                                   sectionClass={sectionClasses[section]}
                               />
                           )}
                       </Section>
                   ))}
               </div>
           </DragDropContext>
           : (
               <div className="empty-list">
                   <p>UPS.! Parece que no tienes tareas creadas aun. 😿</p>
                   <p>Ingresa una nueva tarea y podras organizarte de una mejor manera. 😄</p>
                   <div className="add-task">
                       <input
                           type="text"
                           placeholder="Nueva tarea"
                           className="task-input"
                           // Add your input handling logic here
                       />
                       <button
                           className="add-task-btn"
                           onClick={handleAddTaskClick}
                       >
                           Agregar Tarea
                       </button>
                   </div>
               </div>
           )}

       </>
    )
};

export default TodoList;