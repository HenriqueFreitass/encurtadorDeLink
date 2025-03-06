import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Inicial from './Inicial'
import Inventario from './Inventario'
import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router>
      <Routes>
        <Route path="/login" element={<Inventario/>} />
        <Route path="/" element={<Inicial/>} />
      </Routes>
    </Router>

  </StrictMode>,
)
