import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Card, Col, Row } from 'react-bootstrap'
import { Link, useParams } from 'react-router-dom'
import ShowMovieByGenre from '../../components/GenreComponent/ShowMovieByGenre'

const ShowMovies = () => {
    let { id } = useParams()
    const [movies, setMovies] = useState([])
    const [loaded, setLoaded] = useState(false)
    const [errorMessage, setErrorMessage] = useState(null)

    useEffect(() => {
        const fetchMovies = async () => {
            try {
                const result = await axios(`http://localhost:4001/genres/${id}`);
                await setMovies(result.data.data[0].movies);
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

            ) : (
                <ShowMovieByGenre movies={movies} />
            )}
        </>
    )
}

export default ShowMovies