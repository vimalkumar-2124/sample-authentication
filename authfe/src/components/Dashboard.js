import React, { useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { BaseContext } from '../App'
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import axios from 'axios';

function Dashboard() {
    let baseContext = useContext(BaseContext)
    let navigate = useNavigate()
    let [data, setData] = useState([])

    let loadData = async() => {
        let token = sessionStorage.getItem('token')
        if(token){
            let res = await axios.get(`${baseContext.apiUrl}/users/all`,{
                headers:{
                    'Authorization':`Bearer ${token}`
                }
            })
            if(res.data.statusCode === 200){
                setData(res.data.data)
            }
            else{
                alert(res.data.message)
                
                sessionStorage.clear()
                navigate('/signin')
            }
        }
        else{
            navigate('/signin')
        }
    }

    useEffect(() => {
        loadData()
    }, [])

    let logout = () => {
        sessionStorage.clear()
        navigate('/signin')
    }
    return <>
    <h1 style={{"textAlign":"center"}}>Welcome to Dashboard</h1>
      <Button onClick={()=>loadData()}>Refresh List</Button>
      &nbsp;
      <Button variant='danger' onClick={()=>logout()}> Logout</Button>
      <br></br>
      <Table striped bordered hover>
      <thead>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Mobile</th>
          <th>Email</th>
          <th>Role</th>
        </tr>
      </thead>
      <tbody>
        {
            data.map((e,i) => {
                return <tr key={i}>
                    <td>{i + 1}</td>
                    <td>{e.name}</td>
                    <td>{e.mobile}</td>
                    <td>{e.email}</td>
                    <td>{e.role}</td>
                </tr>
            })
        }
      </tbody>
    </Table>
    
    </>
}

export default Dashboard