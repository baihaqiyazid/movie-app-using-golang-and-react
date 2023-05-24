import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { Button, Card, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';

export const GenreList = () => {

    const [genres, setGenres] = useState([])
    const [loaded, setLoaded] = useState(false)
    const [errorMessage, setErrorMessage] = useState(null)

    const fetchMovies = async () => {
        try {
            const result = await axios(`http://localhost:4001/genres`)
            await setGenres(result.data.data)
            setLoaded(true)
        } catch (err) {
            setErrorMessage(err.response.data)
        }
    }

    useEffect(() => {
        fetchMovies()
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
                <Row>
                    {genres.map((genre, index) => (
                        <Col md={3} className="mb-3" key={index}>
                            <Card>
                                <Card.Body className="text-center">
                                    <Card.Title><Link to={`/genres/${genre.id}`}>{genre.name}</Link></Card.Title>
                                </Card.Body>
                            </Card>
                        </Col>
                    ))}
                </Row>
            )}
        </>
    )
}
