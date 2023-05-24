import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Row } from 'react-bootstrap'
import { useParams } from 'react-router-dom'
import  MovieDetailComponent  from '../../components/MovieComponent/MovieDetailComponent'

export const MovieDetail = () => {
  let { id } = useParams()
  const [movie, setMovie] = useState([])
  const [loaded, setLoaded] = useState(false)
  const [errorMessage, setErrorMessage] = useState(null)

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const result = await axios(`http://localhost:4001/movies/${id}`);
        await setMovie(result.data.data);
        setLoaded(true)
      } catch (err) {
        setErrorMessage(err.response.data)
      }
    }
    fetchMovies();
  }, [])

  return (
    <>
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
              <Row>
                <p>Loading...</p>
              </Row>
            )
          }
        })()

      ) : (<MovieDetailComponent movie={movie} /> )}

    </>
  )
}
