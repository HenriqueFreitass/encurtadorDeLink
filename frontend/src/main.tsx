import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Inicial from './Inicial'
import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Inicial />
  </StrictMode>,
)
