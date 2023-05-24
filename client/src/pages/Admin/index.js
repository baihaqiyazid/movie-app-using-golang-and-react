import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { Dropdown, Row, Table } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import Swal from 'sweetalert2';

export const Admin = () => {
  const [movies, setMovies] = useState([])
  const [loaded, setLoaded] = useState(false)
  const [errorMessage, setErrorMessage] = useState(null)

  const fetchMovies = async () => {
    try {
      const result = await axios(`http://localhost:4001/movies`);
      await setMovies(result.data.data);
      setLoaded(true)
    } catch (err) {
      setErrorMessage(err.response.data)
    }
  }

  useEffect(() => {
    fetchMovies();
  }, [])

  const confirmDelete = (id) => {
    Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
      if (result.isConfirmed) {
        onDelete(id)
        Swal.fire(
          'Deleted!',
          'Your file has been deleted.',
          'success'
        )
        fetchMovies()
      }
    })
  }

  const onDelete = async(id) => {
    const payload = {
      id: id.toString(),
    }

    await axios.delete(
      'http://localhost:4001/admin/movies/delete',
      {data: payload}
    )
  }
  return (
    <>

      <Link to={`movies/create`} className="btn btn-primary mb-3">Add</Link>

      {!loaded ? (

        (() => {
          if (errorMessage) {
            return (
              <Row>
                <p>Opss... {errorMessage}</p>
              </Row>
            )
          } else {
            return (
              <Table striped bordered hover>
                <thead>
                  <tr>
                    <th>#</th>
                    <th>Movie Title</th>
                    <th></th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td colSpan={3} className="text-center">Loading...</td>
                  </tr>
                </tbody>
              </Table>
            )
          }
        })()

      ) : (
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>#</th>
              <th>Movie Title</th>
              <th></th>
            </tr>
          </thead>
          <tbody>

            { movies.length > 0 ? (
            movies.map((movie, index) => (
              <tr key={index}>
                <td>{index + 1}</td>
                <td>
                  <Link to={`/movies/${movie.id}`} className="text-decoration-none text-black">
                    {movie.title}
                  </Link>
                </td>
                <td>
                  <Link to={`movies/${movie.id}/edit`} className="text-white text-decoration-none btn btn-warning btn-sm">
                    Edit
                  </Link>
                  {' '}
                  <span
                    className="text-white text-decoration-none btn btn-danger btn-sm"
                    onClick={() => confirmDelete(movie.id)}
                  >
                    Delete
                  </span>
                </td>
              </tr>
            ))
            ):(
              <tr>
                <td colSpan={3} className="text-center">No Movies</td>
              </tr>
            )}
          </tbody>

        </Table>
      )}
    </>
  );
}
