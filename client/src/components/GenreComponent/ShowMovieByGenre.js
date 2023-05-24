import React from 'react'
import { Card, Col, Row } from 'react-bootstrap'
import { Link } from 'react-router-dom'

const ShowMovieByGenre = ({movies}) => {
    const truncate = (str, n) => {
        return str?.length > n ? str.substr(0, n - 1) + "..." : str;
    };
    
    return (
        <Row>
            {movies.length > 0 ? (
                Object.entries(movies).map((movie, index) =>
                    <Col md={4} className="mb-3" key={index}>
                        <Card>
                            <Card.Body>
                                <Card.Title>{movie[1].title}</Card.Title>
                                <Card.Text>{truncate(movie[1].description, 90)}</Card.Text>
                                <Link to={`/movies/${movie[1].id}`} className="btn btn-primary">Detail</Link>
                            </Card.Body>
                        </Card>
                    </Col>
                )
            ) : (
                <p>No Movies</p>
            )}
        </Row>
    )
}

export default ShowMovieByGenre