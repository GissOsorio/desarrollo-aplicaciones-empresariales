
import './App.css'
import {
    Route, Routes,
} from "react-router-dom";
import TableroPage from "./pages/tablero.page.tsx";
import HomePage from "./pages/home.page.tsx";
import TareasPage from "./pages/tareas.page.tsx";
function App() {
  return (
      <div className="App">
          <Routes>

              <Route path="/" element={ <HomePage /> } />
              <Route path="/tablero-usuario" element={ <TableroPage /> } />
              <Route path="/tareas/:idTablero" element={ <TareasPage /> } />
          </Routes>
      </div>
  )
}

export default App
