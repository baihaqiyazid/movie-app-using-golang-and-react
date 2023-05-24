import React from 'react'
import { Badge, Table } from 'react-bootstrap'
import { Link } from 'react-router-dom'

const MovieDetailComponent = ({movie}) => {
    return (
        <>
            <h2>Movie : {movie.title} {movie.year}</h2>
            <div className='float-start mb-3'>
                <span>Rating : {movie.rating}</span>
            </div>

            <div className='float-end mb-3'>
                {Object.entries(movie.genres).map((genre, index) =>
                    <Link to={`/genres/${genre[1].id}`} key={index} className='text-decoration-none'>
                        <Badge bg="secondary">
                            {genre[1].name}
                        </Badge>{' '}
                    </Link>
                )}
            </div>

            <Table striped bordered hover>
                <thead>
                </thead>
                <tbody>
                    <tr>
                        <td>Title        </td>
                        <td>{movie.title}</td>
                    </tr>
                    <tr>
                        <td>Description </td>
                        <td>{movie.description}</td>
                    </tr>
                    <tr>
                        <td>Duration    </td>
                        <td>{movie.runtime} minutes</td>
                    </tr>
                </tbody>
            </Table>
        </>
    )
}

export default React.memo(MovieDetailComponent);
