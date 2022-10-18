
import './App.css';
import {BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom'
import SignIn from './components/SignIn';
import SignUp from './components/SignUp';
import Dashboard from './components/Dashboard';
import React from 'react';
const apiUrl = 'http://localhost:8000'

export const BaseContext = React.createContext()
function App() {
  return <>
  <div>
    <BaseContext.Provider value={{apiUrl}}>

    <Router>
      <Routes>
        <Route path='/signin' element={<SignIn/>}/>
        <Route path='/signup' element={<SignUp/>}/>
        <Route path='/dashboard' element={<Dashboard/>}/>
        <Route path='*' element={<Navigate to='/signin'/>}/>
      </Routes>
    </Router>
    </BaseContext.Provider>
  </div>
  </>
}

export default App;
