import React,{useContext, useState} from 'react'
import { BaseContext } from '../App'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Spinner from 'react-bootstrap/Spinner';
import axios from 'axios'
import { useNavigate } from 'react-router-dom';

function SignIn() {
    let baseContext = useContext(BaseContext)
    let [email,setEmail] = useState("")
    let [password,setPassword] =  useState("")
    let [spinner, setSpinner] = useState(false)
    let [message, setMessage] = useState("")
    let navigate = useNavigate()
    let handleLogin = async() => {
        
        setSpinner(true)
        let res = await axios.post(`${baseContext.apiUrl}/users/signin`,{
            email,
            password
        })
        if(res.data.statusCode === 200){
            setSpinner(false)
            sessionStorage.setItem('token', res.data.jwtToken)
            navigate('/dashboard')

        }
        else{
            setSpinner(false)
            setMessage(res.data.message)
            setTimeout(() => {
                setMessage("")

            },2 * 1000)
        }

        // console.log(res)
        
    }

    let signupLoginPage = async() =>{
        setSpinner(false)
        navigate('/signup')
    }
  return <>
  <div style={{"textAlign":"center"}}>
    <div className='container'>

    <h1 className='mt-3'>Sign In</h1>
    <p>You're one step away !! Sign-In to Continue</p>
    <Form className='offset-3 col-6'>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Email address</Form.Label>
        <Form.Control type="email" placeholder="Enter email" onChange={(e) => setEmail(e.target.value)}/>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
      </Form.Group>
      <Button variant="primary"  onClick={() => handleLogin()}>
        Submit
      </Button>
      &nbsp;
      &nbsp;
      <Button variant="primary"  onClick={() => signupLoginPage()}>
        Sign Up
      </Button>
    </Form>
    <br></br>
    {
        spinner ?
    <Spinner animation="border" role="status">
      <span className="visually-hidden">Loading...</span>
    </Spinner>
    :
    <></>
    }
    {
        message?<div style={{"color":"red"}}>{message}</div>:<></>
    }
    </div>
    </div>
    </>
}

export default SignIn