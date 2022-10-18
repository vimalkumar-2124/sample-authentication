import React,{useContext, useState} from 'react'
import { BaseContext } from '../App'
import axios from 'axios'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Spinner from 'react-bootstrap/Spinner';
import { useNavigate } from 'react-router-dom';
import Alert from 'react-bootstrap/Alert';


function SignUp() {
    let baseContext = useContext(BaseContext)
    let [name, setName] = useState("")
    let [email,setEmail] = useState("")
    let [mobile, setMobile] = useState("")
    let [role, setRole] = useState("")
    let [password,setPassword] =  useState("")
    let [spinner, setSpinner] = useState(false)
    let [message, setMessage] = useState("")
    let [show, setShow] = useState(true)
    let navigate = useNavigate()
    

    let signupLogin = async() =>{
        setSpinner(true)
        let res = await axios.post(`${baseContext.apiUrl}/users/signup`,{
            name,
            email,
            mobile,
            role,
            password
        })
        if(res.data.statusCode === 201){
            setSpinner(false)
            navigate('/signin')
        }
        else{
            setSpinner(false)
            // alert(res.data.message)
            setMessage(res.data.message)
            // console.log(res.data.data)
        }
    }
 


  return <>
  <div style={{"textAlign":"center"}}>
    <div className='container'>

    <h1 className='mt-3'>Sign Up</h1>
    <p>You're one step away !! Sign-Up to Continue</p>
    {
        message ?  <Alert show={show} variant="danger" className='offset-3 col-sm-6'>
        <Alert.Heading>{message}</Alert.Heading>
        <div className="d-flex justify-content-end">
          <Button onClick={() => setShow(false)} variant="outline-danger">
           Close
          </Button>
        </div>
      </Alert>
     : <></>
    }
    <Form className='offset-3 col-6' >
    <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Name</Form.Label>
        <Form.Control type="text" placeholder="Enter name" onChange={(e) => setName(e.target.value)}/>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Email address</Form.Label>
        <Form.Control type="email" placeholder="Enter email" onChange={(e) => setEmail(e.target.value)}/>
      </Form.Group>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Mobile</Form.Label>
        <Form.Control type="mobile" placeholder="Enter mobile" onChange={(e) => setMobile(e.target.value)}/>
      </Form.Group>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Role</Form.Label>
        <Form.Control type="text" placeholder="Enter role" onChange={(e) => setRole(e.target.value)}/>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
      </Form.Group>
      <Button variant="primary"  onClick={() => signupLogin()}>
        Submit
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
    </div>
    </div>
    </>
}

export default SignUp