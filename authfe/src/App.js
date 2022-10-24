
import './App.css';
import {BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom'
// import SignUp from './components/SignUp';
// import Dashboard from './components/Dashboard';
import React, { lazy, Suspense } from 'react';
const SignIn = lazy(() => import('./components/SignIn'))  // lazy loading - call the component only when its required
const SignUp = lazy(() => import('./components/SignUp'))
const Dashboard = lazy(() => import('./components/Dashboard'))
const apiUrl = 'http://localhost:8000'

export const BaseContext = React.createContext()
function App() {
  return <>
  <div>
    <BaseContext.Provider value={{apiUrl}}>

    <Router>
      <Suspense fallback={<div>Loading...</div>}>

      <Routes>
        <Route path='/signin' element={<SignIn/>}/>
        <Route path='/signup' element={<SignUp/>}/>
        <Route path='/dashboard' element={<Dashboard/>}/>
        <Route path='*' element={<Navigate to='/signin'/>}/>
      </Routes>
      </Suspense>
    </Router>
    </BaseContext.Provider>
  </div>
  </>
}

export default App;
