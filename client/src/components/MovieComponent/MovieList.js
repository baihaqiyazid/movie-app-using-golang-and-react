import React, { useEffect, useState } from 'react'
import { Button, Card, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import axios from 'axios';

export const MovieList = () => {

    const [movies, setMovies] = useState([])
    const [loaded, setLoaded] = useState(false)
    const [errorMessage, setErrorMessage] = useState(null)

    useEffect(() => {
        const fetchMovies = async () => {
            try {
                const result = await axios(`http://localhost:4001/movies`);
                await setMovies(result.data.data);
                setLoaded(true)
            } catch (err) {
                setErrorMessage(err.response.data)
            }
        }
        fetchMovies();
    }, [])

    const truncate = (str, n) => {
        return str?.length > n ? str.substr(0, n - 1) + "..." : str;
    };

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
                <Row>
                    {movies.length > 0 ? (
                        movies.map((movie) => (
                            <Col md={4} className="mb-3" key={movie.id}>
                                <Card>
                                    <Card.Body>
                                        <Card.Title>{movie.title}</Card.Title>
                                        <Card.Text>{truncate(movie.description, 90)}</Card.Text>
                                        <Link to={`/movies/${movie.id}`} className="btn btn-primary">Detail</Link>
                                    </Card.Body>
                                </Card>
                            </Col>
                        ))
                    ) : (
                        <p>No Movies</p>
                    )}
                </Row>
            )}
        </>
    )
}
